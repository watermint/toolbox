package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Add struct {
	rc_recipe.RemarkIrreversible
	Peer           dbx_conn.ConnBusinessMgmt
	GroupName      string
	MemberEmail    string
	OperationLog   rp_model.TransactionReport
	ProgressAdding app_msg.Message
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	info := &UpdateInfo{
		GroupName:   z.GroupName,
		MemberEmail: z.MemberEmail,
	}

	c.UI().Info(z.ProgressAdding.With("Group", z.GroupName).With("Email", z.MemberEmail))

	group, err := sv_group.New(z.Peer.Context()).ResolveByName(z.GroupName)
	if err != nil {
		z.OperationLog.Failure(err, info)
		return err
	}

	_, err = sv_group_member.New(z.Peer.Context(), group).Add(sv_group_member.ByEmail(z.MemberEmail))
	if err != nil {
		z.OperationLog.Failure(err, info)
		return err
	}

	z.OperationLog.Success(info, group)

	return nil
}

func (z *Add) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.GroupName = "Marketing"
		m.MemberEmail = "john@example.com"
	})
	if err, _ = qt_errors.ErrorsForTest(c.Log(), err); err != nil && err != sv_group.ErrorGroupNotFoundForName {
		return err
	}
	return qt_errors.ErrorScenarioTest
}

func (z *Add) Preset() {
	z.OperationLog.SetModel(&UpdateInfo{}, &mo_group.Group{},
		rp_model.HiddenColumns(
			"result.group_id",
			"result.group_external_id",
		),
	)
}
