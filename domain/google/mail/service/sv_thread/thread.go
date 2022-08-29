package sv_thread

import (
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/domain/google/mail/model/mo_thread"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/api/api_request"
)

const (
	FormatFull     = "full"
	FormatMetadata = "metadata"
	FormatMinimal  = "minimal"
)

type Thread interface {
	List() (threads []*mo_thread.Thread, err error)
	Resolve(id string, opts ...ResolveOpt) (thread *mo_thread.Thread, err error)
}

type ResolveOpts struct {
	Format string
}

type ResolveOpt func(o ResolveOpts) ResolveOpts

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

func ResolveFormat(format string) ResolveOpt {
	return func(o ResolveOpts) ResolveOpts {
		o.Format = format
		return o
	}
}

func New(ctx goog_client.Client, userId string) Thread {
	return &threadImpl{
		ctx:    ctx,
		userId: userId,
	}
}

type threadImpl struct {
	ctx    goog_client.Client
	userId string
}

func (z threadImpl) List() (threads []*mo_thread.Thread, err error) {
	res := z.ctx.Get("gmail/v1/users/" + z.userId + "/threads")
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	threads = make([]*mo_thread.Thread, 0)
	if _, found := j.Find("threads"); !found {
		return threads, nil
	}
	err = j.FindArrayEach("threads", func(e es_json.Json) error {
		m := &mo_thread.Thread{}
		if err := e.Model(m); err != nil {
			return err
		}
		threads = append(threads, m)
		return nil
	})
	return threads, err
}

func (z threadImpl) Resolve(id string, opts ...ResolveOpt) (thread *mo_thread.Thread, err error) {
	ro := ResolveOpts{
		Format: "metadata",
	}.Apply(opts...)
	q := struct {
		Format string `url:"format"`
	}{
		Format: ro.Format,
	}

	res := z.ctx.Get("gmail/v1/users/"+z.userId+"/threads/"+id, api_request.Query(q))
	if err, f := res.Failure(); f {
		return nil, err
	}
	j, err := res.Success().AsJson()
	if err != nil {
		return nil, err
	}
	thread = &mo_thread.Thread{}
	if err := j.Model(thread); err != nil {
		return nil, err
	} else {
		return thread, nil
	}
}
