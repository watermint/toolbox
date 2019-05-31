package app_flow

import "github.com/watermint/toolbox/atbx/app_control"

func IsErrorPrefix(prefix string, err error) bool {
	return false
}

type RowDataFile interface {
	EachRow(ctl app_control.Control, exec func(cols []string, rowIndex int) error) error
}
