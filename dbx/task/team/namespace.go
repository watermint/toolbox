package team

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAM_NAMESPACE_LIST = "team/namespace/list"
)

type WorkerTeamNamespaceList struct {
	workflow.SimpleWorkerImpl
	Api      *api.ApiContext
	NextTask string
}

type ContextTeamNamespaceList struct {
	Cursor string `json:"cursor"`
}

type ContextTeamNamespaceListResult struct {
	NamespaceId   string          `json:"namespace_id"`
	NamespaceType string          `json:"namespace_type"`
	Name          string          `json:"name"`
	Namespace     json.RawMessage `json:"namespace"`
}

func (w *WorkerTeamNamespaceList) Prefix() string {
	return WORKER_TEAM_NAMESPACE_LIST
}

func (w *WorkerTeamNamespaceList) Exec(task *workflow.Task) {
	tc := &ContextTeamNamespaceList{}
	workflow.UnmarshalContext(task, tc)

	if tc.Cursor == "" {
		w.list(task)
	} else {
		w.listContinue(tc.Cursor, task)
	}
}

func (w *WorkerTeamNamespaceList) list(task *workflow.Task) {
	seelog.Info("Loading team namespace list")
	cont, res, _ := w.Pipeline.TaskRpc(task, w.Api, "team/namespaces/list", nil)
	if !cont {
		return
	}

	w.processResult(res, task)
}

func (w *WorkerTeamNamespaceList) listContinue(cursor string, task *workflow.Task) {
	type ListContinueParam struct {
		Cursor string `json:"cursor"`
	}
	lp := ListContinueParam{
		Cursor: cursor,
	}

	seelog.Debugf("team/namespaces/list/continue (cursor: %s)", cursor)
	cont, res, _ := w.Pipeline.TaskRpc(task, w.Api, "team/namespaces/list/continue", lp)
	if !cont {
		return
	}

	w.processResult(res, task)
}

func (w *WorkerTeamNamespaceList) processResult(res *api.ApiRpcResponse, task *workflow.Task) {
	namespaces := gjson.Get(res.Body, "namespaces")
	if !namespaces.Exists() {
		seelog.Debugf("`team_folders` data not found")
		return
	}

	for _, namespace := range namespaces.Array() {
		namespaceId := namespace.Get("namespace_id").String()

		c := ContextTeamNamespaceListResult{
			NamespaceId:   namespaceId,
			NamespaceType: namespace.Get("namespace_type.\\.tag").String(),
			Name:          namespace.Get("name").String(),
			Namespace:     json.RawMessage(namespace.Raw),
		}

		w.Pipeline.Enqueue(
			workflow.MarshalTask(
				w.NextTask,
				namespaceId,
				c,
			),
		)
	}

	hasMoreJson := gjson.Get(res.Body, "has_more")
	if hasMoreJson.Exists() && hasMoreJson.Bool() {
		cursorJson := gjson.Get(res.Body, "cursor")
		if !cursorJson.Exists() {
			seelog.Debug("Cursor not found in the response (has_more appear and true)")
			return
		}
		c := ContextTeamNamespaceList{
			Cursor: cursorJson.String(),
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
