package message

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_message"
	"github.com/watermint/toolbox/domain/google/mail/model/to_message"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_message"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Send struct {
	rc_recipe.RemarkIrreversible
	rc_recipe.RemarkSecret
	Peer    goog_conn.ConnGoogleMail
	Sent    rp_model.RowReport
	UserId  string
	To      string
	Subject string
	Body    da_text.TextInput
}

func (z *Send) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailSend,
	)
	z.Sent.SetModel(&mo_message.Message{},
		rp_model.HiddenColumns(
			"id",
			"thread_id",
		),
	)
	z.UserId = "me"
}

func (z *Send) Exec(c app_control.Control) error {
	if err := z.Sent.Open(); err != nil {
		return err
	}

	msgPart := to_message.MessagePart{}
	msgPart = msgPart.WithTo(z.To)
	msgPart = msgPart.WithSubject(z.Subject)
	msgBody, err := z.Body.Content()
	if err != nil {
		return err
	}
	msgPart = msgPart.WithBodyText(string(msgBody))
	msg := to_message.Message{
		Payload: msgPart,
	}

	sent, err := sv_message.New(z.Peer.Context(), z.UserId).Send(msg)
	if err != nil {
		return err
	}
	z.Sent.Row(sent)
	return nil
}

func (z *Send) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("mail", "test mail content")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Send{}, func(r rc_recipe.Recipe) {
		m := r.(*Send)
		m.To = "test@example.com"
		m.Subject = "Test mail"
		m.Body.SetFilePath(f)
	})
}
