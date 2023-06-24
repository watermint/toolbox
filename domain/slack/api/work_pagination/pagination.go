package work_pagination

import (
	"errors"
	"github.com/watermint/toolbox/domain/slack/api/work_client"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
)

func New(ctx work_client.Client) Pagination {
	return &pgImpl{
		ctx: ctx,
	}
}

type Pagination interface {
	WithEndpoint(endpoint string) Pagination
	WithData(d ...api_request.RequestDatum) Pagination
	OnData(path string, handler func(entry es_json.Json) error) error
}

type pgImpl struct {
	ctx      work_client.Client
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

func (z pgImpl) OnData(path string, handler func(entry es_json.Json) error) error {
	l := z.ctx.Log().With(esl.String("endpoint", z.endpoint), esl.String("path", path))

	cursor := ""
	for {
		ll := l.With(esl.String("cursor", cursor))
		ll.Debug("Paginated request")

		type PaginatedParam struct {
			Cursor string `url:"cursor,omitempty" json:"cursor,omitempty"`
		}
		var pp = &PaginatedParam{
			Cursor: cursor,
		}
		data := make([]api_request.RequestDatum, 0)
		data = append(data, z.data...)
		data = append(data, api_request.Query(pp))
		res := z.ctx.Get(z.endpoint, data...)
		if err, fail := res.Failure(); fail {
			ll.Debug("Failure response", esl.Error(err))
			return err
		}

		rj := res.Success().Json()

		if ok, found := rj.FindBool("ok"); found {
			if !ok {
				errorMsg := "error"
				errorMsg0, found := rj.FindString("error")
				if found {
					errorMsg = errorMsg0
				}
				err := errors.New(errorMsg)
				ll.Debug("Error", esl.Error(err))
				return err
			}
		}

		if entries, found := rj.FindArray(path); found {
			ll.Debug("Data found")
			for _, entry := range entries {
				ll.Debug("Process entry", esl.Any("entry", entry))
				if err := handler(entry); err != nil {
					ll.Debug("Handler returned an error", esl.Error(err))
					return err
				}
			}
		} else if entry, found := rj.Find(path); found {
			ll.Debug("Single data found", esl.Any("entry", entry))
			if err := handler(entry); err != nil {
				ll.Debug("Handler returned an error", esl.Error(err))
				return err
			}
		} else {
			ll.Debug("Data not found")
		}

		var found bool
		if cursor, found = rj.FindString("response_metadata.next_cursor"); !found || cursor == "" {
			ll.Debug("Next page not found")
			return nil
		} else {
			ll.Debug("Next page found", esl.Any("cursor", cursor))
		}
	}
}
