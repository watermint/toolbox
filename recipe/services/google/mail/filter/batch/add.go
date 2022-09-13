package batch

import (
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_filter"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_message"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_filter"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_label"
	"github.com/watermint/toolbox/domain/google/mail/service/sv_message"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"regexp"
)

var (
	labelSplit = regexp.MustCompile(`[,;]`)
)

type AddRow struct {
	Query        string `json:"query"`
	AddLabels    string `json:"add_labels"`
	DeleteLabels string `json:"delete_labels"`
}

type Add struct {
	Peer                       goog_conn.ConnGoogleMail
	UserId                     string
	AddLabelIfNotExist         bool
	ApplyToExistingMessages    bool
	Filters                    rp_model.TransactionReport
	Messages                   rp_model.RowReport
	File                       fd_file.RowFeed
	ProgressCreateFilter       app_msg.Message
	ProgressUpdateMessage      app_msg.Message
	ErrorUnableToUpdateMessage app_msg.Message
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		goog_auth.ScopeGmailModify,
		goog_auth.ScopeGmailSettingsBasic,
	)
	z.UserId = "me"
	z.Filters.SetModel(&AddRow{}, &mo_filter.Filter{},
		rp_model.HiddenColumns(
			"result.criteria_from",
			"result.criteria_to",
			"result.criteria_subject",
			"result.criteria_negated_query",
		),
	)
	z.Messages.SetModel(&mo_message.Message{}, rp_model.NoConsoleOutput())
	z.File.SetModel(&AddRow{})
}

func (z *Add) labelIds(labelNames []string, c app_control.Control) (labelIds []string, err error) {
	l := c.Log()
	l.Debug("Looking for label Ids", esl.Strings("labelNames", labelNames))
	if z.AddLabelIfNotExist {
		labelIds, err = sv_label.FindOrAddLabelIdsByNames(z.Peer.Client(), c.UI(), z.UserId, labelNames)
	} else {
		labelIds, err = sv_label.FindLabelIdsByNames(z.Peer.Client(), c.UI(), z.UserId, labelNames)
	}
	return
}

func (z *Add) addDeleteLabels(row *AddRow, c app_control.Control) (addLabelIds, deleteLabelIds []string, err error) {
	l := c.Log().With(esl.Any("row", row))

	if row.AddLabels != "" {
		l.Debug("Add labels")
		names := labelSplit.Split(row.AddLabels, -1)
		addLabelIds, err = z.labelIds(names, c)
		if err != nil {
			l.Debug("Unable to identify labels", esl.Error(err))
			return nil, nil, err
		}
	}
	if row.DeleteLabels != "" {
		l.Debug("Delete labels")
		names := labelSplit.Split(row.DeleteLabels, -1)
		deleteLabelIds, err = z.labelIds(names, c)
		if err != nil {
			l.Debug("Unable to identify labels", esl.Error(err))
			return nil, nil, err
		}
	}
	return
}

func (z *Add) processMessages(query string, addLabelIds, deleteLabelIds []string, c app_control.Control) error {
	ui := c.UI()
	l := c.Log().With(esl.String("query", query), esl.Strings("addLabels", addLabelIds), esl.Strings("deleteLabels", deleteLabelIds))
	l.Debug("Retrieve existing messages that satisfies query")
	svm := sv_message.New(z.Peer.Client(), z.UserId)
	messages, err := svm.List(sv_message.Query(query))
	if err != nil {
		l.Debug("Unable to retrieve messages", esl.Error(err))
		return err
	}

	opts := make([]sv_message.UpdateOpt, 0)
	if addLabelIds != nil {
		opts = append(opts, sv_message.AddLabelIds(addLabelIds))
	}
	if deleteLabelIds != nil {
		opts = append(opts, sv_message.RemoveLabelIds(deleteLabelIds))
	}

	var lastErr error
	for i, message := range messages {
		ll := l.With(esl.Int("index", i), esl.String("messageId", message.Id))
		ui.Progress(z.ProgressUpdateMessage.With("Query", query).With("Index", i+1))
		ll.Debug("Process message")
		msg, err := svm.Update(message.Id, opts...)
		if err != nil {
			ui.Error(z.ErrorUnableToUpdateMessage.With("MessageId", message.Id).With("Error", err))
			ll.Debug("Unable to update message", esl.Error(err))
			lastErr = err
		} else {
			z.Messages.Row(msg)
		}
	}
	return lastErr
}

func (z *Add) addFilter(query string, addLabelIds, deleteLabelIds []string, c app_control.Control) (filter *mo_filter.Filter, err error) {
	l := c.Log().With(esl.String("query", query), esl.Strings("addLabel", addLabelIds), esl.Strings("deleteLabel", deleteLabelIds))
	l.Debug("Process label")
	opts := make([]sv_filter.Opt, 0)
	if addLabelIds != nil {
		opts = append(opts, sv_filter.AddLabelIds(addLabelIds))
	}
	if deleteLabelIds != nil {
		opts = append(opts, sv_filter.RemoveLabelIds(deleteLabelIds))
	}
	opts = append(opts, sv_filter.Query(query))

	svf := sv_filter.New(z.Peer.Client(), z.UserId)

	c.UI().Progress(z.ProgressCreateFilter.With("Query", query))
	return svf.Add(opts...)
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.Filters.Open(); err != nil {
		return err
	}
	if err := z.Messages.Open(); err != nil {
		return err
	}
	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		row := m.(*AddRow)
		addLabels, removeLabels, err := z.addDeleteLabels(row, c)
		if err != nil {
			z.Filters.Failure(err, row)
			return nil
		}

		filter, err := z.addFilter(row.Query, addLabels, removeLabels, c)
		if err != nil {
			z.Filters.Failure(err, row)
			return nil
		}
		z.Filters.Success(row, filter)

		if z.ApplyToExistingMessages {
			err = z.processMessages(row.Query, addLabels, removeLabels, c)
			if err != nil {
				z.Filters.Failure(err, row)
				return err
			}
		}
		return nil
	})
}

func (z *Add) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("batch-add.csv", `from:@google.com,services/google.com`)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecReplay(c, &Add{}, "recipe-services-google-mail-filter-batch-add.json.gz", func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.ApplyToExistingMessages = true
		m.AddLabelIfNotExist = true
		m.File.SetFilePath(f)
	})
}
