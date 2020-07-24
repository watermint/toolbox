package sv_label

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"strings"
)

type MsgFindLabel struct {
	ErrorLabelNotFound app_msg.Message
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
