package stakerdb_test

import (
	"bytes"
	"errors"
	"math/rand"
	"testing"
	"time"

	"github.com/babylonlabs-io/babylon/testutil/datagen"
	"github.com/babylonlabs-io/btc-staker/proto"
	"github.com/babylonlabs-io/btc-staker/stakercfg"
	"github.com/babylonlabs-io/btc-staker/stakerdb"
	"github.com/btcsuite/btcd/btcec/v2"
	"github.com/btcsuite/btcd/btcec/v2/schnorr"
	"github.com/btcsuite/btcd/btcutil"
	"github.com/btcsuite/btcd/chaincfg"
	"github.com/btcsuite/btcd/chaincfg/chainhash"
	"github.com/stretchr/testify/require"
)

func MakeTestStore(t *testing.T) *stakerdb.TrackedTransactionStore {
	// First, create a temporary directory to be used for the duration of
	// this test.
	tempDirName := t.TempDir()

	cfg := stakercfg.DefaultDBConfig()

	cfg.DBPath = tempDirName

	backend, err := stakercfg.GetDBBackend(&cfg)
	require.NoError(t, err)

	t.Cleanup(func() {
		backend.Close()
	})

	store, err := stakerdb.NewTrackedTransactionStore(backend)
	require.NoError(t, err)

	return store
}

func pubKeysEqual(pk1, pk2 *btcec.PublicKey) bool {
	return bytes.Equal(schnorr.SerializePubKey(pk1), schnorr.SerializePubKey(pk2))
}

func pubKeysSliceEqual(pk1, pk2 []*btcec.PublicKey) bool {
	if len(pk1) != len(pk2) {
		return false
	}

	for i := 0; i < len(pk1); i++ {
		if !pubKeysEqual(pk1[i], pk2[i]) {
			return false
		}
	}

	return true
}

func genStoredTransaction(t *testing.T, r *rand.Rand, maxStakingTime uint16) *stakerdb.StoredTransaction {
	btcTx := datagen.GenRandomTx(r)
	outputIdx := r.Uint32()
	priv, err := btcec.NewPrivateKey()
	require.NoError(t, err)
	stakingTime := r.Int31n(int32(maxStakingTime)) + 1

	stakerAddr, err := datagen.GenRandomBTCAddress(r, &chaincfg.MainNetParams)
	require.NoError(t, err)
	require.NoError(t, err)

	numPubKeys := r.Intn(3) + 1

	fpBtcPks := make([]*btcec.PublicKey, numPubKeys)
	for i := 0; i < numPubKeys; i++ {
		fpBtcPks[i] = priv.PubKey()
	}

	return &stakerdb.StoredTransaction{
		StakingTx:               btcTx,
		StakingOutputIndex:      outputIdx,
		StakingTime:             uint16(stakingTime),
		FinalityProvidersBtcPks: fpBtcPks,
		Pop: &stakerdb.ProofOfPossession{
			BtcSigOverBabylonAddr: datagen.GenRandomByteArray(r, 64),
		},
		StakerAddress: stakerAddr.String(),
	}
}

func genNStoredTransactions(t *testing.T, r *rand.Rand, n int, maxStakingTime uint16) []*stakerdb.StoredTransaction {
	storedTxs := make([]*stakerdb.StoredTransaction, n)

	for i := 0; i < n; i++ {
		storedTxs[i] = genStoredTransaction(t, r, maxStakingTime)
	}

	return storedTxs
}

func TestEmptyStore(t *testing.T) {
	t.Parallel()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := MakeTestStore(t)
	hash := datagen.GenRandomBtcdHash(r)
	tx, err := s.GetTransaction(&hash)
	require.Nil(t, tx)
	require.Error(t, err)
	require.True(t, errors.Is(err, stakerdb.ErrTransactionNotFound))
}

