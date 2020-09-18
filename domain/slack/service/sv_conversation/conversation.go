package sv_conversation

import (
	"github.com/watermint/toolbox/domain/slack/api/work_context"
	"github.com/watermint/toolbox/domain/slack/api/work_pagination"
	"github.com/watermint/toolbox/domain/slack/model/mo_conversation"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Conversation interface {
	ListEach(h func(c *mo_conversation.Conversation)) error
}

func New(ctx work_context.Context) Conversation {
	return &conImpl{
		ctx: ctx,
	}
}

type conImpl struct {
	ctx work_context.Context
}

func (z conImpl) ListEach(h func(c *mo_conversation.Conversation)) error {
	pg := work_pagination.New(z.ctx).WithEndpoint("conversations.list")
	return pg.OnData("channels", func(entry es_json.Json) error {
		c := &mo_conversation.Conversation{}
		if err := entry.Model(c); err != nil {
			return err
		}
		h(c)
		return nil
	})
}
