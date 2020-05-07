package es_value

import (
	"github.com/watermint/toolbox/essentials/collections/es_number"
)

func Null() Value {
	return &valueNull{}
}

type valueNull struct {
}

func (z valueNull) IsNumber() bool {
	return false
}

func (z valueNull) Compare(other Value) int {
	if other.IsNull() {
		return 0
	}
	return 1
}

func (z valueNull) String() string {
	return ""
}

func (z valueNull) AsNumber() es_number.Number {
	return es_number.Zero()
}

func (z valueNull) AsInterface() interface{} {
	return nil
}

func (z valueNull) Equals(other Value) bool {
	return other.IsNull()
}

func (z valueNull) IsNull() bool {
	return true
}

func (z valueNull) Hash() string {
	return ""
}
