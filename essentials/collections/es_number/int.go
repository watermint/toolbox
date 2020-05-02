package es_number

import (
	"github.com/watermint/toolbox/essentials/lang"
	"golang.org/x/text/message"
	"strconv"
)

type intImpl struct {
	v int64
}

func (z intImpl) Compare(other Number) int {
	if other.IsInt() {
		o := other.Int64()
		switch {
		case z.v < o:
			return 1
		case z.v == o:
			return 0
		default:
			return -1
		}
	}
	o := int64(other.Float64())
	switch {
	case z.v < o:
		return 1
	case z.v == o:
		return 0
	default:
		return -1
	}
}

func (z intImpl) IsValid() bool {
	return true
}

func (z intImpl) IsInt() bool {
	return true
}

func (z intImpl) IsBigInt() bool {
	return false
}

func (z intImpl) IsFloat() bool {
	return false
}

func (z intImpl) IsBigFloat() bool {
	return false
}

func (z intImpl) IsNaN() bool {
	return false
}

func (z intImpl) Float32() float32 {
	return float32(z.v)
}

func (z intImpl) Float64() float64 {
	return float64(z.v)
}

func (z intImpl) Int() int {
	return int(z.v)
}

func (z intImpl) Int8() int8 {
	return int8(z.v)
}

func (z intImpl) Int16() int16 {
	return int16(z.v)
}

func (z intImpl) Int32() int32 {
	return int32(z.v)
}

func (z intImpl) Int64() int64 {
	return z.v
}

func (z intImpl) String() string {
	return strconv.FormatInt(z.v, 10)
}

func (z intImpl) Pretty(l lang.Lang) string {
	p := message.NewPrinter(l.Tag())
	return p.Sprintf("%d", z.v)
}
