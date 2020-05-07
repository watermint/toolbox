package es_value

import "github.com/watermint/toolbox/essentials/collections/es_number"

type valueNumber struct {
	v es_number.Number
}

func (z valueNumber) IsNumber() bool {
	return true
}

func (z valueNumber) Compare(other Value) int {
	on := other.AsNumber()
	return z.v.Compare(on)
}

func (z valueNumber) String() string {
	return z.v.String()
}

func (z valueNumber) AsNumber() es_number.Number {
	return z.v
}

func (z valueNumber) AsInterface() interface{} {
	return z.v
}

func (z valueNumber) Equals(other Value) bool {
	return z.Hash() == other.Hash()
}

func (z valueNumber) IsNull() bool {
	return false
}

func (z valueNumber) Hash() string {
	return z.v.String()
}
