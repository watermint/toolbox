package sv_conversation_history

import (
	"github.com/watermint/toolbox/domain/slack/api/work_client"
	"github.com/watermint/toolbox/domain/slack/api/work_pagination"
	"github.com/watermint/toolbox/domain/slack/model/mo_message"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type HistoryListOpts struct {
	Channel            string `url:"channel"`
	Cursor             string `url:"cursor,omitempty"`
	Before             string `url:"latest,omitempty"`
	After              string `url:"oldest,omitempty"`
	Inclusive          *bool  `url:"inclusive,omitempty"`
	IncludeAllMetadata *bool  `url:"include_all_metadata,omitempty"`
}

type History interface {
	ListEach(channel string, f func(m mo_message.Message), opts HistoryListOpts) error
}

func NewHistory(client work_client.Client) History {
	return &historyImpl{
		client: client,
	}
}

type historyImpl struct {
	client work_client.Client
}

func (z historyImpl) ListEach(channel string, f func(m mo_message.Message), opts HistoryListOpts) error {
	opts.Channel = channel
	pg := work_pagination.New(z.client).WithEndpoint("conversations.history").WithData(api_request.Query(&opts))
	return pg.OnData("channels", func(entry es_json.Json) error {
		m := &mo_message.Message{}
		if err := entry.Model(m); err != nil {
			return err
		}
		f(*m)
		return nil
	})
}
