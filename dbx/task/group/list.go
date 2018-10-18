package group

import (
	"encoding/json"
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
	Api      *api.ApiContext
	NextTask string
}

type ContextTeamGroupResult struct {
	GroupId             string `json:"group_id"`
	GroupName           string `json:"group_name"`
	GroupManagementType string `json:"group_management_type"`
	GroupExternalId     string `json:"group_external_id"`
	MemberCount         int64  `json:"member_count"`
}

type ContextTeamGroupList struct {
	Cursor string `json:"cursor"`
}

func (w *WorkerTeamGroupList) Prefix() string {
	return WORKER_TEAM_GROUP_LIST
}

func (w *WorkerTeamGroupList) Exec(task *workflow.Task) {
	tc := &ContextTeamGroupList{}
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
	cont, res, _ := w.Pipeline.TaskRpc(task, w.Api, "team/groups/list", lp)
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
	cont, res, _ := w.Pipeline.TaskRpc(task, w.Api, "team/groups/list/continue", lp)
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

		c := ContextTeamGroupResult{
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
		c := ContextTeamGroupList{
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

type ContextTeamGroupMemberList struct {
	GroupId      string          `json:"group_id"`
	GroupName    string          `json:"group_name"`
	TeamMemberId string          `json:"team_member_id"`
	Member       json.RawMessage `json:"member"`
}

func (w *WorkerTeamGroupMemberList) Prefix() string {
	return WORKER_TEAM_GROUP_MEMBER_LIST
}

func (w *WorkerTeamGroupMemberList) Exec(task *workflow.Task) {
	tc := &ContextTeamGroupResult{}
	workflow.UnmarshalContext(task, tc)

	type GroupSelector struct {
		Tag     string `json:".tag"`
		GroupId string `json:"group_id"`
	}
	type ListParam struct {
		Group GroupSelector `json:"group"`
		Limit int           `json:"limit"`
	}

	lp := ListParam{
		Group: GroupSelector{
			Tag:     "group_id",
			GroupId: tc.GroupId,
		},
		Limit: 3,
	}
	seelog.Debug("groups/members/list")
	cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/groups/members/list", lp)
	if !cont {
		return
	}

	w.processResult(res, tc)

	if !gjson.Get(res.Body, "has_more").Bool() {
		return
	}
	cursor := gjson.Get(res.Body, "cursor").String()

	for {
		seelog.Debug("groups/members/list")
		type CursorParam struct {
			Cursor string `json:"cursor"`
		}
		cp := CursorParam{
			Cursor: cursor,
		}
		cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/groups/members/list/continue", cp)
		if !cont {
			return
		}

		w.processResult(res, tc)

		cursor = gjson.Get(res.Body, "cursor").String()

		if !gjson.Get(res.Body, "has_more").Bool() {
			return
		}
	}
}

func (w *WorkerTeamGroupMemberList) processResult(res *api.ApiRpcResponse, tc *ContextTeamGroupResult) {
	members := gjson.Get(res.Body, "members")
	if !members.Exists() {
		seelog.Debugf("`members` data not found")
		return
	}

	for _, member := range members.Array() {
		teamMemberId := member.Get("profile.team_member_id").String()
		key := tc.GroupId + "@" + teamMemberId

		c := ContextTeamGroupMemberList{
			GroupId:      tc.GroupId,
			GroupName:    tc.GroupName,
			TeamMemberId: teamMemberId,
			Member:       json.RawMessage(member.Raw),
		}

		w.Pipeline.Enqueue(
			workflow.MarshalTask(
				w.NextTask,
				key,
				c,
			),
		)
	}
}
