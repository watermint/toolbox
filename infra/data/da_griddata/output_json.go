package da_griddata

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"io"
)

func NewJsonWriter() PlainGridDataWriter {
	return &jsonWriter{}
}

type jsonWriter struct {
}

func (z jsonWriter) FileSuffix() string {
	return ".json"
}

func (z jsonWriter) WriteRow(l esl.Logger, w io.Writer, formatter GridDataFormatter, row int, column []interface{}) error {
	colData := make([]interface{}, 0)
	for c := range column {
		colData = append(colData, formatter.Format(column[c], c, row))
	}
	jsonData, err := json.Marshal(colData)
	if err != nil {
		l.Debug("Unable to marshal data", esl.Error(err))
		return err
	}
	_, err = w.Write(jsonData)
	if err != nil {
		l.Debug("Unable to write a row", esl.Error(err))
		return err
	}
	_, err = w.Write([]byte("\n"))
	return err
}
