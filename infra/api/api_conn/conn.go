package api_conn

import (
	"github.com/watermint/toolbox/infra/control/app_control"
)

const (
	ServiceUtility         = ""
	ServiceDropbox         = "dropbox"
	ServiceDropboxBusiness = "dropbox_business"
	ServiceGithub          = "github"
	ServiceGoogleMail      = "google_mail"
	ServiceGoogleSheets    = "google_sheets"
	ServiceAsana           = "asana"
	ServiceSlack           = "slack"
)

var (
	Services = []string{
		ServiceDropbox,
		ServiceDropboxBusiness,
		ServiceGithub,
		ServiceGoogleMail,
		ServiceAsana,
		ServiceSlack,
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
