package rp_spec

import (
	"github.com/watermint/toolbox/infra/report/rp_model"
)

// Deprecated:
type ReportSpec interface {
	rp_model.Spec

	// Deprecated:
	Open(opts ...rp_model.ReportOpt) (rp_model.SideCarReport, error)
}
