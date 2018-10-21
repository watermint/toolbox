package group

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
	"github.com/watermint/toolbox/workflow"
)

type WorkerTeamGroupList struct {
	workflow.SimpleWorkerImpl
	Api      *dbx_api.Context
	NextTask string
}

type ContextTeamGroupResult struct {
	GroupId             string `json:"group_id,omitempty"`
	GroupName           string `json:"group_name,omitempty"`
	GroupManagementType string `json:"group_management_type,omitempty"`
	GroupExternalId     string `json:"group_external_id,omitempty"`
	MemberCount         int64  `json:"member_count,omitempty"`
}

func (w *WorkerTeamGroupList) Prefix() string {
	return "team/group/list"
}

func (w *WorkerTeamGroupList) Exec(task *workflow.Task) {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/groups/list",
		EndpointListContinue: "team/groups/list/continue",
		UseHasMore:           true,
		ResultTag:            "groups",
		HandlerError:         w.Pipeline.HandleGeneralFailure,
		HandlerEntry:         w.processResult,
	}

	list.List(w.Api, nil)
}

func (w *WorkerTeamGroupList) processResult(group gjson.Result) bool {
	groupIdJson := group.Get("group_id")

	if !groupIdJson.Exists() {
		seelog.Debugf("one of group info (group_id) not found `%s`", group.Raw)
		return false
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
	return true
}

type WorkerTeamGroupMemberList struct {
	workflow.SimpleWorkerImpl
	Api      *dbx_api.Context
	NextTask string
}

type ContextTeamGroupMemberList struct {
	GroupId      string          `json:"group_id"`
	GroupName    string          `json:"group_name"`
	TeamMemberId string          `json:"team_member_id"`
	Member       json.RawMessage `json:"member"`
}

func (w *WorkerTeamGroupMemberList) Prefix() string {
	return "team/group/member/list"
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
	}

	lp := ListParam{
		Group: GroupSelector{
			Tag:     "group_id",
			GroupId: tc.GroupId,
		},
	}

	processEntry := func(member gjson.Result) bool {
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
		return true
	}

	list := dbx_rpc.RpcList{
		EndpointList:         "team/groups/members/list",
		EndpointListContinue: "team/groups/members/list/continue",
		UseHasMore:           true,
		ResultTag:            "members",
		HandlerError:         w.Pipeline.HandleGeneralFailure,
		HandlerEntry:         processEntry,
	}

	list.List(w.Api, lp)
}
