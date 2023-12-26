package label

import (
	"errors"
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_label"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Delete struct {
	Peer               goog_conn.ConnGoogleMail
	UserId             string
	Name               string
	ErrorLabelNotFound app_msg.Message
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailLabels,
	)
	z.UserId = "me"
}

func (z *Delete) Exec(c app_control.Control) error {
	l := c.Log()
	svl := sv_label.New(z.Peer.Client(), z.UserId)

	labels, err := svl.List()
	if err != nil {
		return err
	}
	for _, label := range labels {
		ll := l.With(esl.String("labelId", label.Id), esl.String("labelName", label.Name))
		ll.Debug("Scanning label")
		if label.Name == z.Name {
			ll.Debug("Target label found")
			err := svl.Remove(label.Id)
			return err
		}
	}
	c.UI().Error(z.ErrorLabelNotFound.With("Label", z.Name))
	return errors.New("label not found for the name")
}

func (z *Delete) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &Delete{}, "recipe-services-google-mail-label-delete.json.gz", func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Name = "xxxx"
	})
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.Name = "delete_test"
	})
}
