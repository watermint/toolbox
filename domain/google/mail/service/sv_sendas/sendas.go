package sv_sendas

import (
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_sendas"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type SendAs interface {
	List() (sa []*mo_sendas.SendAs, err error)
	Add(sendAsEmail string, opts ...SendAsOpt) (sa *mo_sendas.SendAs, err error)
	Verify(sendAsEmail string) (err error)
	Remove(sendAsEmail string) (err error)
}

type SendAsOpts struct {
	SendAsEmail    string `json:"sendAsEmail"`
	DisplayName    string `json:"displayName,omitempty"`
	ReplyToAddress string `json:"replyToAddress,omitempty"`
	Signature      string `json:"signature,omitempty"`
}

func (z SendAsOpts) Apply(opts []SendAsOpt) SendAsOpts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

type SendAsOpt func(o SendAsOpts) SendAsOpts

func ReplyTo(replyTo string) SendAsOpt {
	return func(o SendAsOpts) SendAsOpts {
		o.ReplyToAddress = replyTo
		return o
	}
}

func DisplayName(displayName string) SendAsOpt {
	return func(o SendAsOpts) SendAsOpts {
		o.DisplayName = displayName
		return o
	}
}

func New(ctx goog_client.Client, userId string) SendAs {
	return &sendAsImpl{
		ctx:    ctx,
		userId: userId,
	}
}

type sendAsImpl struct {
	ctx    goog_client.Client
	userId string
}

func (z sendAsImpl) List() (sa []*mo_sendas.SendAs, err error) {
	res := z.ctx.Get("gmail/v1/users/" + z.userId + "/settings/sendAs")
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	sa = make([]*mo_sendas.SendAs, 0)
	if _, found := j.Find("sendAs"); !found {
		return sa, nil
	}
	err = j.FindArrayEach("sendAs", func(e es_json.Json) error {
		m := &mo_sendas.SendAs{}
		if err := e.Model(m); err != nil {
			return err
		}
		sa = append(sa, m)
		return nil
	})
	return sa, err
}

func (z sendAsImpl) Add(sendAsEmail string, opts ...SendAsOpt) (sa *mo_sendas.SendAs, err error) {
	p := SendAsOpts{
		SendAsEmail: sendAsEmail,
	}.Apply(opts)

	res := z.ctx.Post("gmail/v1/users/"+z.userId+"/settings/sendAs", api_request.Param(&p))
	if err, f := res.Failure(); f {
		return nil, err
	}
	sa = &mo_sendas.SendAs{}
	err = res.Success().Json().Model(sa)
	return
}

func (z sendAsImpl) Verify(sendAsEmail string) (err error) {
	res := z.ctx.Post("gmail/v1/users/" + z.userId + "/settings/sendAs/" + sendAsEmail + "/verify")
	if err, f := res.Failure(); f {
		return err
	}
	return nil
}

func (z sendAsImpl) Remove(sendAsEmail string) (err error) {
	res := z.ctx.Delete("gmail/v1/users/" + z.userId + "/settings/sendAs/" + sendAsEmail)
	if err, f := res.Failure(); f {
		return err
	}
	return nil
}
