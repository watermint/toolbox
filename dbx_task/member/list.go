package member

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
	"github.com/watermint/toolbox/workflow"
)

type WorkerTeamMemberList struct {
	workflow.SimpleWorkerImpl
	Api      *dbx_api.Context
	NextTask string

	IncludeRemoved bool
}

type ContextTeamMemberResult struct {
	Email        string          `json:"email"`
	TeamMemberId string          `json:"team_member_id"`
	Member       json.RawMessage `json:"member"`
}

func (w *WorkerTeamMemberList) Prefix() string {
	return "team/member/list"
}

func (w *WorkerTeamMemberList) Exec(task *workflow.Task) {
	type ListParam struct {
		IncludeRemoved bool `json:"include_removed"`
	}
	lp := ListParam{
		IncludeRemoved: w.IncludeRemoved,
	}

	list := dbx_rpc.RpcList{
		EndpointList:         "team/members/list",
		EndpointListContinue: "team/members/list/continue",
		UseHasMore:           true,
		ResultTag:            "members",
		HandlerError:         w.Pipeline.HandleGeneralFailure,
		HandlerEntry:         w.processResult,
	}

	list.List(w.Api, lp)
}

func (w *WorkerTeamMemberList) processResult(res gjson.Result) bool {
	emailJson := res.Get("profile.email")

	if !emailJson.Exists() {
		seelog.Debugf("one of member info (email) not found `%s`", res.Raw)
		return false
	}

	c := ContextTeamMemberResult{
		Email:        emailJson.String(),
		TeamMemberId: res.Get("profile.team_member_id").String(),
		Member:       json.RawMessage(res.Raw),
	}

	w.Pipeline.Enqueue(
		workflow.MarshalTask(
			w.NextTask,
			emailJson.String(),
			c,
		),
	)
	return true
}
