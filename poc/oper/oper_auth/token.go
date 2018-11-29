package oper_auth

import (
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/poc/oper/oper_api"
)

type DropboxBusinessManagement struct {
	Context *dbx_api.Context
}

func (z *DropboxBusinessManagement) AppTypeLabel() string {
	return "Dropbox Business API"
}

func (z *DropboxBusinessManagement) AppAccessLabel() string {
	return "Team member management"
}

func (z *DropboxBusinessManagement) ApiKey() string {
	return app.BusinessManagementAppKey
}

func (z *DropboxBusinessManagement) ApiSecret() string {
	return app.BusinessManagementAppSecret
}

func (z *DropboxBusinessManagement) WithApi(api *dbx_api.Context) oper_api.DropboxApiToken {
	return &DropboxBusinessManagement{
		Context: api,
	}
}

func (z *DropboxBusinessManagement) Tag() string {
	return "business_management"
}

func (z *DropboxBusinessManagement) Api() *dbx_api.Context {
	return z.Context
}
