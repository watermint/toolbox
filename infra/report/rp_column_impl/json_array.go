package rp_column_impl

import (
	"encoding/json"
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/report/rp_column"
	"go.uber.org/zap"
)

func NewBson(header []string) rp_column.Column {
	return &JsonArray{header: header}
}

type JsonArray struct {
	header []string
}

func (z *JsonArray) Header() []string {
	return z.header
}

func (z *JsonArray) Values(r interface{}) (cols []interface{}) {
	l := app_root.Log()
	b := r.([]byte)
	if err := json.Unmarshal(b, &cols); err != nil {
		l.Error("Unable to unmarshal", zap.Error(err))
		return
	}
	return
}

func (z *JsonArray) ValueStrings(r interface{}) (cols []string) {
	l := app_root.Log()
	b := r.([]byte)
	rawCols := make([]interface{}, 0)
	if err := json.Unmarshal(b, &rawCols); err != nil {
		l.Error("Unable to unmarshal", zap.Error(err))
		return
	}
	cols = make([]string, 0)
	for _, c := range rawCols {
		cols = append(cols, fmt.Sprintf("%v", c))
	}
	return
}
