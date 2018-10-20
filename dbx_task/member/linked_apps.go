package member

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/workflow"
)

const (
	WORKER_TEAM_MEMBER_LINKEDAPPS = "team/member/linkedapps"
)

type WorkerTeamMemberLinkedApps struct {
	workflow.SimpleWorkerImpl
	ApiManagement *dbx_api.ApiContext
	NextTask      string
}

type ContextTeamMemberLinkedApps struct {
	Cursor string `json:"cursor"`
}

type ContextTeamMemberLinkedAppsResult struct {
	TeamMemberId   string          `json:"team_member_id"`
	LinkedApiAppId string          `json:"linked_api_app_id"`
	LinkedApiApp   json.RawMessage `json:"linked_api_app"`
}

func (w *WorkerTeamMemberLinkedApps) Prefix() string {
	return WORKER_TEAM_MEMBER_LINKEDAPPS
}

func (w *WorkerTeamMemberLinkedApps) Exec(task *workflow.Task) {
	tc := &ContextTeamMemberLinkedApps{}
	workflow.UnmarshalContext(task, tc)

	if tc.Cursor == "" {
		w.list(task)
	} else {
		w.listContinue(tc.Cursor, task)
	}

}

func (w *WorkerTeamMemberLinkedApps) list(task *workflow.Task) {
	type ListContinueParam struct {
		Cursor string `json:"cursor"`
	}
	lp := ListContinueParam{
		Cursor: "",
	}
	cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/linked_apps/list_members_linked_apps", lp)
	if !cont {
		return
	}
	w.processResult(res, task)
}

func (w *WorkerTeamMemberLinkedApps) listContinue(cursor string, task *workflow.Task) {
	type ListContinueParam struct {
		Cursor string `json:"cursor"`
	}
	lp := ListContinueParam{
		Cursor: cursor,
	}
	cont, res, _ := w.Pipeline.TaskRpc(task, w.ApiManagement, "team/linked_apps/list_members_linked_apps", lp)
	if !cont {
		return
	}

	w.processResult(res, task)
}

func (w *WorkerTeamMemberLinkedApps) processResult(res *dbx_api.ApiRpcResponse, task *workflow.Task) {
	membersApps := gjson.Get(res.Body, "apps")

	for _, member := range membersApps.Array() {
		teamMemberId := member.Get("team_member_id").String()

		for _, app := range member.Get("linked_api_apps").Array() {
			appId := app.Get("app_id").String()
			c := ContextTeamMemberLinkedAppsResult{
				TeamMemberId:   teamMemberId,
				LinkedApiAppId: appId,
				LinkedApiApp:   json.RawMessage(app.Raw),
			}
			w.Pipeline.Enqueue(
				workflow.MarshalTask(
					w.NextTask,
					teamMemberId+"@"+appId,
					c,
				),
			)
		}
	}

	hasMoreJson := gjson.Get(res.Body, "has_more")
	if hasMoreJson.Exists() && hasMoreJson.Bool() {
		cursorJson := gjson.Get(res.Body, "cursor")
		if !cursorJson.Exists() {
			seelog.Debug("Cursor not found in the response (has_more appear and true)")
			return
		}
		c := ContextTeamMemberLinkedApps{
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
