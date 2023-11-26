package es_number_deprecated

import (
	"github.com/watermint/toolbox/essentials/go/es_lang"
	"golang.org/x/text/message"
	"math"
	"strconv"
)

type floatImpl struct {
	v float64
}

func (z floatImpl) Compare(other Number) int {
	if other.IsFloat() {
		o := other.Float64()
		switch {
		case z.v < o:
			return -1
		case z.v == o:
			return 0
		default:
			return 1
		}
	}
	o := float64(other.Int64())
	switch {
	case z.v < o:
		return -1
	case z.v == o:
		return 0
	default:
		return 1
	}
}

func (z floatImpl) IsValid() bool {
	return true
}

func (z floatImpl) IsInt() bool {
	return false
}

func (z floatImpl) IsFloat() bool {
	return true
}

func (z floatImpl) IsNaN() bool {
	return z.v == math.NaN()
}

func (z floatImpl) Int() int {
	return int(z.v)
}

func (z floatImpl) Int8() int8 {
	return int8(z.v)
}

func (z floatImpl) Int16() int16 {
	return int16(z.v)
}

func (z floatImpl) Int32() int32 {
	return int32(z.v)
}

func (z floatImpl) Int64() int64 {
	return int64(z.v)
}

func (z floatImpl) Float32() float32 {
	return float32(z.v)
}

func (z floatImpl) Float64() float64 {
	return z.v
}

func (z floatImpl) String() string {
	return strconv.FormatFloat(z.v, 'f', -1, 64)
}

func (z floatImpl) Pretty(l es_lang.Lang) string {
	p := message.NewPrinter(l.Tag())
	return p.Sprintf("%f", z.v)
}
