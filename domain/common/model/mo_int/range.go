package mo_int

import (
	"github.com/watermint/toolbox/infra/util/ut_math"
)

type RangeInt interface {
	Value() int
	Range() (min, max int)
	SetRange(min, max, preFill int)
	SetValue(value int)
	IsValid() bool
}

func NewRange() RangeInt {
	return &rangeInt{}
}

type rangeInt struct {
	min   int
	max   int
	value int
}

func (z *rangeInt) SetValue(value int) {
	z.value = value
}

func (z *rangeInt) Value() int {
	return z.value
}

func (z *rangeInt) Range() (min, max int) {
	return z.min, z.max
}

func (z *rangeInt) SetRange(min, max, preFill int) {
	z.min = ut_math.MinInt(min, max)
	z.max = ut_math.MaxInt(min, max)
	z.value = preFill
}

func (z *rangeInt) IsValid() bool {
	return z.min <= z.value && z.value <= z.max
}
