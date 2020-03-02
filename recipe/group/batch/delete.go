package batch

import (
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type GroupName struct {
	Name string `json:"name"`
}

type Delete struct {
	ErrGroupNotFound    app_msg.Message
	ErrUnableToDelete   app_msg.Message
	File                fd_file.RowFeed
	OperationLog        rp_model.TransactionReport
	Peer                rc_conn.ConnBusinessMgmt
	ProgressDeleteGroup app_msg.Message
}

func (z *Delete) Exec(c app_control.Control) error {
	ui := c.UI()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	svg := sv_group.New(z.Peer.Context())

	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		r := m.(*GroupName)
		ui.Info(z.ProgressDeleteGroup.With("Name", r.Name))

		group, err := svg.ResolveByName(r.Name)
		if err != nil {
			ui.Error(z.ErrGroupNotFound.With("Name", r.Name).With("Error", err.Error()))
			z.OperationLog.Failure(err, r)
			return nil
		}

		if err = svg.Remove(group.GroupId); err != nil {
			ui.Error(z.ErrUnableToDelete.With("Name", r.Name).With("Error", err.Error()))
			z.OperationLog.Failure(err, r)
			return nil
		}
		z.OperationLog.Success(r, group)
		return nil
	})
}

func (z *Delete) Test(c app_control.Control) error {
	return qt_errors.ErrorImplementMe
}

func (z *Delete) Preset() {
	z.File.SetModel(&GroupName{})
	z.OperationLog.SetModel(&GroupName{}, &mo_group.Group{})
}
