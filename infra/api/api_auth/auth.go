package api_auth

import (
	"golang.org/x/oauth2"
)

const (
	DropboxTokenFull               = "user_full"
	DropboxTokenApp                = "user_app"
	DropboxTokenBusinessInfo       = "business_info"
	DropboxTokenBusinessAudit      = "business_audit"
	DropboxTokenBusinessFile       = "business_file"
	DropboxTokenBusinessManagement = "business_management"
	Github                         = "github"
)

// Application key/secret manager
type App interface {
	// Key/secret for token type
	AppKey(scope string) (key, secret string)

	// OAuth2 config
	Config(scope string) *oauth2.Config
}

// Auth interface for console UI
type Console interface {
	PeerName() string

	Auth(scope string) (token Context, err error)
}
