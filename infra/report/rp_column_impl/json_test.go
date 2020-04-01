package rp_column_impl

import (
	"encoding/json"
	"github.com/google/go-cmp/cmp"
	"testing"
)

func TestHeaders(t *testing.T) {
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

		headers, err := Headers(s, func(name string) bool {
			return false
		})
		if err != nil {
			t.Error(err)
		}
		if !cmp.Equal(headers, []string{"sku", "quantity"}) {
			t.Error(headers)
		}
	}

	// nested struct
	{

		s := struct {
			SKU      string `json:"sku"`
			Quantity int    `json:"quantity"`
			Supplier struct {
				Name    string `json:"name"`
				Contact string `json:"contact"`
				Phone   string `json:"phone"`
			} `json:"supplier"`
		}{
			SKU:      "A123",
			Quantity: 87,
			Supplier: struct {
				Name    string `json:"name"`
				Contact string `json:"contact"`
				Phone   string `json:"phone"`
			}{
				Name:    "XYZ",
				Contact: "John",
				Phone:   "0123-456-789",
			},
		}

		headers, err := Headers(s, func(name string) bool {
			return name == "supplier.phone"
		})
		if err != nil {
			t.Error(err)
		}
		if !cmp.Equal(headers, []string{"sku", "quantity", "supplier.name", "supplier.contact"}) {
			t.Error(headers)
		}
	}

	// invalid type
	{
		a := make(chan int)
		_, err := Headers(a, func(name string) bool { return false })
		if err == nil {
			t.Error("invalid")
		}
	}
}

func TestParse(t *testing.T) {
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

		_, err := Parse(s)
		if err != nil {
			t.Error(err)
		}
	}
}
