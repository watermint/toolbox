package api_auth

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"golang.org/x/oauth2"
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

type TokenContainer struct {
	Token       string
	TokenType   string
	PeerName    string
	Description string
}

type TokenContext interface {
	Token() TokenContainer
}

// Application key/secret manager
type App interface {
	// Key/secret for token type
	AppKey(tokenType string) (key, secret string)

	// OAuth2 config
	Config(tokenType string) *oauth2.Config
}

// Auth interface for console UI
type Console interface {
	Auth(tokenType string) (ctx api_context.Context, err error)
}

// Auth interface for web UI
type Web interface {
	// Create new state and url.
	New(tokenType, redirectUrl string) (state, url string, err error)

	// Proceed authorisation process.
	Auth(state, code string) (peerName string, ctx api_context.Context, err error)

	// Retrieve existing connection.
	Get(state string) (peerName string, ctx api_context.Context, err error)

	// List existing connections
	List(tokenType string) (token []TokenContainer, err error)
}
