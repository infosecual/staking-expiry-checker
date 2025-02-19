package db

import (
	"context"
	"testing"

	"github.com/babylonlabs-io/staking-expiry-checker/internal/config"
	"github.com/babylonlabs-io/staking-expiry-checker/internal/types"
	go_fuzz_utils "github.com/trailofbits/go-fuzz-utils"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func Fuzz_Nosy_Database_DeleteExpiredDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var id primitive.ObjectID
		fill_err = tp.Fill(&id)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.DeleteExpiredDelegation(c3, id)
	})
}

func Fuzz_Nosy_Database_FindExpiredDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var btcTipHeight uint64
		fill_err = tp.Fill(&btcTipHeight)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.FindExpiredDelegations(c3, btcTipHeight)
	})
}

func Fuzz_Nosy_Database_GetBTCDelegationByStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetBTCDelegationByStakingTxHash(c3, stakingTxHash)
	})
}

func Fuzz_Nosy_Database_GetBTCDelegationState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		var stakingTxHash string
		fill_err = tp.Fill(&stakingTxHash)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetBTCDelegationState(c3, stakingTxHash)
	})
}

func Fuzz_Nosy_Database_GetBTCDelegationsByStates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.GetBTCDelegationsByStates(c3, states, paginationToken)
	})
}

func Fuzz_Nosy_Database_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.Ping(c3)
	})
}

func Fuzz_Nosy_Database_SaveTimeLockExpireCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.SaveTimeLockExpireCheck(c3, stakingTxHashHex, expireHeight, txType)
	})
}

func Fuzz_Nosy_Database_Shutdown__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
		if fill_err != nil {
			return
		}
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.Shutdown(c3)
	})
}

func Fuzz_Nosy_Database_TransitionToUnbondedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.TransitionToUnbondedState(c3, stakingTxHashHex, eligiblePreviousStates)
	})
}

func Fuzz_Nosy_Database_TransitionToUnbondingState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.TransitionToUnbondingState(c3, stakingTxHashHex, unbondingStartHeight, unbondingTimelock, unbondingOutputIndex, unbondingTxHex, unbondingStartTimestamp)
	})
}

func Fuzz_Nosy_Database_TransitionToWithdrawnState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c1 context.Context
		fill_err = tp.Fill(&c1)
		if fill_err != nil {
			return
		}
		var cfg *config.DbConfig
		fill_err = tp.Fill(&cfg)
		if fill_err != nil {
			return
		}
		var c3 context.Context
		fill_err = tp.Fill(&c3)
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
		if cfg == nil {
			return
		}

		d1, err := New(c1, cfg)
		if err != nil {
			return
		}
		d1.TransitionToWithdrawnState(c3, stakingTxHashHex, eligiblePreviousStates)
	})
}

// skipping Fuzz_Nosy_Database_transitionState__ because parameters include func, chan, or unsupported interface: map[string]interface{}

func Fuzz_Nosy_DuplicateKeyError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *DuplicateKeyError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_InvalidPaginationTokenError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *InvalidPaginationTokenError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_NotFoundError_Error__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var e *NotFoundError
		fill_err = tp.Fill(&e)
		if fill_err != nil {
			return
		}
		if e == nil {
			return
		}

		e.Error()
	})
}

func Fuzz_Nosy_dbWithMetrics_DeleteExpiredDelegation__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.DeleteExpiredDelegation(ctx, id)
	})
}

func Fuzz_Nosy_dbWithMetrics_FindExpiredDelegations__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.FindExpiredDelegations(ctx, btcTipHeight)
	})
}

func Fuzz_Nosy_dbWithMetrics_GetBTCDelegationByStakingTxHash__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.GetBTCDelegationByStakingTxHash(ctx, stakingTxHash)
	})
}

func Fuzz_Nosy_dbWithMetrics_GetBTCDelegationState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.GetBTCDelegationState(ctx, stakingTxHash)
	})
}

func Fuzz_Nosy_dbWithMetrics_GetBTCDelegationsByStates__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.GetBTCDelegationsByStates(ctx, states, paginationToken)
	})
}

func Fuzz_Nosy_dbWithMetrics_Ping__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}
		var ctx context.Context
		fill_err = tp.Fill(&ctx)
		if fill_err != nil {
			return
		}
		if d == nil {
			return
		}

		d.Ping(ctx)
	})
}

func Fuzz_Nosy_dbWithMetrics_SaveTimeLockExpireCheck__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.SaveTimeLockExpireCheck(ctx, stakingTxHashHex, expireHeight, txType)
	})
}

func Fuzz_Nosy_dbWithMetrics_TransitionToUnbondedState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.TransitionToUnbondedState(ctx, stakingTxHashHex, eligiblePreviousStates)
	})
}

func Fuzz_Nosy_dbWithMetrics_TransitionToUnbondingState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.TransitionToUnbondingState(ctx, stakingTxHashHex, unbondingStartHeight, unbondingTimelock, unbondingOutputIndex, unbondingTxHex, unbondingStartTimestamp)
	})
}

func Fuzz_Nosy_dbWithMetrics_TransitionToWithdrawnState__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d *dbWithMetrics
		fill_err = tp.Fill(&d)
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
		if d == nil {
			return
		}

		d.TransitionToWithdrawnState(ctx, stakingTxHashHex, eligiblePreviousStates)
	})
}

// skipping Fuzz_Nosy_dbWithMetrics_run__ because parameters include func, chan, or unsupported interface: func() error

// skipping Fuzz_Nosy_DbInterface_DeleteExpiredDelegation__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_FindExpiredDelegations__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetBTCDelegationByStakingTxHash__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetBTCDelegationState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_GetBTCDelegationsByStates__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_Ping__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_SaveTimeLockExpireCheck__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_TransitionToUnbondedState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_TransitionToUnbondingState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_DbInterface_TransitionToWithdrawnState__ because parameters include func, chan, or unsupported interface: github.com/babylonlabs-io/staking-expiry-checker/internal/db.DbInterface

// skipping Fuzz_Nosy_IsDuplicateKeyError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_IsInvalidPaginationTokenError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_IsNotFoundError__ because parameters include func, chan, or unsupported interface: error

// skipping Fuzz_Nosy_runAndMeasureLatency__ because parameters include func, chan, or unsupported interface: func() (T, error)
