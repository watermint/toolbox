package goog_conn

import (
	"github.com/watermint/toolbox/domain/google/api/goog_client"
	"github.com/watermint/toolbox/essentials/api/api_conn"
)

type ConnGoogleApi interface {
	api_conn.ScopedConnection

	Client() goog_client.Client
}

type ConnGoogleMail interface {
	ConnGoogleApi
	IsGmail() bool
}

type ConnGoogleSheets interface {
	ConnGoogleApi
	IsSheets() bool
}

type ConnGoogleCalendar interface {
	ConnGoogleApi
	IsCalendar() bool
}

type ConnGoogleTranslate interface {
	ConnGoogleApi
	IsTranslate() bool
}
