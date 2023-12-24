package label

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_label"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_label"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Add struct {
	Peer                  goog_conn.ConnGoogleMail
	Name                  string
	UserId                string
	LabelListVisibility   mo_string.SelectString
	MessageListVisibility mo_string.SelectString
	ColorBackground       mo_string.SelectString
	ColorText             mo_string.SelectString
	Label                 rp_model.RowReport
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailLabels,
	)
	z.Label.SetModel(&mo_label.Label{})
	z.UserId = "me"
	z.LabelListVisibility.SetOptions(sv_label.VisibilityLabelListShow, sv_label.VisibilityLabelList...)
	z.MessageListVisibility.SetOptions(sv_label.VisibilityMessageListShow, sv_label.VisibilityMessageList...)
	z.ColorText.SetOptions("", sv_label.ValidColors...)
	z.ColorBackground.SetOptions("", sv_label.ValidColors...)
}

func (z *Add) Exec(c app_control.Control) error {
	svl := sv_label.New(z.Peer.Client(), z.UserId)
	label, err := svl.Add(z.Name,
		sv_label.LabelListVisibility(z.LabelListVisibility.Value()),
		sv_label.MessageListVisibility(z.MessageListVisibility.Value()),
		sv_label.ColorBackground(z.ColorBackground.Value()),
		sv_label.ColorText(z.ColorText.Value()),
	)
	if err != nil {
		return err
	}
	if err := z.Label.Open(); err != nil {
		return err
	}
	z.Label.Row(label)
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &Add{}, "recipe-services-google-mail-label-add.json.gz", func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Name = "test"
	})
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Name = "test"
	})
}
