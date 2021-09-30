package batch

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/essentials/lang"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type Add struct {
	Peer              dbx_conn.ConnScopedTeam
	File              fd_file.RowFeed
	ManagementType    mo_string.SelectString
	OperationLog      rp_model.TransactionReport
	SkipAlreadyExists app_msg.Message
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsWrite,
	)
	z.File.SetModel(&GroupName{})
	z.ManagementType.SetOptions(
		"company_managed",
		"company_managed", "user_managed",
	)
	z.OperationLog.SetModel(&GroupName{}, &mo_group.Group{},
		rp_model.HiddenColumns(
			"result.group_id",
			"result.group_external_id",
		),
	)
}

func (z *Add) addGroup(gn *GroupName, c app_control.Control) error {
	group, err := sv_group.New(z.Peer.Context()).Create(
		gn.Name,
		sv_group.ManagementType(z.ManagementType.Value()),
	)
	de := dbx_error.NewErrors(err)
	switch {
	case de == nil:
		z.OperationLog.Success(gn, group)
		return nil

	case de.IsGroupNameAlreadyUsed():
		z.OperationLog.Skip(z.SkipAlreadyExists, gn)
		return nil

	default:
		z.OperationLog.Failure(err, gn)
		return err
	}
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	var lastErr, queueErr error

	c.Sequence().Do(func(s eq_sequence.Stage) {
		s.Define("add_group", z.addGroup, c)
		q := s.Get("add_group")

		queueErr = z.File.EachRow(func(m interface{}, rowIndex int) error {
			q.Enqueue(m)
			return nil
		})
	}, eq_sequence.ErrorHandler(func(err error, mouldId, batchId string, p interface{}) {
		lastErr = err
	}))

	return lang.NewMultiErrorOrNull(lastErr, queueErr)
}

func (z *Add) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		tf, err := qt_file.MakeTestFile("group-batch-add", "Marketing\nSales\n")
		if err != nil {
			return
		}
		m := r.(*Add)
		m.File.SetFilePath(tf)
	})
}
