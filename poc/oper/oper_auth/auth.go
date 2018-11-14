package oper_auth

import "github.com/watermint/toolbox/model/dbx_api"

type DropboxApiToken interface {
	Api() *dbx_api.Context
}

type DropboxBusinessManagement struct {
	Context *dbx_api.Context
}

func (z *DropboxBusinessManagement) Api() *dbx_api.Context {
	return z.Context
}
