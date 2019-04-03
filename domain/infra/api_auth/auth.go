package api_auth

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
)

const (
	DropboxTokenNoAuth             = ""
	DropboxTokenFull               = "user_full"
	DropboxTokenApp                = "user_app"
	DropboxTokenBusinessInfo       = "business_info"
	DropboxTokenBusinessAudit      = "business_audit"
	DropboxTokenBusinessFile       = "business_file"
	DropboxTokenBusinessManagement = "business_management"
)

type Auth interface {
	Auth(tokenType string) (ctx api_context.Context, err error)
}

type TokenContainer struct {
	Token     string
	TokenType string
	PeerName  string
}

type TokenContext interface {
	Token() TokenContainer
}
