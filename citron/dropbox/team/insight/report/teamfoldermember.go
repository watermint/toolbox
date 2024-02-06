package report

import (
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_insight_reports"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Teamfoldermember struct {
	Database string
	Entry    rp_model.RowReport
}

func (z *Teamfoldermember) Preset() {
	z.Entry.SetModel(&uc_insight_reports.TeamFolderMember{})
}

func (z *Teamfoldermember) Exec(c app_control.Control) error {
	return nil
}

func (z *Teamfoldermember) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
