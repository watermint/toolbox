package group

import (
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
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
	gl := dbx_team.GroupList{
		OnError: w.Pipeline.HandleGeneralFailure,
		OnEntry: func(group *dbx_team.Group) bool {
			w.Pipeline.Enqueue(
				workflow.MarshalTask(
					w.NextTask,
					group.GroupId,
					group,
				),
			)
			return true
		},
	}
	gl.List(w.Api)
}

type WorkerTeamGroupMemberList struct {
	workflow.SimpleWorkerImpl
	Api      *dbx_api.Context
	NextTask string
}

func (w *WorkerTeamGroupMemberList) Prefix() string {
	return "team/group/member/list"
}

func (w *WorkerTeamGroupMemberList) Exec(task *workflow.Task) {
	tc := &dbx_team.Group{}
	workflow.UnmarshalContext(task, tc)

	gml := dbx_team.GroupMemberList{
		OnError: w.Pipeline.HandleGeneralFailure,
		OnEntry: func(r *dbx_team.GroupMember) bool {
			key := r.GroupId + "@" + r.TeamMemberId
			w.Pipeline.Enqueue(
				workflow.MarshalTask(
					w.NextTask,
					key,
					r,
				),
			)
			return true
		},
	}
	gml.List(w.Api, tc)
}
