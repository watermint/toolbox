package dbx_sharing

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type SharedLink struct {
	SharedLinkId  string          `json:"shared_link_id"`
	AsMemberId    string          `json:"as_member_id"`
	AsMemberEmail string          `json:"as_member_email"`
	Link          json.RawMessage `json:"link"`
}

type SharedLinkList struct {
	Path          string
	AsMemberId    string
	AsMemberEmail string
	OnError       func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry       func(link *SharedLink) bool
}

func (a *SharedLinkList) List(c *dbx_api.Context) bool {
	type ListParam struct {
		Path string `json:"path,omitempty"`
	}
	lp := ListParam{
		Path: a.Path,
	}
	list := dbx_rpc.RpcList{
		EndpointList:         "sharing/list_shared_links",
		EndpointListContinue: "sharing/list_shared_links",
		AsMemberId:           a.AsMemberId,
		UseHasMore:           true,
		ResultTag:            "links",
		OnError:              a.OnError,
		OnEntry: func(link gjson.Result) bool {
			if a.OnEntry == nil {
				return true
			}

			linkId := link.Get("id").String()
			s := &SharedLink{
				SharedLinkId:  linkId,
				AsMemberId:    a.AsMemberId,
				AsMemberEmail: a.AsMemberEmail,
				Link:          json.RawMessage(link.Raw),
			}

			return a.OnEntry(s)
		},
	}

	return list.List(c, lp)
}
