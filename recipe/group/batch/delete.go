package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type GroupName struct {
	Name string `json:"name"`
}

type Delete struct {
	rc_recipe.RemarkIrreversible
	ErrGroupNotFound    app_msg.Message
	ErrUnableToDelete   app_msg.Message
	File                fd_file.RowFeed
	OperationLog        rp_model.TransactionReport
	Peer                dbx_conn.ConnScopedTeam
	ProgressDeleteGroup app_msg.Message
}

func (z *Delete) Exec(c app_control.Control) error {
	ui := c.UI()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	svg := sv_group.New(z.Peer.Context())

	groups, err := svg.List()
	if err != nil {
		return err
	}
	groupByName := make(map[string]*mo_group.Group)
	for _, group := range groups {
		groupByName[group.GroupName] = group
	}

	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		r := m.(*GroupName)
		ui.Info(z.ProgressDeleteGroup.With("Name", r.Name))

		group, ok := groupByName[r.Name]
		if !ok {
			ui.Error(z.ErrGroupNotFound.With("Name", r.Name))
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
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		tf, err := qt_file.MakeTestFile("group-batch-delete", "Marketing\nSales\n")
		if err != nil {
			return
		}
		m := r.(*Delete)
		m.File.SetFilePath(tf)
	})
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsWrite,
	)
	z.File.SetModel(&GroupName{})
	z.OperationLog.SetModel(&GroupName{}, &mo_group.Group{},
		rp_model.HiddenColumns(
			"result.group_id",
			"result.group_external_id",
		),
	)
}
