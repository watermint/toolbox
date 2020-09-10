package mo_int

import (
	"github.com/watermint/toolbox/essentials/collections/es_number"
)

type RangeInt interface {
	Value() int
	Value64() int64
	Range() (min, max int64)
	SetRange(min, max, preFill int64)
	SetValue(value int64)
	IsValid() bool
}

func NewRange() RangeInt {
	return &rangeInt{}
}

type rangeInt struct {
	min   int64
	max   int64
	value int64
}

func (z *rangeInt) SetValue(value int64) {
	z.value = value
}

func (z *rangeInt) Value() int {
	return int(z.value)
}

func (z *rangeInt) Value64() int64 {
	return z.value
}

func (z *rangeInt) Range() (min, max int64) {
	return z.min, z.max
}

func (z *rangeInt) SetRange(min, max, preFill int64) {
	z.min = es_number.Min(min, max).Int64()
	z.max = es_number.Max(min, max).Int64()
	z.value = preFill
}

func (z *rangeInt) IsValid() bool {
	return z.min <= z.value && z.value <= z.max
}
