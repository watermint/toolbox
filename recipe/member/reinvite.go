package member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Reinvite struct {
	Peer             dbx_conn.ConnBusinessMgmt
	Silent           bool
	OperationLog     rp_model.TransactionReport
	ProgressReinvite app_msg.Message
}

func (z *Reinvite) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()

	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	for _, member := range members {
		ll := l.With(es_log.Any("member", member))
		if member.Status != "invited" {
			ll.Debug("Skip")
			continue
		}

		ui.Info(z.ProgressReinvite.With("MemberEmail", member.Email))
		if err = sv_member.New(z.Peer.Context()).Remove(member); err != nil {
			ll.Debug("Unable to remove", es_log.Error(err))
			z.OperationLog.Failure(err, member)
			continue
		}
		opts := make([]sv_member.AddOpt, 0)
		if z.Silent {
			opts = append(opts, sv_member.AddWithoutSendWelcomeEmail())
		}
		invite, err := sv_member.New(z.Peer.Context()).Add(member.Email, opts...)
		if err != nil {
			ll.Debug("Unable to invite", es_log.Error(err))
			z.OperationLog.Failure(err, member)
			continue
		}

		z.OperationLog.Success(member, invite)
	}
	return nil
}

func (z *Reinvite) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Reinvite{}, func(r rc_recipe.Recipe) {
		m := r.(*Reinvite)
		m.Silent = true
	})
	if e, _ := qt_errors.ErrorsForTest(c.Log(), err); e != nil {
		return e
	}
	return qt_errors.ErrorHumanInteractionRequired
}

func (z *Reinvite) Preset() {
	z.OperationLog.SetModel(
		&mo_member.Member{},
		&mo_member.Member{},
		rp_model.HiddenColumns(
			"input.team_member_id",
			"input.familiar_name",
			"input.abbreviated_name",
			"input.member_folder_id",
			"input.external_id",
			"input.account_id",
			"input.persistent_id",
			"result.team_member_id",
			"result.familiar_name",
			"result.abbreviated_name",
			"result.member_folder_id",
			"result.external_id",
			"result.account_id",
			"result.persistent_id",
		),
	)
}
