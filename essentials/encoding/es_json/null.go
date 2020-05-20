package es_json

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/collections/es_number"
)

func Null() Json {
	return &nullImpl{}
}

type nullImpl struct {
}

func (z nullImpl) RawString() string {
	return ""
}

func (z nullImpl) FindArrayEach(path string, f func(e Json) error) error {
	return ErrorNotFound
}

func (z nullImpl) Raw() json.RawMessage {
	return json.RawMessage("null")
}

func (z nullImpl) IsNull() bool {
	return true
}

func (z nullImpl) Array() (v []Json, t bool) {
	return nil, false
}

func (z nullImpl) ArrayEach(f func(e Json) error) error {
	return ErrorNotAnArray
}

func (z nullImpl) Object() (v map[string]Json, t bool) {
	return nil, false
}

func (z nullImpl) Bool() (v bool, t bool) {
	return false, false
}

func (z nullImpl) Number() (v es_number.Number, t bool) {
	return nil, false
}

func (z nullImpl) String() (v string, t bool) {
	return "", false
}

func (z nullImpl) Model(v interface{}) (err error) {
	return ErrorNotFound
}

func (z nullImpl) Find(path string) (j Json, found bool) {
	return nil, false
}

func (z nullImpl) FindModel(path string, v interface{}) (err error) {
	return ErrorNotFound
}

func (z nullImpl) FindArray(path string) (v []Json, t bool) {
	return nil, false
}

func (z nullImpl) FindObject(path string) (v map[string]Json, t bool) {
	return nil, false
}

func (z nullImpl) FindBool(path string) (v bool, t bool) {
	return false, false
}

func (z nullImpl) FindNumber(path string) (v es_number.Number, t bool) {
	return nil, false
}

func (z nullImpl) FindString(path string) (v string, t bool) {
	return "", false
}
