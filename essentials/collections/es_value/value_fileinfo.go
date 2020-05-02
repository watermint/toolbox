package es_value

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/collections/es_number"
	"os"
	"strings"
)

type valueFileInfo struct {
	v os.FileInfo
}

func (z valueFileInfo) IsNumber() bool {
	return false
}

func (z valueFileInfo) Compare(other Value) int {
	return strings.Compare(z.String(), other.String())
}

func (z valueFileInfo) String() string {
	return z.v.Name()
}

func (z valueFileInfo) AsNumber() es_number.Number {
	return es_number.New(z.v.Name())
}

func (z valueFileInfo) AsInterface() interface{} {
	return z.v
}

func (z valueFileInfo) Equals(other Value) bool {
	return z.Hash() == other.Hash()
}

func (z valueFileInfo) IsNull() bool {
	return false
}

func (z valueFileInfo) Hash() string {
	return fmt.Sprintf("%s-%t-%x-%d-%s", z.v.Name(), z.v.IsDir(), z.v.Mode(), z.v.Size(), z.v.ModTime().String())
}
