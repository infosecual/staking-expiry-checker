package utils

import (
	"testing"

	"github.com/babylonlabs-io/staking-expiry-checker/internal/types"
	"github.com/btcsuite/btcd/wire"
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

func Fuzz_Nosy_SupportedBtcNetwork_String__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var c SupportedBtcNetwork
		fill_err = tp.Fill(&c)
		if fill_err != nil {
			return
		}

		c.String()
	})
}

// skipping Fuzz_Nosy_Contains__ because parameters include func, chan, or unsupported interface: []T

func Fuzz_Nosy_GetFunctionName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, depth int) {
		GetFunctionName(depth)
	})
}

// skipping Fuzz_Nosy_PushOrQuit__ because parameters include func, chan, or unsupported interface: chan<- T

func Fuzz_Nosy_QualifiedStatesToUnbonded__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var unbondTxType types.StakingTxType
		fill_err = tp.Fill(&unbondTxType)
		if fill_err != nil {
			return
		}

		QualifiedStatesToUnbonded(unbondTxType)
	})
}

func Fuzz_Nosy_SerializeBtcTransaction__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var tx *wire.MsgTx
		fill_err = tp.Fill(&tx)
		if fill_err != nil {
			return
		}
		if tx == nil {
			return
		}

		SerializeBtcTransaction(tx)
	})
}

func Fuzz_Nosy_shortFuncName__(f *testing.F) {
	f.Fuzz(func(t *testing.T, fullName string) {
		shortFuncName(fullName)
	})
}
