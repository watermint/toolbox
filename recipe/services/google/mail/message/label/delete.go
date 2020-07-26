package label

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/model/mo_message"
	"github.com/watermint/toolbox/domain/google/service/sv_label"
	"github.com/watermint/toolbox/domain/google/service/sv_message"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
)

type Delete struct {
	Peer                goog_conn.ConnGoogleMail
	UserId              string
	MessageId           string
	Label               string
	Message             rp_model.RowReport
	ErrorLabelsNotFound app_msg.Message
}

func (z *Delete) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailModify,
	)
	z.Message.SetModel(&mo_message.Message{})
	z.UserId = "me"
}

func (z *Delete) Exec(c app_control.Control) error {
	labelNames := strings.Split(z.Label, ",")
	labels, missing, err := sv_label.NewCached(z.Peer.Context(), z.UserId).ResolveByNames(labelNames)
	if err != nil {
		c.UI().Error(z.ErrorLabelsNotFound.With("Labels", strings.Join(missing, ",")))
		return err
	}
	labelIds := make([]string, 0)
	for _, label := range labels {
		labelIds = append(labelIds, label.Id)
	}
	if err := z.Message.Open(); err != nil {
		return err
	}

	message, err := sv_message.New(z.Peer.Context(), z.UserId).Update(z.MessageId, sv_message.RemoveLabelIds(labelIds))
	if err != nil {
		return err
	}
	z.Message.Row(message)
	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.MessageId = "abc123def456"
		m.Label = "test"
	})
}
