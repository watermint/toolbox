package auth

import "github.com/watermint/toolbox/infra/api/api_auth"

var (
	Scopes = []string{
		api_auth.DropboxTokenFull,
		api_auth.DropboxTokenBusinessInfo,
		api_auth.DropboxTokenBusinessFile,
		api_auth.DropboxTokenBusinessManagement,
		api_auth.DropboxTokenBusinessAudit,
	}
)
