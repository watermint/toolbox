package rp_writer

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Writer interface {
	Name() string
	Row(r interface{})
	Open(ctl app_control.Control, model interface{}, opts ...rp_model.ReportOpt) error
	Close()
}
