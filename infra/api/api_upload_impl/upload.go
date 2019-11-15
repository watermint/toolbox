package api_upload_impl

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_rpc"
	"github.com/watermint/toolbox/infra/api/api_upload"
	"io"
)

func New(ctx api_context.Context, endpoint string) api_upload.Upload {
	return &Upload{
		ctx:      ctx,
		endpoint: endpoint,
	}
}

type Upload struct {
	ctx      api_context.Context
	p        interface{}
	r        io.Reader
	endpoint string
}

func (z *Upload) Param(p interface{}) api_upload.Upload {
	return &Upload{
		ctx:      z.ctx,
		p:        p,
		r:        z.r,
		endpoint: z.endpoint,
	}
}

func (z *Upload) Content(r io.Reader) api_upload.Upload {
	return &Upload{
		ctx:      z.ctx,
		p:        z.p,
		r:        r,
		endpoint: z.endpoint,
	}
}

func (z *Upload) Call() (res api_rpc.Response, err error) {
	return z.ctx.Request(z.endpoint).Param(z.p).Upload(z.r)
}
