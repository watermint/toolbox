package es_jsonl

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestReadEachJson(t *testing.T) {
	{
		source := []string{
			`{"a":1}`,
			`{"a":2}`,
			`{"a":3}`,
			`{"a":4}`,
		}
		index := 1
		type SourceType struct {
			A int `json:"a"`
		}
		var obj SourceType
		err := ReadEachJson(strings.NewReader(strings.Join(source, "\n")), func(line []byte) error {
			if err := json.Unmarshal(line, &obj); err != nil {
				return err
			}
			if obj.A != index {
				t.Error(obj, index)
			}
			index++

			return nil
		})
		if index != 5 {
			t.Error(index)
		}
		if err != nil {
			t.Error(err)
		}
	}

	{
		source := `
{
"a":
1}`
		type SourceType struct {
			A int `json:"a"`
		}
		index := 1
		var obj SourceType
		err := ReadEachJson(strings.NewReader(source), func(line []byte) error {
			if err := json.Unmarshal(line, &obj); err != nil {
				return err
			}
			if obj.A != index {
				t.Error(obj.A, index)
			}
			index++
			return nil
		})
		if index != 2 {
			t.Error(index)
		}
		if err != nil {
			t.Error(err)
		}
	}
}
