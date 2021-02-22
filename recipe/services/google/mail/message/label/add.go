package label

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_message"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_label"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_message"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
)

type Add struct {
	Peer                goog_conn.ConnGoogleMail
	UserId              string
	MessageId           string
	Label               string
	Message             rp_model.RowReport
	ErrorLabelsNotFound app_msg.Message
	AddLabelIfNotExist  bool
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailModify,
	)
	z.Message.SetModel(&mo_message.Message{})
	z.UserId = "me"
}

func (z *Add) Exec(c app_control.Control) error {
	labelNames := strings.Split(z.Label, ",")
	var labelIds []string
	var err error
	if z.AddLabelIfNotExist {
		labelIds, err = sv_label.FindOrAddLabelIdsByNames(z.Peer.Context(), c.UI(), z.UserId, labelNames)
	} else {
		labelIds, err = sv_label.FindLabelIdsByNames(z.Peer.Context(), c.UI(), z.UserId, labelNames)
	}
	if err := z.Message.Open(); err != nil {
		return err
	}

	message, err := sv_message.New(z.Peer.Context(), z.UserId).Update(z.MessageId, sv_message.AddLabelIds(labelIds))
	if err != nil {
		return err
	}
	z.Message.Row(message)
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &Add{}, "recipe-services-google-mail-filter-add.json.gz", func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.MessageId = "abc123def456"
		m.Label = "test"
	})
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.MessageId = "abc123def456"
		m.Label = "test"
	})
}
