package sharedlink

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
	"github.com/watermint/toolbox/workflow"
)

type WorkerSharedLinkList struct {
	workflow.SimpleWorkerImpl
	Api      *dbx_api.Context
	NextTask string
}

type ContextSharedLink struct {
	AsMemberId    string `json:"as_member_id"`
	AsMemberEmail string `json:"as_member_email"`
	Cursor        string `json:"cursor"`
	Path          string `json:"path"`
}

type ContextSharedLinkResult struct {
	SharedLinkId  string          `json:"shared_link_id"`
	AsMemberId    string          `json:"as_member_id"`
	AsMemberEmail string          `json:"as_member_email"`
	Link          json.RawMessage `json:"link"`
}

func (w *WorkerSharedLinkList) Prefix() string {
	return "sharedlink/list"
}

func (w *WorkerSharedLinkList) Exec(task *workflow.Task) {
	tc := &ContextSharedLink{}
	workflow.UnmarshalContext(task, tc)

	type ListParam struct {
		Path   string `json:"path,omitempty"`
		Cursor string `json:"cursor,omitempty"`
	}
	lp := ListParam{
		Path: tc.Path,
	}
	list := dbx_rpc.RpcList{
		EndpointList:         "sharing/list_shared_links",
		EndpointListContinue: "sharing/list_shared_links",
		AsMemberId:           tc.AsMemberId,
		UseHasMore:           true,
		ResultTag:            "links",
		HandlerError:         w.Pipeline.HandleGeneralFailure,
		HandlerEntry: func(link gjson.Result) bool {
			linkId := link.Get("id").String()

			c := ContextSharedLinkResult{
				SharedLinkId:  linkId,
				AsMemberId:    tc.AsMemberId,
				AsMemberEmail: tc.AsMemberEmail,
				Link:          json.RawMessage(link.Raw),
			}

			w.Pipeline.Enqueue(
				workflow.MarshalTask(
					w.NextTask,
					linkId,
					c,
				),
			)
			return true
		},
	}

	list.List(w.Api, lp)
}
