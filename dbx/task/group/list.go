package group

import (
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAM_GROUP_LIST        = "team/group/list"
	WORKER_TEAM_GROUP_MEMBER_LIST = "team/group/member/list"
)

type WorkerTeamGroupList struct {
	workflow.SimpleWorkerImpl
	ApiManagement *api.ApiContext
	NextTask      string
}

type ContentTeamGroupResult struct {
	GroupId             string `json:"group_id"`
	GroupName           string `json:"group_name"`
	GroupManagementType string `json:"group_management_type"`
	GroupExternalId     string `json:"group_external_id"`
	MemberCount         int64  `json:"member_count"`
}

type ContentTeamGroupList struct {
	Cursor string `json:"cursor"`
}

func (w *WorkerTeamGroupList) Prefix() string {
	return WORKER_TEAM_GROUP_LIST
}

func (w *WorkerTeamGroupList) Exec(task *workflow.Task) {
	tc := &ContentTeamGroupList{}
	workflow.UnmarshalContext(task, tc)

	if tc.Cursor == "" {
		w.list(task)
	} else {
		w.listContinue(tc.Cursor, task)
	}
}

func (w *WorkerTeamGroupList) list(task *workflow.Task) {
	type ListParam struct {
		Limit int `json:"limit"`
	}
	lp := ListParam{
		Limit: 100,
	}

	seelog.Debug("groups/list")
	cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/groups/list", lp)
	if !cont {
		return
	}
	w.processResult(res, task)
}

func (w *WorkerTeamGroupList) listContinue(cursor string, task *workflow.Task) {
	type ListContinueParam struct {
		Cursor string `json:"cursor"`
	}
	lp := ListContinueParam{
		Cursor: cursor,
	}

	seelog.Debugf("groups/list/continue (cursor: %s)", cursor)
	cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/groups/list/continue", lp)
	if !cont {
		return
	}

	w.processResult(res, task)
}

func (w *WorkerTeamGroupList) processResult(res *api.ApiRpcResponse, task *workflow.Task) {
	groups := gjson.Get(res.Body, "groups")
	if !groups.Exists() {
		seelog.Debugf("`groups` data not found")
		return
	}

	for _, group := range groups.Array() {
		groupIdJson := group.Get("group_id")

		if !groupIdJson.Exists() {
			seelog.Debugf("one of group info (group_id) not found `%s`", group.Raw)
			continue
		}

		c := ContentTeamGroupResult{
			GroupId:             groupIdJson.String(),
			GroupName:           group.Get("group_name").String(),
			GroupManagementType: group.Get("group_management_type.\\.tag").String(),
			GroupExternalId:     group.Get("group_external_id").String(),
			MemberCount:         group.Get("member_count").Int(),
		}

		w.Pipeline.Enqueue(
			workflow.MarshalTask(
				w.NextTask,
				groupIdJson.String(),
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
		c := ContentTeamGroupList{
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

type WorkerTeamGroupMemberList struct {
	workflow.SimpleWorkerImpl
	ApiManagement *api.ApiContext
	NextTask      string
}
