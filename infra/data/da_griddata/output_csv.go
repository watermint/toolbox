package da_griddata

import (
	"bytes"
	"encoding/csv"
	"fmt"
	"io"

	"github.com/watermint/toolbox/essentials/log/esl"
)

func NewCsvWriter() PlainGridDataWriter {
	return &csvWriter{}
}

type csvWriter struct {
}

func (z csvWriter) FileSuffix() string {
	return ".csv"
}

func (z csvWriter) WriteRow(l esl.Logger, w io.Writer, formatter GridDataFormatter, row int, column []interface{}) error {
	csvRow := make([]string, 0)
	for c := range column {
		v := formatter.Format(column[c], c, row)
		switch v0 := v.(type) {
		case string:
			csvRow = append(csvRow, v0)
		case int, int8, int16, int32, int64, uint, uint8, uint16, uint32, uint64:
			csvRow = append(csvRow, fmt.Sprintf("%d", v0))
		case float32, float64:
			csvRow = append(csvRow, fmt.Sprintf("%f", v0))
		default:
			csvRow = append(csvRow, fmt.Sprintf("%v", v0))
		}
	}
	csvData := &bytes.Buffer{}
	csvWriter := csv.NewWriter(csvData)
	_ = csvWriter.Write(csvRow)
	csvWriter.Flush()
	_, err := w.Write(csvData.Bytes())
	if err != nil {
		l.Debug("Unable to write a row", esl.Error(err))
		return err
	}
	// Option to write an additional newline is currently disabled.
	return nil
}
