package number

import (
	"github.com/watermint/toolbox/essentials/lang"
)

type invalid struct {
}

func (z invalid) IsValid() bool {
	return false
}

func (z invalid) String() string {
	return "invalid"
}

func (z invalid) Pretty(l lang.Lang) string {
	return "invalid"
}

func (z invalid) IsNaN() bool {
	return false
}

func (z invalid) IsInt() bool {
	return false
}

func (z invalid) IsFloat() bool {
	return false
}

func (z invalid) Int() int {
	return 0
}

func (z invalid) Int8() int8 {
	return 0
}

func (z invalid) Int16() int16 {
	return 0
}

func (z invalid) Int32() int32 {
	return 0
}

func (z invalid) Int64() int64 {
	return 0
}

func (z invalid) Float32() float32 {
	return 0
}

func (z invalid) Float64() float64 {
	return 0
}
