package dbx_device

import (
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_rpc"
)

type RevokeSession struct {
	OnError func(err error) bool `json:"-"`
}

func (z *RevokeSession) WebSession(c *dbx_api.Context, teamMemberId, sessionId string) bool {
	p := struct {
		Tag          string `json:".tag"`
		SessionId    string `json:"session_id"`
		TeamMemberId string `json:"team_member_id"`
	}{
		Tag:          "web_session",
		SessionId:    sessionId,
		TeamMemberId: teamMemberId,
	}
	req := dbx_rpc.RpcRequest{
		Endpoint: "team/devices/revoke_device_session",
		Param:    p,
	}
	_, err := req.Call(c)
	if err != nil {
		return z.OnError(err)
	}
	return true
}

func (z *RevokeSession) DesktopClient(c *dbx_api.Context, teamMemberId, sessionId string, deleteOnUnlink bool) bool {
	p := struct {
		Tag            string `json:".tag"`
		SessionId      string `json:"session_id"`
		TeamMemberId   string `json:"team_member_id"`
		DeleteOnUnlink bool   `json:"delete_on_unlink,omitempty"`
	}{
		Tag:            "desktop_client",
		SessionId:      sessionId,
		TeamMemberId:   teamMemberId,
		DeleteOnUnlink: deleteOnUnlink,
	}
	req := dbx_rpc.RpcRequest{
		Endpoint: "team/devices/revoke_device_session",
		Param:    p,
	}
	_, err := req.Call(c)
	if err != nil {
		return z.OnError(err)
	}
	return true
}

func (z *RevokeSession) MobileClient(c *dbx_api.Context, teamMemberId, sessionId string) bool {
	p := struct {
		Tag          string `json:".tag"`
		SessionId    string `json:"session_id"`
		TeamMemberId string `json:"team_member_id"`
	}{
		Tag:          "mobile_client",
		SessionId:    sessionId,
		TeamMemberId: teamMemberId,
	}
	req := dbx_rpc.RpcRequest{
		Endpoint: "team/devices/revoke_device_session",
		Param:    p,
	}
	_, err := req.Call(c)
	if err != nil {
		return z.OnError(err)
	}
	return true
}
