package member

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
	"github.com/watermint/toolbox/workflow"
)

type WorkerTeamMemberLinkedApps struct {
	workflow.SimpleWorkerImpl
	Api      *dbx_api.Context
	NextTask string
}

type ContextTeamMemberLinkedAppsResult struct {
	TeamMemberId   string          `json:"team_member_id"`
	LinkedApiAppId string          `json:"linked_api_app_id"`
	LinkedApiApp   json.RawMessage `json:"linked_api_app"`
}

func (w *WorkerTeamMemberLinkedApps) Prefix() string {
	return "team/member/linked_apps"
}

func (w *WorkerTeamMemberLinkedApps) Exec(task *workflow.Task) {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/linked_apps/list_members_linked_apps",
		EndpointListContinue: "team/linked_apps/list_members_linked_apps",
		UseHasMore:           true,
		ResultTag:            "apps",
		OnError:              w.Pipeline.HandleGeneralFailure,
		OnEntry:              w.processResult,
	}

	list.List(w.Api, nil)
}

func (w *WorkerTeamMemberLinkedApps) processResult(member gjson.Result) bool {
	//teamMemberId := member.Get("team_member_id").String()
	//apps := member.Get("linked_api_apps")
	//if !apps.Exists() || !apps.IsArray() {
	//	seelog.Debugf("Apps not found in the result [%s]", member)
	//	return false
	//}
	//for _, app := range apps.Array() {
	//	appId := app.Get("app_id").String()
	//	c := ContextTeamMemberLinkedAppsResult{
	//		TeamMemberId:   teamMemberId,
	//		LinkedApiAppId: appId,
	//		LinkedApiApp:   json.RawMessage(app.Raw),
	//	}
	//}
	return true
}
