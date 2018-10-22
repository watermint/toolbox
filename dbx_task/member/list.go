package member

import (
	"encoding/json"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
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
	ml := dbx_team.MembersList{
		OnError: w.Pipeline.HandleGeneralFailure,
		OnEntry: func(p *dbx_task.Profile) bool {
			w.Pipeline.Enqueue(
				workflow.MarshalTask(
					w.NextTask,
					p.Email,
					p,
				),
			)
			return true
		},
	}
	ml.List(w.Api, w.IncludeRemoved)
}
