package team

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
	"github.com/watermint/toolbox/workflow"
)

type WorkerTeamNamespaceList struct {
	workflow.SimpleWorkerImpl
	Api      *dbx_api.Context
	NextTask string
}

type ContextTeamNamespaceListResult struct {
	NamespaceId   string          `json:"namespace_id"`
	NamespaceType string          `json:"namespace_type"`
	Name          string          `json:"name"`
	Namespace     json.RawMessage `json:"namespace"`
}

func (w *WorkerTeamNamespaceList) Prefix() string {
	return "team/namespace/list"
}

func (w *WorkerTeamNamespaceList) Exec(task *workflow.Task) {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/namespaces/list",
		EndpointListContinue: "team/namespaces/list/continue",
		UseHasMore:           true,
		ResultTag:            "namespaces",
		HandlerError:         w.Pipeline.HandleGeneralFailure,
		HandlerEntry: func(namespace gjson.Result) bool {
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
			return true
		},
	}

	list.List(w.Api, nil)
}
