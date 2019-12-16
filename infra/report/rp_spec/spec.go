package rp_spec

import (
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type ReportSpec interface {
	Name() string
	Row() interface{}
	Desc() app_msg.Message
	Columns() []string
	ColumnDesc(col string) app_msg.Message
	Options() []rp_model.ReportOpt
	Open(opts ...rp_model.ReportOpt) (rp_model.Report, error)
}
