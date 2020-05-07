package es_value

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"os"
)

type Value interface {
	fmt.Stringer
	AsNumber() es_number.Number
	AsInterface() interface{}
	Equals(other Value) bool
	Compare(other Value) int
	IsNull() bool
	IsNumber() bool
	Hash() string
}

func New(v interface{}) Value {
	if v == nil {
		return &valueNull{}
	}

	switch w := v.(type) {
	case Value:
		return w
	case string:
		return &valueString{v: w}
	case es_number.Number:
		return &valueNumber{v: w}
	case int, int8, int16, int32, int64:
		return &valueNumber{v: es_number.New(w)}
	case uint, uint8, uint16, uint32, uint64:
		return &valueNumber{v: es_number.New(w)}
	case float32, float64:
		return &valueNumber{v: es_number.New(w)}
	case os.FileInfo:
		return &valueFileInfo{v: w}
	default:
		return &valueOther{v: v}
	}
}
