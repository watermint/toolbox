package sv_label

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"strings"
)

type MsgFindLabel struct {
	ErrorLabelNotFound  app_msg.Message
	ErrorFailedAddLabel app_msg.Message
}

var (
	MFindLabel = app_msg.Apply(&MsgFindLabel{}).(*MsgFindLabel)
)

func FindLabelIdsByNames(ctx goog_context.Context, ui app_ui.UI, userId string, names []string) (labelIds []string, err error) {
	l := ctx.Log()
	l.Debug("Build query param: labels")
	queryLabelIds := make([]string, 0)
	l.Debug("Search for labels", esl.Strings("labels", names))
	queryLabels, missing, err := NewCached(ctx, userId).ResolveByNames(names)
	if err != nil {
		ui.Error(MFindLabel.ErrorLabelNotFound.With("Label", strings.Join(missing, ", ")))
		return nil, err
	}
	for _, q := range queryLabels {
		queryLabelIds = append(queryLabelIds, q.Id)
	}
	return queryLabelIds, nil
}

func FindOrAddLabelIdsByNames(ctx goog_context.Context, ui app_ui.UI, userId string, names []string) (labelIds []string, err error) {
	l := ctx.Log()
	l.Debug("Build query param: labels")
	queryLabelIds := make([]string, 0)
	l.Debug("Search for labels", esl.Strings("labels", names))
	svl := NewCached(ctx, userId)
	labels, err := svl.List()
	if err != nil {
		return nil, err
	}

	for _, name := range names {
		l.Debug("Looking for a label", esl.String("name", name))
		found := false
		for _, label := range labels {
			if label.Name == name {
				found = true
				queryLabelIds = append(queryLabelIds, label.Id)
				break
			}
		}
		if !found {
			label, err := svl.Add(name)
			if err != nil {
				ui.Error(MFindLabel.ErrorFailedAddLabel.With("Error", err).With("Label", name))
				return nil, err
			}
			queryLabelIds = append(queryLabelIds, label.Id)
		}
	}

	return queryLabelIds, nil
}
