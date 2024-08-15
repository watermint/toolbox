package api_conn

import (
	"github.com/watermint/toolbox/infra/control/app_control"
)

const (
	ServiceTagUtility         = ""
	ServiceTagDropbox         = "dropbox"
	ServiceTagDropboxBusiness = "dropbox_business"
	ServiceTagDropboxSign     = "dropbox_sign"
	ServiceTagAsana           = "asana"
	ServiceTagDeepl           = "deepl"
	ServiceTagFigma           = "figma"
	ServiceTagGithub          = "github"
	ServiceTagGoogleCalendar  = "google_calendar"
	ServiceTagGoogleMail      = "google_mail"
	ServiceTagGoogleSheets    = "google_sheets"
	ServiceTagGoogleTranslate = "google_translate"
	ServiceTagSlack           = "slack"
)

var (
	Services = []string{
		ServiceTagDropbox,
		ServiceTagDropboxBusiness,
		ServiceTagDropboxSign,
		ServiceTagAsana,
		ServiceTagDeepl,
		ServiceTagFigma,
		ServiceTagGithub,
		ServiceTagGoogleCalendar,
		ServiceTagGoogleMail,
		ServiceTagGoogleSheets,
		ServiceTagSlack,
		ServiceTagUtility,
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

// KeyConnection Key auth connection
type KeyConnection interface {
	Connection

	IsKeyAuth() bool
}

// ScopedConnection OAuth2 Scoped connection
type ScopedConnection interface {
	Connection

	// Update scopes
	SetScopes(scopes ...string)

	// Scopes
	Scopes() []string
}
