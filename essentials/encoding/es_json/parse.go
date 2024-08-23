package es_json

import (
	"encoding/json"
	"strings"
)

func ParseAny(data []byte) (interface{}, error) {
	if strings.TrimSpace(string(data)) == "null" {
		return nil, nil
	}
	var vb bool
	if err := json.Unmarshal(data, &vb); err == nil {
		return vb, nil
	}
	var vi int
	if err := json.Unmarshal(data, &vi); err == nil {
		return vi, nil
	}
	var vf float64
	if err := json.Unmarshal(data, &vf); err == nil {
		return vf, nil
	}
	var vs string
	if err := json.Unmarshal(data, &vs); err == nil {
		return vs, nil
	}
	var va []interface{}
	if err := json.Unmarshal(data, &va); err == nil {
		return va, nil
	}
	var vo map[string]interface{}
	if err := json.Unmarshal(data, &vo); err == nil {
		return vo, nil
	}
	return nil, ErrorInvalidJSONFormat
}
