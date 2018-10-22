package dbx_team

import (
	"encoding/json"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_rpc"
)

type LinkedApp struct {
	TeamMemberId   string          `json:"team_member_id"`
	LinkedApiAppId string          `json:"linked_api_app_id"`
	LinkedApiApp   json.RawMessage `json:"linked_api_app"`
}

type LinkedAppList struct {
	OnError func(annotation dbx_api.ErrorAnnotation) bool
	OnEntry func(profile *LinkedApp) bool
}

func (a *LinkedAppList) List(c *dbx_api.Context) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/linked_apps/list_members_linked_apps",
		EndpointListContinue: "team/linked_apps/list_members_linked_apps",
		UseHasMore:           true,
		ResultTag:            "apps",
		OnError:              a.OnError,
		OnEntry: func(result gjson.Result) bool {
			teamMemberId := result.Get("team_member_id").String()
			apps := result.Get("linked_api_apps")
			if !apps.Exists() || !apps.IsArray() {
				seelog.Debugf("Apps not found in the result [%s]", result)
				return false
			}
			for _, app := range apps.Array() {
				appId := app.Get("app_id").String()
				e := &LinkedApp{
					TeamMemberId:   teamMemberId,
					LinkedApiAppId: appId,
					LinkedApiApp:   json.RawMessage(app.Raw),
				}
				if a.OnEntry != nil {
					if !a.OnEntry(e) {
						return false
					}
				}
			}
			return true
		},
	}

	return list.List(c, nil)
}
