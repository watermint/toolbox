package api_conn

import (
	"github.com/watermint/toolbox/infra/control/app_control"
)

const (
	DefaultPeerName = "default"
)

type Connection interface {
	// Connect to api
	Connect(ctl app_control.Control) (err error)

	// PeerName returns the peer name of the connection.
	// The peer name is the key to identify auth entity in the auth repository.
	// If the user want multiple connections to the same service, the peer name should be different.
	PeerName() string

	// SetPeerName set peer name
	SetPeerName(name string)

	// ScopeLabel returns the scope label of the connection.
	// The scope label is used to categorize the service in the UI, documentation, etc.
	ScopeLabel() string

	// AppKeyName returns the app key name that identifies app key in the auth repository.
	AppKeyName() string
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
