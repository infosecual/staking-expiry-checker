package mocks

import (
	context "context"
	"testing"

	types "github.com/babylonlabs-io/staking-expiry-checker/internal/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"
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

func Fuzz_Nosy_BtcInterface_GetBlockCount__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BtcInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBlockCount()
	})
}

func Fuzz_Nosy_BtcInterface_GetBlockTimestamp__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BtcInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var height uint64
		fill_err = tp.Fill(&height)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBlockTimestamp(height)
	})
}

func Fuzz_Nosy_BtcInterface_IsUTXOSpent__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *BtcInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var txid string
		fill_err = tp.Fill(&txid)
		if fill_err != nil {
			return
		}
		var vout uint32
		fill_err = tp.Fill(&vout)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.IsUTXOSpent(txid, vout)
	})
}

func Fuzz_Nosy_DbInterface_DeleteExpiredDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var id primitive.ObjectID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.DeleteExpiredDelegation(ctx, id)
	})
}

func Fuzz_Nosy_DbInterface_FindExpiredDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var btcTipHeight uint64
		fill_err = tp.Fill(&btcTipHeight)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.FindExpiredDelegations(ctx, btcTipHeight)
	})
}

func Fuzz_Nosy_DbInterface_GetBTCDelegationByStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBTCDelegationByStakingTxHash(ctx, stakingTxHash)
	})
}

func Fuzz_Nosy_DbInterface_GetBTCDelegationState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBTCDelegationState(ctx, stakingTxHash)
	})
}

func Fuzz_Nosy_DbInterface_GetBTCDelegationsByStates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		var states []types.DelegationState
		fill_err = tp.Fill(&states)
		if fill_err != nil {
			return
		}
		var paginationToken string
		fill_err = tp.Fill(&paginationToken)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.GetBTCDelegationsByStates(ctx, states, paginationToken)
	})
}

func Fuzz_Nosy_DbInterface_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.Ping(ctx)
	})
}

func Fuzz_Nosy_DbInterface_SaveTimeLockExpireCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
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
		var expireHeight uint64
		fill_err = tp.Fill(&expireHeight)
		if fill_err != nil {
			return
		}
		var txType string
		fill_err = tp.Fill(&txType)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.SaveTimeLockExpireCheck(ctx, stakingTxHashHex, expireHeight, txType)
	})
}

func Fuzz_Nosy_DbInterface_TransitionToUnbondedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
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
		var eligiblePreviousStates []types.DelegationState
		fill_err = tp.Fill(&eligiblePreviousStates)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.TransitionToUnbondedState(ctx, stakingTxHashHex, eligiblePreviousStates)
	})
}

func Fuzz_Nosy_DbInterface_TransitionToUnbondingState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
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
		var unbondingStartHeight uint64
		fill_err = tp.Fill(&unbondingStartHeight)
		if fill_err != nil {
			return
		}
		var unbondingTimelock uint64
		fill_err = tp.Fill(&unbondingTimelock)
		if fill_err != nil {
			return
		}
		var unbondingOutputIndex uint64
		fill_err = tp.Fill(&unbondingOutputIndex)
		if fill_err != nil {
			return
		}
		var unbondingTxHex string
		fill_err = tp.Fill(&unbondingTxHex)
		if fill_err != nil {
			return
		}
		var unbondingStartTimestamp int64
		fill_err = tp.Fill(&unbondingStartTimestamp)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.TransitionToUnbondingState(ctx, stakingTxHashHex, unbondingStartHeight, unbondingTimelock, unbondingOutputIndex, unbondingTxHex, unbondingStartTimestamp)
	})
}

func Fuzz_Nosy_DbInterface_TransitionToWithdrawnState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var _m *DbInterface
		fill_err = tp.Fill(&_m)
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
		var eligiblePreviousStates []types.DelegationState
		fill_err = tp.Fill(&eligiblePreviousStates)
		if fill_err != nil {
			return
		}
		if _m == nil {
			return
		}

		_m.TransitionToWithdrawnState(ctx, stakingTxHashHex, eligiblePreviousStates)
	})
}
