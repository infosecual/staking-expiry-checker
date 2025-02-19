package services

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/networks/parameters/parser"
	"github.com/babylonlabs-io/staking-expiry-checker/internal/db/model"
	"github.com/babylonlabs-io/staking-expiry-checker/internal/types"
	"github.com/btcsuite/btcd/wire"
	"github.com/lightningnetwork/lnd/chainntnfs"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
)

func GetTypeProvider(data []byte) (*go_fuzz_utils.TypeProvider, error) {
	tp, err := go_fuzz_utils.NewTypeProvider(data)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsStringBounds(0, 1024)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsSliceBounds(0, 4096)
	if err != nil {
		return nil, err
	}
	err = tp.SetParamsBiases(0, 0, 0, 0)
	if err != nil {
		return nil, err
	}
	return tp, nil
}

func Fuzz_Nosy_Service_IsValidUnbondingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var delegation *model.DelegationDocument
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var params *parser.ParsedVersionedGlobalParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil || delegation == nil || params == nil {
			return
		}

		s.IsValidUnbondingTx(tx, delegation, params)
	})
}

func Fuzz_Nosy_Service_RunUntilShutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.RunUntilShutdown(ctx)
	})
}

func Fuzz_Nosy_Service_SaveNewTimeLockExpire__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var startHeight uint64
		fill_err = tp.Fill(&startHeight)
		if fill_err != nil {
			return
		}
		var timelock uint64
		fill_err = tp.Fill(&timelock)
		if fill_err != nil {
			return
		}
		var txType types.StakingTxType
		fill_err = tp.Fill(&txType)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.SaveNewTimeLockExpire(ctx, stakingTxHashHex, startHeight, timelock, txType)
	})
}

func Fuzz_Nosy_Service_TransitionToUnbondedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxType types.StakingTxType
		fill_err = tp.Fill(&stakingTxType)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.TransitionToUnbondedState(ctx, stakingTxType, stakingTxHashHex)
	})
}

func Fuzz_Nosy_Service_getVersionedParams__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.getVersionedParams(height)
	})
}

func Fuzz_Nosy_Service_handleSpendingStakingTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var spendingTx *wire.MsgTx
		fill_err = tp.Fill(&spendingTx)
		if fill_err != nil {
			return
		}
		var spendingHeight uint32
		fill_err = tp.Fill(&spendingHeight)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if s == nil || spendingTx == nil {
			return
		}

		s.handleSpendingStakingTransaction(ctx, spendingTx, spendingHeight, spendingInputIdx, stakingTxHashHex)
	})
}

func Fuzz_Nosy_Service_handleSpendingUnbondingTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var spendingTx *wire.MsgTx
		fill_err = tp.Fill(&spendingTx)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if s == nil || spendingTx == nil {
			return
		}

		s.handleSpendingUnbondingTransaction(ctx, spendingTx, spendingInputIdx, stakingTxHashHex)
	})
}

func Fuzz_Nosy_Service_handleUnbondingDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.handleUnbondingDelegation(ctx)
	})
}

func Fuzz_Nosy_Service_handleWithdrawnDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.handleWithdrawnDelegation(ctx)
	})
}

func Fuzz_Nosy_Service_processBTCSubscriber__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processBTCSubscriber(ctx)
	})
}

func Fuzz_Nosy_Service_processExpiredDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.processExpiredDelegations(ctx)
	})
}

func Fuzz_Nosy_Service_quitContext__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.quitContext()
	})
}

func Fuzz_Nosy_Service_registerStakingSpendNotification__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var stakingTxHex string
		fill_err = tp.Fill(&stakingTxHex)
		if fill_err != nil {
			return
		}
		var stakingOutputIdx uint32
		fill_err = tp.Fill(&stakingOutputIdx)
		if fill_err != nil {
			return
		}
		var stakingStartHeight uint32
		fill_err = tp.Fill(&stakingStartHeight)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.registerStakingSpendNotification(stakingTxHashHex, stakingTxHex, stakingOutputIdx, stakingStartHeight)
	})
}

func Fuzz_Nosy_Service_registerUnbondingSpendNotification__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		var unbondingTxHex string
		fill_err = tp.Fill(&unbondingTxHex)
		if fill_err != nil {
			return
		}
		var spendHeightHint uint32
		fill_err = tp.Fill(&spendHeightHint)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.registerUnbondingSpendNotification(stakingTxHashHex, unbondingTxHex, spendHeightHint)
	})
}

func Fuzz_Nosy_Service_startBTCSubscriberPoller__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.startBTCSubscriberPoller(ctx)
	})
}

func Fuzz_Nosy_Service_startExpiryPoller__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if s == nil {
			return
		}

		s.startExpiryPoller(ctx)
	})
}

func Fuzz_Nosy_Service_validateWithdrawalTxFromStaking__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var delegation *model.DelegationDocument
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var params *parser.ParsedVersionedGlobalParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil || delegation == nil || params == nil {
			return
		}

		s.validateWithdrawalTxFromStaking(tx, spendingInputIdx, delegation, params)
	})
}

func Fuzz_Nosy_Service_validateWithdrawalTxFromUnbonding__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		var delegation *model.DelegationDocument
		fill_err = tp.Fill(&delegation)
		if fill_err != nil {
			return
		}
		var spendingInputIdx uint32
		fill_err = tp.Fill(&spendingInputIdx)
		if fill_err != nil {
			return
		}
		var params *parser.ParsedVersionedGlobalParams
		fill_err = tp.Fill(&params)
		if fill_err != nil {
			return
		}
		if s == nil || tx == nil || delegation == nil || params == nil {
			return
		}

		s.validateWithdrawalTxFromUnbonding(tx, delegation, spendingInputIdx, params)
	})
}

func Fuzz_Nosy_Service_watchForSpendStakingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var spendEvent *chainntnfs.SpendEvent
		fill_err = tp.Fill(&spendEvent)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if s == nil || spendEvent == nil {
			return
		}

		s.watchForSpendStakingTx(spendEvent, stakingTxHashHex)
	})
}

func Fuzz_Nosy_Service_watchForSpendUnbondingTx__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var s *Service
		fill_err = tp.Fill(&s)
		if fill_err != nil {
			return
		}
		var spendEvent *chainntnfs.SpendEvent
		fill_err = tp.Fill(&spendEvent)
		if fill_err != nil {
			return
		}
		var stakingTxHashHex string
		fill_err = tp.Fill(&stakingTxHashHex)
		if fill_err != nil {
			return
		}
		if s == nil || spendEvent == nil {
			return
		}

		s.watchForSpendUnbondingTx(spendEvent, stakingTxHashHex)
	})
}

func Fuzz_Nosy_TrackedSubscriptions_AddSubscription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		ts := NewTrackedSubscriptions()
		ts.AddSubscription(stakingTxHash)
	})
}

func Fuzz_Nosy_TrackedSubscriptions_IsSubscribed__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		ts := NewTrackedSubscriptions()
		ts.IsSubscribed(stakingTxHash)
	})
}

func Fuzz_Nosy_TrackedSubscriptions_RemoveSubscription__(f *testing.F) {
	f.Fuzz(func(t *testing.T, stakingTxHash string) {
		ts := NewTrackedSubscriptions()
		ts.RemoveSubscription(stakingTxHash)
	})
}
