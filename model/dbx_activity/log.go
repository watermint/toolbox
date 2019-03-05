package dbx_activity

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type ActivityLog struct {
	AccountId string                                        `json:"account_id,omitempty"`
	OnError   func(annotation dbx_api.ErrorAnnotation) bool `json:"-"`
	OnEvent   func(event Event) bool                        `json:"-"`
}

type Event struct {
	Raw json.RawMessage
}

func (z *ActivityLog) Events(c *dbx_api.Context) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team_log/get_events",
		EndpointListContinue: "team_log/get_events/continue",
		UseHasMore:           true,
		ResultTag:            "events",
		OnError:              z.OnError,
		OnEntry: func(event gjson.Result) bool {
			ev := Event{}
			err := c.ParseModel(&ev, event)
			if err != nil {
				return z.OnError(dbx_api.ErrorAnnotation{
					ErrorType: dbx_api.ErrorUnexpectedDataType,
					Error:     err,
				})
			}
			return z.OnEvent(ev)
		},
	}

	return list.List(c, z)
}
