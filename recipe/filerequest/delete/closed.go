package delete

import (
	"github.com/watermint/toolbox/domain/model/mo_filerequest"
	"github.com/watermint/toolbox/domain/service/sv_filerequest"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Closed struct {
	Peer    rc_conn.ConnUserFile
	Deleted rp_model.RowReport
}

func (z *Closed) Preset() {
	z.Deleted.SetModel(&mo_filerequest.FileRequest{})
}

func (z *Closed) Exec(c app_control.Control) error {
	if err := z.Deleted.Open(); err != nil {
		return err
	}
	frs, err := sv_filerequest.New(z.Peer.Context()).DeleteAllClosed()
	if err != nil {
		return err
	}
	for _, fr := range frs {
		z.Deleted.Row(fr)
	}
	return nil
}

func (z *Closed) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Closed{}, rc_recipe.NoCustomValues)
}
