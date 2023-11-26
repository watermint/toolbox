package es_value_deprecated

import (
	"github.com/watermint/toolbox/essentials/collections/es_number_deprecated"
	"strings"
)

type valueString struct {
	v string
}

func (z valueString) IsNumber() bool {
	return false
}

func (z valueString) Compare(other Value) int {
	return strings.Compare(z.String(), other.String())
}

func (z valueString) String() string {
	return z.v
}

func (z valueString) AsNumber() es_number_deprecated.Number {
	return es_number_deprecated.New(z.v)
}

func (z valueString) AsInterface() interface{} {
	return z.v
}

func (z valueString) Equals(other Value) bool {
	return other.String() == z.v
}

func (z valueString) IsNull() bool {
	return false
}

func (z valueString) Hash() string {
	return z.v
}
