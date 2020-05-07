package es_json

import "github.com/tidwall/gjson"

func Validate(j []byte) bool {
	return gjson.ValidBytes(j)
}

func ValidateString(j string) bool {
	return gjson.Valid(j)
}
