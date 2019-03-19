package dbx_team

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type LinkedApp struct {
	TeamMemberId   string          `json:"team_member_id"`
	LinkedApiAppId string          `json:"linked_api_app_id"`
	LinkedApiApp   json.RawMessage `json:"linked_api_app"`
}

type LinkedAppList struct {
	OnError func(err error) bool
	OnEntry func(profile *LinkedApp) bool
}

func (z *LinkedAppList) List(c *dbx_api.DbxContext) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/linked_apps/list_members_linked_apps",
		EndpointListContinue: "team/linked_apps/list_members_linked_apps",
		UseHasMore:           true,
		ResultTag:            "apps",
		OnError:              z.OnError,
		OnEntry: func(result gjson.Result) bool {
			teamMemberId := result.Get("team_member_id").String()
			apps := result.Get("linked_api_apps")
			if !apps.Exists() || !apps.IsArray() {
				c.Log().Debug(
					"app not found in the result",
					zap.String("body", result.Str),
				)
				return false
			}
			for _, app := range apps.Array() {
				appId := app.Get("app_id").String()
				e := &LinkedApp{
					TeamMemberId:   teamMemberId,
					LinkedApiAppId: appId,
					LinkedApiApp:   json.RawMessage(app.Raw),
				}
				if z.OnEntry != nil {
					if !z.OnEntry(e) {
						return false
					}
				}
			}
			return true
		},
	}

	return list.List(c, nil)
}
