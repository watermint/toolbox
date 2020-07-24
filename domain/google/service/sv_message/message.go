package sv_message

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/model/mo_message"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

const (
	FormatFull     = "full"
	FormatMetadata = "metadata"
	FormatMinimal  = "minimal"
	FormatRaw      = "raw"
)

type Message interface {
	Resolve(id string, opts ...ResolveOpt) (message *mo_message.Message, err error)
	List(q ...QueryOpt) (messages []*mo_message.Message, err error)
}

func New(ctx goog_context.Context, userId string) Message {
	return &msgImpl{
		ctx:    ctx,
		userId: userId,
	}
}

type QueryOpts struct {
	IncludeSpamTrash bool     `url:"includeSpamTrash,omitempty"`
	LabelIds         []string `url:"labelIds,omitempty"`
	MaxResults       int      `url:"maxResults,omitempty"`
	Query            string   `url:"q,omitempty"`
}

func (z QueryOpts) Apply(opts ...QueryOpt) QueryOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:]...)
	}
}

type QueryOpt func(o QueryOpts) QueryOpts

func IncludeSpamTrash(enabled bool) QueryOpt {
	return func(o QueryOpts) QueryOpts {
		o.IncludeSpamTrash = enabled
		return o
	}
}
func LabelIds(labels []string) QueryOpt {
	return func(o QueryOpts) QueryOpts {
		o.LabelIds = labels
		return o
	}
}
func MaxResults(v int) QueryOpt {
	return func(o QueryOpts) QueryOpts {
		o.MaxResults = v
		return o
	}
}
func Query(q string) QueryOpt {
	return func(o QueryOpts) QueryOpts {
		o.Query = q
		return o
	}
}

type ResolveOpts struct {
	Format string
}

func (z ResolveOpts) Apply(opts ...ResolveOpt) ResolveOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:]...)
	}
}

type ResolveOpt func(o ResolveOpts) ResolveOpts

func ResolveFormat(format string) ResolveOpt {
	return func(o ResolveOpts) ResolveOpts {
		o.Format = format
		return o
	}
}

type msgImpl struct {
	ctx    goog_context.Context
	userId string
}

func (z msgImpl) Resolve(id string, opts ...ResolveOpt) (message *mo_message.Message, err error) {
	ro := ResolveOpts{
		Format: "metadata",
	}.Apply(opts...)
	q := struct {
		Format string `url:"format"`
	}{
		Format: ro.Format,
	}

	res := z.ctx.Get("gmail/v1/users/"+z.userId+"/messages/"+id, api_request.Query(q))
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	message = &mo_message.Message{}
	if err := j.Model(message); err != nil {
		return nil, err
	} else {
		return message, nil
	}
}

func (z msgImpl) List(q ...QueryOpt) (messages []*mo_message.Message, err error) {
	p := QueryOpts{}.Apply(q...)
	res := z.ctx.Get("gmail/v1/users/"+z.userId+"/messages", api_request.Query(&p))
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	messages = make([]*mo_message.Message, 0)
	err = j.FindArrayEach("messages", func(e es_json.Json) error {
		m := &mo_message.Message{}
		if err := e.Model(m); err != nil {
			return err
		}
		messages = append(messages, m)
		return nil
	})
	return messages, err
}
