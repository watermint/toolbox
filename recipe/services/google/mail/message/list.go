package message

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/model/mo_message"
	"github.com/watermint/toolbox/domain/google/service/sv_message"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer     goog_conn.ConnGoogleMail
	Messages rp_model.RowReport
	UserId   string
	Format   mo_string.SelectString
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailReadonly,
	)
	z.Messages.SetModel(&mo_message.Message{},
		rp_model.HiddenColumns(
			"id",
			"thread_id",
		),
	)
	z.Format.SetOptions(
		sv_message.FormatMetadata,
		sv_message.FormatFull, sv_message.FormatMetadata, sv_message.FormatMinimal, sv_message.FormatRaw,
	)
	z.UserId = "me"
}

func (z *List) Exec(c app_control.Control) error {
	svm := sv_message.New(z.Peer.Context(), z.UserId)
	messages, err := svm.List()
	if err != nil {
		return err
	}
	if err := z.Messages.Open(); err != nil {
		return err
	}
	for _, msgId := range messages {
		message, err := svm.Resolve(msgId.Id, sv_message.ResolveFormat(z.Format.Value()))
		if err != nil {
			return err
		}
		z.Messages.Row(message)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
