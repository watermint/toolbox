package number

import (
	"github.com/watermint/toolbox/essentials/lang"
	"strconv"
	"strings"
)

type Format interface {
	String() string
	Pretty(l lang.Lang) string
}

type Precision interface {
	IsValid() bool
	IsNaN() bool
	IsInt() bool
	IsFloat() bool
}

type Integer interface {
	Format
	Precision
	Int() int
	Int8() int8
	Int16() int16
	Int32() int32
	Int64() int64
}

type Float interface {
	Format
	Precision
	Float32() float32
	Float64() float64
}

type Number interface {
	Format
	Integer
	Float
}

func Zero() Number {
	return &intImpl{v: 0}
}

func New(v interface{}) Number {
	switch x := v.(type) {
	case Number:
		return x
	case int:
		return &intImpl{v: int64(x)}
	case int8:
		return &intImpl{v: int64(x)}
	case int16:
		return &intImpl{v: int64(x)}
	case int32:
		return &intImpl{v: int64(x)}
	case int64:
		return &intImpl{v: x}
	case uint8:
		return &intImpl{v: int64(x)}
	case uint16:
		return &intImpl{v: int64(x)}
	case uint32:
		return &intImpl{v: int64(x)}
	case uint64: // convert to float
		return &floatImpl{v: float64(x)}
	case float32:
		return &floatImpl{v: float64(x)}
	case float64:
		return &floatImpl{v: x}
	case string:
		x1 := strings.TrimSpace(x)
		vi, err := strconv.ParseInt(x1, 10, 64)
		if err == nil {
			return &intImpl{v: vi}
		}
		vf, err := strconv.ParseFloat(x1, 64)
		if err == nil {
			return &floatImpl{v: vf}
		}
		return &invalid{}

	default:
		return &invalid{}
	}
}