func FuzzStoringTxs(f *testing.F) {
	// only 3 seeds as this is pretty slow test opening/closing db
	datagen.AddRandomSeedsToFuzzer(f, 3)

	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		s := MakeTestStore(t)
		maxCreatedTx := 30
		numTx := r.Intn(maxCreatedTx) + 1
		generatedStoredTxs := genNStoredTransactions(t, r, numTx, 200)
		for _, storedTx := range generatedStoredTxs {
			stakerAddr, err := btcutil.DecodeAddress(storedTx.StakerAddress, &chaincfg.MainNetParams)
			require.NoError(t, err)
			require.NoError(t, err)
			err = s.AddTransactionSentToBTC(
				storedTx.StakingTx,
				storedTx.StakingOutputIndex,
				storedTx.StakingTime,
				storedTx.FinalityProvidersBtcPks,
				storedTx.Pop,
				stakerAddr,
			)
			require.NoError(t, err)
		}
		var expectedIdx uint64 = 1
		for _, storedTx := range generatedStoredTxs {
			hash := storedTx.StakingTx.TxHash()
			tx, err := s.GetTransaction(&hash)
			require.NoError(t, err)
			require.Equal(t, storedTx.StakingTx, tx.StakingTx)
			require.Equal(t, storedTx.StakingOutputIndex, tx.StakingOutputIndex)
			require.Equal(t, storedTx.StakingTime, tx.StakingTime)
			require.True(t, pubKeysSliceEqual(storedTx.FinalityProvidersBtcPks, tx.FinalityProvidersBtcPks))
			require.Equal(t, storedTx.Pop, tx.Pop)
			require.Equal(t, storedTx.StakerAddress, tx.StakerAddress)
			require.Equal(t, expectedIdx, tx.StoredTransactionIdx)
			expectedIdx++
		}

		storedResult, err := s.QueryStoredTransactions(stakerdb.DefaultStoredTransactionQuery())
		require.NoError(t, err)

		require.Equal(t, len(generatedStoredTxs), len(storedResult.Transactions))
		require.Equal(t, len(generatedStoredTxs), int(storedResult.Total))

		// transactions are returned in order of insertion
		for i, storedTx := range generatedStoredTxs {
			require.Equal(t, storedTx.StakingTx, storedResult.Transactions[i].StakingTx)
			require.Equal(t, storedTx.StakingOutputIndex, storedResult.Transactions[i].StakingOutputIndex)
			require.Equal(t, storedTx.StakingTime, storedResult.Transactions[i].StakingTime)
			require.True(t, pubKeysSliceEqual(storedTx.FinalityProvidersBtcPks, storedResult.Transactions[i].FinalityProvidersBtcPks))
			require.Equal(t, storedTx.Pop, storedResult.Transactions[i].Pop)
			require.Equal(t, storedTx.StakerAddress, storedResult.Transactions[i].StakerAddress)
		}

		// scan transactions
		i := 0
		err = s.ScanTrackedTransactions(func(tx *stakerdb.StoredTransaction) error {
			require.Equal(t, generatedStoredTxs[i].StakingTx, tx.StakingTx)
			i++
			return nil
		}, func() {})
		require.NoError(t, err)
	})
}

func TestStateTransitions(t *testing.T) {
	t.Parallel()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := MakeTestStore(t)
	tx := genStoredTransaction(t, r, 200)
	stakerAddr, err := btcutil.DecodeAddress(tx.StakerAddress, &chaincfg.MainNetParams)
	require.NoError(t, err)
	txHash := tx.StakingTx.TxHash()
	err = s.AddTransactionSentToBTC(
		tx.StakingTx,
		tx.StakingOutputIndex,
		tx.StakingTime,
		tx.FinalityProvidersBtcPks,
		tx.Pop,
		stakerAddr,
	)
	require.NoError(t, err)

	// Initial state
	storedTx, err := s.GetTransaction(&txHash)
	require.NoError(t, err)
	require.Equal(t, proto.TransactionState_SENT_TO_BTC, storedTx.State)
	require.Equal(t, uint64(1), storedTx.StoredTransactionIdx)
	// Confirmed
	hash := datagen.GenRandomBtcdHash(r)
	height := r.Uint32()

	err = s.SetTxConfirmed(&txHash, &hash, height)
	require.NoError(t, err)
	storedTx, err = s.GetTransaction(&txHash)
	require.NoError(t, err)
	require.Equal(t, proto.TransactionState_CONFIRMED_ON_BTC, storedTx.State)
	require.NotNil(t, storedTx.StakingTxConfirmationInfo)
	require.True(t, hash.IsEqual(&storedTx.StakingTxConfirmationInfo.BlockHash))
	require.Equal(t, height, storedTx.StakingTxConfirmationInfo.Height)

	// Sent to Babylon
	err = s.SetTxSentToBabylon(&txHash, hash.String(), tx.StakingTx, tx.StakingTime)
	require.NoError(t, err)
	storedTx, err = s.GetTransaction(&txHash)
	require.NoError(t, err)
	require.Equal(t, proto.TransactionState_SENT_TO_BABYLON, storedTx.State)

	// Spent on BTC
	err = s.SetTxSpentOnBtc(&txHash)
	require.NoError(t, err)
	storedTx, err = s.GetTransaction(&txHash)
	require.NoError(t, err)
	require.Equal(t, proto.TransactionState_SPENT_ON_BTC, storedTx.State)
	require.NotNil(t, storedTx.UnbondingTxData)
	require.Equal(t, tx.StakingTx, storedTx.UnbondingTxData.UnbondingTx)
	require.Equal(t, tx.StakingTime, storedTx.UnbondingTxData.UnbondingTime)
}

