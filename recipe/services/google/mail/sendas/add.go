package sendas

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_sendas"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_sendas"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Add struct {
	Peer              goog_conn.ConnGoogleMail
	SendAs            rp_model.RowReport
	UserId            string
	Email             string
	ReplyTo           mo_string.OptionalString
	DisplayName       mo_string.OptionalString
	SkipVerify        bool
	ErrorVerification app_msg.Message
}

func (z *Add) Preset() {
	z.SendAs.SetModel(
		&mo_sendas.SendAs{},
		rp_model.HiddenColumns(),
	)
	z.Peer.SetScopes(
		goog_auth.ScopeGmailSettingsSharing,
	)
	z.UserId = "me"
}

func (z *Add) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.SendAs.Open(); err != nil {
		return err
	}

	opts := make([]sv_sendas.SendAsOpt, 0)
	if z.DisplayName.IsExists() {
		opts = append(opts, sv_sendas.DisplayName(z.DisplayName.Value()))
	}
	if z.ReplyTo.IsExists() {
		opts = append(opts, sv_sendas.ReplyTo(z.ReplyTo.Value()))
	}
	svs := sv_sendas.New(z.Peer.Client(), z.UserId)
	entry, err := svs.Add(z.Email, opts...)
	if err != nil {
		l.Debug("Unable to create send as", esl.Error(err))
		return err
	}

	z.SendAs.Row(entry)

	if z.SkipVerify {
		l.Debug("Skip verification")
		return nil
	}

	if err := svs.Verify(z.Email); err != nil {
		c.UI().Error(z.ErrorVerification.With("Error", err))
		return err
	}
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Email = "send_as@example.com"
	})
}
