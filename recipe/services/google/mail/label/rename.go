package label

import (
	"errors"
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_label"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_label"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Rename struct {
	Peer               goog_conn.ConnGoogleMail
	CurrentName        string
	NewName            string
	UserId             string
	Label              rp_model.RowReport
	ErrorLabelNotFound app_msg.Message
}

func (z *Rename) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailLabels,
	)
	z.Label.SetModel(&mo_label.Label{})
	z.UserId = "me"
}

func (z *Rename) Exec(c app_control.Control) error {
	l := c.Log()
	svl := sv_label.New(z.Peer.Client(), z.UserId)

	labels, err := svl.List()
	if err != nil {
		return err
	}
	if err := z.Label.Open(); err != nil {
		return err
	}
	for _, label := range labels {
		ll := l.With(esl.String("labelId", label.Id), esl.String("labelName", label.Name))
		ll.Debug("Scanning label")
		if label.Name == z.CurrentName {
			ll.Debug("Target label found")
			label, err := svl.Update(label.Id, sv_label.Name(z.NewName))
			if err != nil {
				return err
			}
			z.Label.Row(label)
			return nil
		}
	}
	c.UI().Error(z.ErrorLabelNotFound.With("Label", z.CurrentName))
	return errors.New("label not found for the name")
}

func (z *Rename) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &Rename{}, "recipe-services-google-mail-label-rename.json.gz", func(r rc_recipe.Recipe) {
		m := r.(*Rename)
		m.CurrentName = "xxxx"
		m.NewName = "test-new"
	})
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Rename{}, func(r rc_recipe.Recipe) {
		m := r.(*Rename)
		m.CurrentName = "test"
		m.NewName = "test-new"
	})
}
