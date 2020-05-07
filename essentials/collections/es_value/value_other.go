package es_value

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"strings"
)

type valueOther struct {
	v interface{}
}

func (z valueOther) Compare(other Value) int {
	return strings.Compare(z.String(), other.String())
}

func (z valueOther) IsNumber() bool {
	return false
}

func (z valueOther) String() string {
	return fmt.Sprintf("%v", z.v)
}

func (z valueOther) AsNumber() es_number.Number {
	return es_number.New(z.v)
}

func (z valueOther) AsInterface() interface{} {
	return z.v
}

func (z valueOther) Equals(other Value) bool {
	return z.Hash() == other.Hash()
}

func (z valueOther) IsNull() bool {
	return z.v == nil
}

func (z valueOther) Hash() string {
	return z.String()
}
