package api_parser

import (
	"encoding/json"
	"testing"
)

func TestParseModel(t *testing.T) {
	type T struct {
		Id     string `path:"id"`
		Name   string `path:"name"`
		Volume int    `path:"volume"`
	}

	{
		v := &T{}
		j := `{"id":"abc123","name":"earth","volume":1234}`

		if err := ParseModelString(v, j); err != nil || v.Id != "abc123" || v.Name != "earth" || v.Volume != 1234 {
			t.Error("invalid")
		}
		if err := ParseModelRaw(v, json.RawMessage(j)); err != nil || v.Id != "abc123" || v.Name != "earth" || v.Volume != 1234 {
			t.Error("invalid")
		}
	}

	{ // broken json
		v := &T{}
		j := `{"id":"abc123","name":"earth","volume":1234,}`

		if err := ParseModelString(v, j); err == nil {
			t.Error("invalid")
		}
		if err := ParseModelRaw(v, json.RawMessage(j)); err == nil {
			t.Error("invalid")
		}
	}
}
