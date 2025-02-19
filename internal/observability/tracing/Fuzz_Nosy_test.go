package tracing

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

func Fuzz_Nosy_TracingInfo_addSpanDetail__(f *testing.F) {
	f.Fuzz(func(t *testing.T, data []byte) {

		tp, fill_err := GetTypeProvider(data)
		if fill_err != nil {
			return
		}
		var t1 *TracingInfo
		fill_err = tp.Fill(&t1)
		if fill_err != nil {
			return
		}
		var detail SpanDetail
		fill_err = tp.Fill(&detail)
		if fill_err != nil {
			return
		}
		if t1 == nil {
			return
		}

		t1.addSpanDetail(detail)
	})
}

// skipping Fuzz_Nosy_WrapWithSpan__ because parameters include func, chan, or unsupported interface: func() (Result, error)
