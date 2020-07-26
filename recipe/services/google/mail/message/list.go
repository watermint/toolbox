package message

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
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

type List struct {
	Peer               goog_conn.ConnGoogleMail
	Messages           rp_model.RowReport
	UserId             string
	Labels             mo_string.OptionalString
	IncludeSpamTrash   bool
	Query              mo_string.OptionalString
	MaxResults         int
	Format             mo_string.SelectString
	ErrorLabelNotFound app_msg.Message
	ProgressGetMessage app_msg.Message
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailReadonly,
	)
	z.MaxResults = 20
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
	l := c.Log()
	ui := c.UI()

	svm := sv_message.New(z.Peer.Context(), z.UserId)
	queries := make([]sv_message.QueryOpt, 0)
	queries = append(queries, sv_message.IncludeSpamTrash(z.IncludeSpamTrash))
	queries = append(queries, sv_message.MaxResults(z.MaxResults))
	if z.Query.IsExists() {
		l.Debug("Build query param: query")
		queries = append(queries, sv_message.Query(z.Query.Value()))
	}
	if z.Labels.IsExists() {
		l.Debug("Build query param: labels")
		queryLabelNames := strings.Split(z.Labels.Value(), ",")
		queryLabelIds, err := sv_label.FindLabelIdsByNames(z.Peer.Context(), c.UI(), z.UserId, queryLabelNames)
		if err != nil {
			return err
		}
		queries = append(queries, sv_message.LabelIds(queryLabelIds))
	}

	messages, err := svm.List(queries...)
	if err != nil {
		return err
	}
	if err := z.Messages.Open(); err != nil {
		return err
	}
	for i, msgId := range messages {
		ui.Progress(z.ProgressGetMessage.With("Index", i+1))
		message, err := svm.Resolve(msgId.Id, sv_message.ResolveFormat(z.Format.Value()))
		if err != nil {
			return err
		}
		z.Messages.Row(message)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.ExecReplay(c, &List{}, "recipe-services-google-mail-message-list.json.gz", rc_recipe.NoCustomValues)
	if err != nil {
		return err
	}

	return rc_exec.ExecMock(c, &List{}, rc_recipe.NoCustomValues)
}
