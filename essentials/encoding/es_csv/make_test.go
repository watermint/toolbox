package es_csv

import (
	"encoding/csv"
	"testing"
)

func TestMakeCsv(t *testing.T) {
	c := MakeCsv(func(w *csv.Writer) {
		err := w.Write([]string{"sku", "quantity"})
		if err != nil {
			t.Error(err)
		}
		err = w.Write([]string{"A123", "30"})
		if err != nil {
			t.Error(err)
		}
	})
	if c != "sku,quantity\nA123,30\n" {
		t.Error(c)
	}
}
