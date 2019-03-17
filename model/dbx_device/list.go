package dbx_device

import (
	"encoding/json"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_device/flat_device"
	"github.com/watermint/toolbox/model/dbx_rpc"
	"go.uber.org/zap"
)

type DeviceList struct {
	IncludeWebSessions    bool `json:"include_web_sessions"`
	IncludeDesktopClients bool `json:"include_desktop_clients"`
	IncludeMobileClients  bool `json:"include_mobile_clients"`

	OnError   func(err error) bool                          `json:"-"`
	OnWeb     func(web *flat_device.WebSession) bool        `json:"-"`
	OnDesktop func(desktop *flat_device.DesktopClient) bool `json:"-"`
	OnMobile  func(mobile *flat_device.MobileClient) bool   `json:"-"`
}

func (z *DeviceList) List(c *dbx_api.DbxContext) bool {
	list := dbx_rpc.RpcList{
		EndpointList:         "team/devices/list_members_devices",
		EndpointListContinue: "team/devices/list_members_devices",
		UseHasMore:           true,
		ResultTag:            "devices",
		OnError:              z.OnError,
		OnEntry: func(result gjson.Result) bool {
			d := Device{}
			err := json.Unmarshal([]byte(result.Raw), &d)
			if err != nil {
				c.Log().Debug("unable to unmarshal response", zap.Error(err))
			} else {
				for _, x := range d.Desktop() {
					z.OnDesktop(x)
				}
				for _, x := range d.Mobile() {
					z.OnMobile(x)
				}
				for _, x := range d.Web() {
					z.OnWeb(x)
				}
			}
			return true
		},
	}

	return list.List(c, z)
}
