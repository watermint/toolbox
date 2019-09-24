package app_file

import (
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type ColDataFile interface {
	EachRow(ctl app_control.Control, exec func(cols []string, rowIndex int) error) error
}

type JsonDataFile interface {
	EachRow(ctl app_control.Control, exec func(j gjson.Result, rowIndex int) error) error
}
