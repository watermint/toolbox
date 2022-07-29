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
	DropboxScopedIndividual        = "dropbox_scoped_individual"
	DropboxScopedTeam              = "dropbox_scoped_team"
	Github                         = "github"
	GoogleMail                     = "google_mail"
	GoogleSheets                   = "google_sheets"
	GoogleCalendar                 = "google_calendar"
	Asana                          = "asana"
	Slack                          = "slack"
)

// App OAuth Application key/secret manager
type App interface {
	// Config OAuth2 config
	Config(scope []string) *oauth2.Config

	// UsePKCE Use PKCE on authentication
	UsePKCE() bool
}

// Console OAuth interface for console UI
type Console interface {
	PeerName() string

	Auth(scope []string) (token Context, err error)
}

// Basic is for Basic Authentication interface
type Basic interface {
	PeerName() string
}
