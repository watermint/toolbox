package mo_int

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
	rangeMin int64
	rangeMax int64
	value    int64
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
	return z.rangeMin, z.rangeMax
}

func (z *rangeInt) SetRange(rangeMin, rangeMax, preFill int64) {
	z.rangeMin = min(rangeMin, rangeMax)
	z.rangeMax = max(rangeMin, rangeMax)
	z.value = preFill
}

func (z *rangeInt) IsValid() bool {
	return z.rangeMin <= z.value && z.value <= z.rangeMax
}
