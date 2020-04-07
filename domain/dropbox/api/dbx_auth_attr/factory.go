package dbx_auth_attr

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
)

func NewConsole(c app_control.Control, peerName string) api_auth.Console {
	//oa := dbx_auth.NewConsoleOAuth(c, peerName)
	oa := dbx_auth.NewConsoleRedirect(c, peerName)
	aa := NewConsoleAttr(c, oa)
	ca := dbx_auth.NewConsoleCache(c, aa)
	return ca
}

func NewConsoleAttr(c app_control.Control, auth api_auth.Console) api_auth.Console {
	return &Attr{
		ctl:  c,
		auth: auth,
	}
}
