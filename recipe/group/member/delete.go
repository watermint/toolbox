package member

import (
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Delete struct {
	Peer             rc_conn.ConnBusinessMgmt
	GroupName        string
	MemberEmail      string
	OperationLog     rp_model.TransactionReport
	ProgressDeleting app_msg.Message
}

func (z *Delete) Exec(c app_control.Control) error {
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	info := &UpdateInfo{
		GroupName:   z.GroupName,
		MemberEmail: z.MemberEmail,
	}

	c.UI().Info(z.ProgressDeleting.With("Group", z.GroupName).With("Email", z.MemberEmail))

	group, err := sv_group.New(z.Peer.Context()).ResolveByName(z.GroupName)
	if err != nil {
		z.OperationLog.Failure(err, info)
		return err
	}

	_, err = sv_group_member.New(z.Peer.Context(), group).Remove(sv_group_member.ByEmail(z.MemberEmail))
	if err != nil {
		z.OperationLog.Failure(err, info)
		return err
	}

	z.OperationLog.Success(info, group)

	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return qt_errors.ErrorScenarioTest
}

func (z *Delete) Preset() {
	z.OperationLog.SetModel(&UpdateInfo{}, &mo_group.Group{})
}
