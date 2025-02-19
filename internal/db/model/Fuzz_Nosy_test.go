package model

import (
	"testing"

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

func Fuzz_Nosy_BuildDelegationScanPaginationToken__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var d DelegationDocument
		fill_err = tp.Fill(&d)
		if fill_err != nil {
			return
		}

		BuildDelegationScanPaginationToken(d)
	})
}

func Fuzz_Nosy_DecodePaginationToken__(f *testing.F) {
	f.Fuzz(func(t *testing.T, token string) {
		DecodePaginationToken(token)
	})
}

// skipping Fuzz_Nosy_GetPaginationToken__ because parameters include func, chan, or unsupported interface: PaginationType
