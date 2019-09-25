package app_file

import "github.com/watermint/toolbox/infra/control/app_control"

type Data interface {
	Model(ctl app_control.Control, m interface{}) error
	EachRow(exec func(m interface{}, rowIndex int) error) error
}