func TestPaginator(t *testing.T) {
	t.Parallel()
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	s := MakeTestStore(t)
	numTx := 45
	batchSize := 20

	generatedStoredTxs := genNStoredTransactions(t, r, numTx, 200)
	for _, storedTx := range generatedStoredTxs {
		stakerAddr, err := btcutil.DecodeAddress(storedTx.StakerAddress, &chaincfg.MainNetParams)
		require.NoError(t, err)
		err = s.AddTransactionSentToBTC(
			storedTx.StakingTx,
			storedTx.StakingOutputIndex,
			storedTx.StakingTime,
			storedTx.FinalityProvidersBtcPks,
			storedTx.Pop,
			stakerAddr,
		)
		require.NoError(t, err)
	}

	query := stakerdb.DefaultStoredTransactionQuery()
	query.IndexOffset = 0
	query.NumMaxTransactions = uint64(batchSize)
	storedResult1, err := s.QueryStoredTransactions(query)

	require.NoError(t, err)
	require.Equal(t, batchSize, len(storedResult1.Transactions))
	require.Equal(t, numTx, int(storedResult1.Total))

	query = stakerdb.DefaultStoredTransactionQuery()
	query.IndexOffset = uint64(batchSize)
	query.NumMaxTransactions = uint64(batchSize)
	storedResult2, err := s.QueryStoredTransactions(query)

	require.NoError(t, err)
	require.Equal(t, batchSize, len(storedResult2.Transactions))
	require.Equal(t, numTx, int(storedResult2.Total))

	query = stakerdb.DefaultStoredTransactionQuery()
	query.IndexOffset = 2 * uint64(batchSize)
	query.NumMaxTransactions = uint64(batchSize)
	storedResult3, err := s.QueryStoredTransactions(query)
	require.NoError(t, err)
	// 2 batches of 20, 1 batch of 5
	require.Equal(t, 5, len(storedResult3.Transactions))
	require.Equal(t, numTx, int(storedResult3.Total))

	var allTransactionsFromDB []stakerdb.StoredTransaction
	allTransactionsFromDB = append(allTransactionsFromDB, storedResult1.Transactions...)
	allTransactionsFromDB = append(allTransactionsFromDB, storedResult2.Transactions...)
	allTransactionsFromDB = append(allTransactionsFromDB, storedResult3.Transactions...)

	require.Equal(t, len(generatedStoredTxs), len(allTransactionsFromDB))
	for i, storedTx := range generatedStoredTxs {
		require.Equal(t, storedTx.StakingTx, allTransactionsFromDB[i].StakingTx)
		require.Equal(t, storedTx.StakingOutputIndex, allTransactionsFromDB[i].StakingOutputIndex)
		require.Equal(t, storedTx.StakingTime, allTransactionsFromDB[i].StakingTime)
		require.True(t, pubKeysSliceEqual(storedTx.FinalityProvidersBtcPks, allTransactionsFromDB[i].FinalityProvidersBtcPks))
		require.Equal(t, storedTx.Pop, allTransactionsFromDB[i].Pop)
		require.Equal(t, storedTx.StakerAddress, allTransactionsFromDB[i].StakerAddress)
	}
}

