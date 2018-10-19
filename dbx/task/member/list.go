package member

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAM_MEMBER_LIST = "team/member/list"
)

type WorkerTeamMemberList struct {
	workflow.SimpleWorkerImpl
	ApiManagement *api.ApiContext
	NextTask      string

	IncludeRemoved bool
}

type ContextTeamMemberResult struct {
	Email        string          `json:"email"`
	TeamMemberId string          `json:"team_member_id"`
	Member       json.RawMessage `json:"member"`
}

type ContextTeamMemberList struct {
	Cursor string `json:"cursor"`
}

func (w *WorkerTeamMemberList) Prefix() string {
	return WORKER_TEAM_MEMBER_LIST
}

func (w *WorkerTeamMemberList) Exec(task *workflow.Task) {
	tc := &ContextTeamMemberList{}
	workflow.UnmarshalContext(task, tc)

	if tc.Cursor == "" {
		w.list(task)
	} else {
		w.listContinue(tc.Cursor, task)
	}
}

func (w *WorkerTeamMemberList) list(task *workflow.Task) {
	type ListParam struct {
		IncludeRemoved bool `json:"include_removed"`
	}
	lp := ListParam{
		IncludeRemoved: w.IncludeRemoved,
	}

	seelog.Debug("members/list")
	seelog.Info("Loading team member list")
	cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/members/list", lp)
	if !cont {
		return
	}

	w.processResult(res, task)
}

func (w *WorkerTeamMemberList) listContinue(cursor string, task *workflow.Task) {
	type ListContinueParam struct {
		Cursor string `json:"cursor"`
	}
	lp := ListContinueParam{
		Cursor: cursor,
	}

	seelog.Debugf("members/list/continue (cursor: %s)", cursor)
	cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/members/list/continue", lp)
	if !cont {
		return
	}

	w.processResult(res, task)
}

func (w *WorkerTeamMemberList) processResult(res *api.ApiRpcResponse, task *workflow.Task) {
	members := gjson.Get(res.Body, "members")
	if !members.Exists() {
		seelog.Debugf("`members` data not found")
		return
	}

	for _, member := range members.Array() {
		emailJson := member.Get("profile.email")

		if !emailJson.Exists() {
			seelog.Debugf("one of member info (email) not found `%s`", member.Raw)
			continue
		}

		c := ContextTeamMemberResult{
			Email:        emailJson.String(),
			TeamMemberId: member.Get("profile.team_member_id").String(),
			Member:       json.RawMessage(member.Raw),
		}

		w.Pipeline.Enqueue(
			workflow.MarshalTask(
				w.NextTask,
				emailJson.String(),
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
		c := ContextTeamMemberList{
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
