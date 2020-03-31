package rp_column_impl

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestColumnStream_Header(t *testing.T) {
	// simple struct
	{
		s := struct {
			SKU      string          `json:"sku"`
			Quantity int             `json:"quantity"`
			Notes    []string        `json:"notes"` // should be ignored
			Raw      json.RawMessage // should be ignored
		}{
			SKU:      "A123",
			Quantity: 87,
			Notes:    []string{"Export control passed"},
		}

		cs := NewStream(s)
		if !cmp.Equal(cs.Header(), []string{"sku", "quantity"}) {
			t.Error(cs.Header())
		}
	}
}

func TestColumnStream_Values(t *testing.T) {
	// simple struct
	{
		s := struct {
			SKU      string          `json:"sku"`
			Quantity int             `json:"quantity"`
			Notes    []string        `json:"notes"` // should be ignored
			Raw      json.RawMessage // should be ignored
		}{
			SKU:      "A123",
			Quantity: 87,
			Notes:    []string{"Export control passed"},
		}

		cs := NewStream(s)
		if !cmp.Equal(cs.Values(s), []interface{}{"A123", float64(87)}) {
			t.Error(cs.Values(s))
		}
		if !cmp.Equal(cs.ValueStrings(s), []string{"A123", "87"}) {
			t.Error(cs.ValueStrings(s))
		}
	}
}