func FuzzQuerySpendableTx(f *testing.F) {
	// only 3 seeds as this is pretty slow test opening/closing db
	datagen.AddRandomSeedsToFuzzer(f, 3)

	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		s := MakeTestStore(t)
		// generate random transactions between 20 and 50
		maxCreatedTx := int(r.Int31n(31) + 20)
		// random staking time between 150 and 250 blocks
		maxStakingTime := r.Int31n(101) + 150
		stored := genNStoredTransactions(t, r, maxCreatedTx, uint16(maxStakingTime))
		for _, storedTx := range stored {
			stakerAddr, err := btcutil.DecodeAddress(storedTx.StakerAddress, &chaincfg.MainNetParams)
			require.NoError(t, err)
			require.NoError(t, err)
			err = s.AddTransactionSentToBTC(
				storedTx.StakingTx,
				storedTx.StakingOutputIndex,
				storedTx.StakingTime,
				storedTx.FinalityProvidersBtcPks,
				storedTx.Pop,
				stakerAddr,
			)
			require.NoError(t, err)
		}

		query := stakerdb.DefaultStoredTransactionQuery()
		// random confirmation block
		confirmationBlock := uint32(r.Int31n(1000) + 1)
		halfOfMaxStaking := maxStakingTime / 2
		currentBestBlock := confirmationBlock + uint32(r.Int31n(halfOfMaxStaking)+1)
		filteredQuery := query.WithdrawableTransactionsFilter(currentBestBlock)

		var hashesWithExpiredTimeLock []*chainhash.Hash
		for _, storedTx := range stored {
			if currentBestBlock+1 >= uint32(storedTx.StakingTime)+confirmationBlock {
				txHash := storedTx.StakingTx.TxHash()
				hashesWithExpiredTimeLock = append(hashesWithExpiredTimeLock, &txHash)
			}
		}

		storedResult, err := s.QueryStoredTransactions(filteredQuery)
		require.NoError(t, err)
		// at this point, all transactions should be non spendable
		require.Len(t, storedResult.Transactions, 0)

		for _, storedTx := range stored {
			txHash := storedTx.StakingTx.TxHash()
			err := s.SetTxConfirmed(&txHash, &txHash, confirmationBlock)
			require.NoError(t, err)
		}

		storedResult, err = s.QueryStoredTransactions(filteredQuery)
		require.NoError(t, err)
		require.Len(t, storedResult.Transactions, len(hashesWithExpiredTimeLock))
		require.Equal(t, storedResult.Total, uint64(maxCreatedTx))

		for _, storedTx := range stored {
			txHash := storedTx.StakingTx.TxHash()
			err := s.SetTxSentToBabylon(
				&txHash,
				txHash.String(),
				storedTx.StakingTx,
				storedTx.StakingTime,
			)
			require.NoError(t, err)
		}

		storedResult, err = s.QueryStoredTransactions(filteredQuery)
		require.NoError(t, err)
		require.Len(t, storedResult.Transactions, len(hashesWithExpiredTimeLock))
		require.Equal(t, storedResult.Total, uint64(maxCreatedTx))

		for _, storedTx := range stored {
			txHash := storedTx.StakingTx.TxHash()
			err := s.SetTxUnbondingConfirmedOnBtc(
				&txHash,
				&txHash,
				confirmationBlock,
			)
			require.NoError(t, err)
		}

		// we should receive the same resuls as with staking confirmed transactions
		storedResult, err = s.QueryStoredTransactions(filteredQuery)
		require.NoError(t, err)
		require.Len(t, storedResult.Transactions, len(hashesWithExpiredTimeLock))
		require.Equal(t, storedResult.Total, uint64(maxCreatedTx))
	})
}

func FuzzTrackInputs(f *testing.F) {
	// only 3 seeds as this is pretty slow test opening/closing db
	datagen.AddRandomSeedsToFuzzer(f, 3)

	f.Fuzz(func(t *testing.T, seed int64) {
		r := rand.New(rand.NewSource(seed))
		s := MakeTestStore(t)
		numTx := 45
		// all gene
		generatedStoredTxs := genNStoredTransactions(t, r, numTx, 200)

		randomUnbondingTx := datagen.GenRandomTx(r)
		unbodningTime := uint16(r.Int31n(int32(200)) + 1)

		for _, storedTx := range generatedStoredTxs {
			stakerAddr, err := btcutil.DecodeAddress(storedTx.StakerAddress, &chaincfg.MainNetParams)
			require.NoError(t, err)
			err = s.AddTransactionSentToBabylon(
				storedTx.StakingTx,
				storedTx.StakingOutputIndex,
				storedTx.StakingTime,
				storedTx.FinalityProvidersBtcPks,
				storedTx.Pop,
				stakerAddr,
				randomUnbondingTx,
				unbodningTime,
				"txHashBtcDelegation",
			)
			require.NoError(t, err)
		}

		for _, storedTx := range generatedStoredTxs {
			// check all inputs are used
			for _, inp := range storedTx.StakingTx.TxIn {
				used, err := s.OutpointUsed(&inp.PreviousOutPoint)
				require.NoError(t, err)
				require.True(t, used)
			}
		}

		// generate few not saved transactions
		notSaved := genNStoredTransactions(t, r, 20, 200)

		// check all input are not used
		for _, storedTx := range notSaved {
			for _, inp := range storedTx.StakingTx.TxIn {
				used, err := s.OutpointUsed(&inp.PreviousOutPoint)
				require.NoError(t, err)
				require.False(t, used)
			}
		}

		for _, storedTx := range generatedStoredTxs {
			stakingTxHash := storedTx.StakingTx.TxHash()
			err := s.SetDelegationActiveOnBabylon(&stakingTxHash)
			require.NoError(t, err)
		}

		for _, storedTx := range generatedStoredTxs {
			// check all inputs are deleted
			for _, inp := range storedTx.StakingTx.TxIn {
				used, err := s.OutpointUsed(&inp.PreviousOutPoint)
				require.NoError(t, err)
				require.False(t, used)
			}
		}
	})
}
