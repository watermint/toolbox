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
	ServiceHelloSign       = "hellosign"
	ServiceSlack           = "slack"
)

var (
	Services = []string{
		ServiceDropbox,
		ServiceDropboxBusiness,
		ServiceAsana,
		ServiceGithub,
		ServiceGoogleCalendar,
		ServiceGoogleMail,
		ServiceGoogleSheets,
		ServiceSlack,
		ServiceHelloSign,
		ServiceUtility,
	}
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

type ScopedConnection interface {
	Connection

	// Update scopes
	SetScopes(scopes ...string)

	// Scopes
	Scopes() []string
}
