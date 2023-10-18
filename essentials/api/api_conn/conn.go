package api_conn

import (
	"github.com/watermint/toolbox/infra/control/app_control"
)

const (
	ServiceUtility         = ""
	ServiceDropbox         = "dropbox"
	ServiceDropboxBusiness = "dropbox_business"
	ServiceAsana           = "asana"
	ServiceGithub          = "github"
	ServiceGoogleCalendar  = "google_calendar"
	ServiceGoogleMail      = "google_mail"
	ServiceGoogleSheets    = "google_sheets"
	ServiceGoogleTranslate = "google_translate"
	ServiceDropboxSign     = "dropbox_sign"
	ServiceSlack           = "slack"
	ServiceFigma           = "figma"
)

var (
	Services = []string{
		ServiceAsana,
		ServiceDropbox,
		ServiceDropboxBusiness,
		ServiceDropboxSign,
		ServiceFigma,
		ServiceGithub,
		ServiceGoogleCalendar,
		ServiceGoogleMail,
		ServiceGoogleSheets,
		ServiceSlack,
		ServiceUtility,
	}
)

const (
	DefaultPeerName = "default"
)

type Connection interface {
	// Connect to api
	Connect(ctl app_control.Control) (err error)

	// Peer name
	PeerName() string

	// Update peer name
	SetPeerName(name string)

	// Scope label
	ScopeLabel() string

	// Name tag of the service
	ServiceName() string
}

// BasicConnection Basic auth connection
type BasicConnection interface {
	Connection

	IsBasic() bool
}

// ScopedConnection OAuth2 Scoped connection
type ScopedConnection interface {
	Connection

	// Update scopes
	SetScopes(scopes ...string)

	// Scopes
	Scopes() []string
}
