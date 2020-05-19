package es_csv

import (
	"bytes"
	"encoding/csv"
)

func MakeCsv(f func(w *csv.Writer)) string {
	var buf bytes.Buffer
	cw := csv.NewWriter(&buf)
	f(cw)
	cw.Flush()
	return buf.String()
}
