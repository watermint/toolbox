package app_file

import "github.com/watermint/toolbox/experimental/app_control"

type RowDataFile interface {
	EachRow(ctl app_control.Control, exec func(cols []string, rowIndex int) error) error
}
