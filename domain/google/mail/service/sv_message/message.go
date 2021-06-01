package sv_message

import (
	"github.com/watermint/toolbox/domain/google/api/goog_context"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_message"
	"github.com/watermint/toolbox/domain/google/mail/model/to_message"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_request"
	"github.com/watermint/toolbox/infra/ui/app_ui"
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
	Update(id string, opts ...UpdateOpt) (message *mo_message.Message, err error)
	Send(send to_message.Message) (sent *mo_message.Message, err error)
}

func New(ctx goog_context.Context, userId string) Message {
	return &msgImpl{
		ctx:    ctx,
		userId: userId,
	}
}

type UpdateOpts struct {
	Ids            []string `json:"ids,omitempty"`
	AddLabelIds    []string `json:"addLabelIds,omitempty"`
	RemoveLabelIds []string `json:"removeLabelIds,omitempty"`
}

func (z UpdateOpts) Apply(opts ...UpdateOpt) UpdateOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:]...)
	}
}

type UpdateOpt func(o UpdateOpts) UpdateOpts

func AddLabelIds(labels []string) UpdateOpt {
	return func(o UpdateOpts) UpdateOpts {
		o.AddLabelIds = labels
		return o
	}
}
func RemoveLabelIds(label []string) UpdateOpt {
	return func(o UpdateOpts) UpdateOpts {
		o.RemoveLabelIds = label
		return o
	}
}

type QueryOpts struct {
	NextPageToken    string   `url:"pageToken,omitempty"`
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

func (z msgImpl) Send(send to_message.Message) (sent *mo_message.Message, err error) {
	res := z.ctx.Post("gmail/v1/users/"+z.userId+"/messages/send", api_request.Param(&send))
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	sent = &mo_message.Message{}
	err = j.Model(sent)
	return sent, err
}

func (z msgImpl) Update(id string, opts ...UpdateOpt) (message *mo_message.Message, err error) {
	uo := UpdateOpts{}.Apply(opts...)
	res := z.ctx.Post("gmail/v1/users/"+z.userId+"/messages/"+id+"/modify", api_request.Param(&uo))
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	message = &mo_message.Message{}
	err = j.Model(message)
	return message, err
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

func (z msgImpl) listChunk(nextPageToken string, p QueryOpts) (messages []*mo_message.Message, newNextPageToken string, err error) {
	l := z.ctx.Log().With(esl.String("userId", z.userId), esl.String("nextPageToken", nextPageToken))
	if nextPageToken != "" {
		p.NextPageToken = nextPageToken
	}
	l.Debug("Execute query", esl.Any("query", p))
	res := z.ctx.Get("gmail/v1/users/"+z.userId+"/messages", api_request.Query(&p))
	if err, f := res.Failure(); f {
		return nil, "", err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, "", err
	}
	messages = make([]*mo_message.Message, 0)
	if _, found := j.Find("messages"); !found {
		l.Debug("No message found")
		return messages, "", nil
	}
	err = j.FindArrayEach("messages", func(e es_json.Json) error {
		m := &mo_message.Message{}
		if err := e.Model(m); err != nil {
			l.Debug("Unable to unmarshal", esl.Error(err))
			return err
		}
		messages = append(messages, m)
		return nil
	})
	if v, found := j.FindString("nextPageToken"); found {
		newNextPageToken = v
		l.Debug("nextPageToken found", esl.String("newNextPageToken", newNextPageToken))
	} else {
		l.Debug("nextPageToken NOT found")
	}
	return messages, newNextPageToken, err
}

func (z msgImpl) List(q ...QueryOpt) (messages []*mo_message.Message, err error) {
	opts := QueryOpts{}.Apply(q...)
	messages = make([]*mo_message.Message, 0)
	maxResults := opts.MaxResults
	var chunk []*mo_message.Message
	nextPageToken := ""
	for {
		chunk, nextPageToken, err = z.listChunk(nextPageToken, opts)
		if err != nil {
			return nil, err
		}
		messages = append(messages, chunk...)
		if nextPageToken == "" || (0 < maxResults && maxResults <= len(messages)) {
			return messages, nil
		}
		app_ui.ShowLongRunningProgress(z.ctx.UI(), z.userId, MProgress.ProgressRetrieve)
	}
}
