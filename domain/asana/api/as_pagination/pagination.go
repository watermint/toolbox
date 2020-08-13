package as_pagination

import (
	"github.com/watermint/toolbox/domain/asana/api/as_context"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_request"
)

const (
	defaultLimit = 20
)

type Pagination interface {
	WithEndpoint(endpoint string) Pagination
	WithData(d ...api_request.RequestDatum) Pagination
	OnData(handler func(entry es_json.Json) error) error
}

func New(ctx as_context.Context) Pagination {
	return NewWithLimit(ctx, defaultLimit)
}

func NewWithLimit(ctx as_context.Context, limit int) Pagination {
	return &pgImpl{
		ctx:   ctx,
		limit: limit,
	}
}

type pgImpl struct {
	ctx      as_context.Context
	limit    int
	data     []api_request.RequestDatum
	endpoint string
}

func (z pgImpl) WithEndpoint(endpoint string) Pagination {
	z.endpoint = endpoint
	return z
}

func (z pgImpl) WithData(d ...api_request.RequestDatum) Pagination {
	z.data = d
	return z
}

func (z pgImpl) OnData(handler func(entry es_json.Json) error) error {
	l := z.ctx.Log().With(esl.String("endpoint", z.endpoint), esl.Int("limit", z.limit))

	offset := ""
	for {
		ll := l.With(esl.String("offset", offset))
		ll.Debug("Paginated request")
		res := z.ctx.GetWithPagination(z.endpoint, offset, z.limit, z.data...)
		if err, fail := res.Failure(); fail {
			ll.Debug("Failure response", esl.Error(err))
			return err
		}

		rj := res.Success().Json()

		if entries, found := rj.FindArray("data"); found {
			ll.Debug("Data found")
			for _, entry := range entries {
				ll.Debug("Process entry", esl.Any("entry", entry))
				if err := handler(entry); err != nil {
					ll.Debug("Handler returned an error", esl.Error(err))
					return err
				}
			}
		} else {
			ll.Debug("Data not found")
		}

		if np, found := rj.Find("next_page"); !found {
			ll.Debug("Next page not found")
			return nil
		} else {
			ll.Debug("Next page found", esl.Any("nextPage", np))
			if offset, found = np.FindString("offset"); found {
				ll.Debug("Offset found", esl.String("offset", offset))
				continue
			} else {
				ll.Debug("Offset not found")
				return nil
			}
		}
	}
}
