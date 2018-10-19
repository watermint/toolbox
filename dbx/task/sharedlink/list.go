package sharedlink

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_SHAREDLINK_LIST = "sharedlink/list"
)

type WorkerSharedLinkList struct {
	workflow.SimpleWorkerImpl
	Api      *api.ApiContext
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
	return WORKER_SHAREDLINK_LIST
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

	cont, res, _ := w.Pipeline.TaskRpcAsMemberId(task, w.Api, "sharing/list_shared_links", lp, tc.AsMemberId)
	if !cont {
		return
	}
	w.processResult(tc, res, task)
}

func (w *WorkerSharedLinkList) processResult(tc *ContextSharedLink, res *api.ApiRpcResponse, task *workflow.Task) {
	for _, link := range gjson.Get(res.Body, "links").Array() {
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
	}

	hasMoreJson := gjson.Get(res.Body, "has_more")
	if hasMoreJson.Exists() && hasMoreJson.Bool() {
		cursorJson := gjson.Get(res.Body, "cursor")

		c := ContextSharedLink{
			Cursor:        cursorJson.String(),
			AsMemberId:    tc.AsMemberId,
			AsMemberEmail: tc.AsMemberEmail,
		}

		w.Pipeline.Enqueue(
			workflow.MarshalTask(
				w.Prefix(),
				cursorJson.String(),
				c,
			),
		)
	}
}
