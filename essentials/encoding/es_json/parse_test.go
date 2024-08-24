package es_json

import (
	"reflect"
	"testing"
)

func TestParseAny(t *testing.T) {
	patterns := []struct {
		data   []byte
		expect interface{}
	}{
		{
			data:   []byte(`true`),
			expect: true,
		},
		{
			data:   []byte(`false`),
			expect: false,
		},
		{
			data:   []byte(`1`),
			expect: 1,
		},
		{
			data:   []byte(`1.0`),
			expect: 1.0,
		},
		{
			data:   []byte(`"hello"`),
			expect: "hello",
		},
		{
			data:   []byte(`["a", "b", "c"]`),
			expect: []interface{}{"a", "b", "c"},
		},
		{
			data:   []byte(`{"a": "AAA", "b": "BBB"}`),
			expect: map[string]interface{}{"a": "AAA", "b": "BBB"},
		},
		{
			data:   []byte(`null`),
			expect: nil,
		},
	}

	for _, p := range patterns {
		if v, err := ParseAny(p.data); err != nil {
			t.Error(err)
		} else if !reflect.DeepEqual(v, p.expect) {
			t.Errorf("unexpected value: actual[%v], expected[%v]", v, p.expect)
		}
	}
}
