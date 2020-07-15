package goog_auth

import (
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

const (
	// Create, read, update, and delete labels only.
	ScopeGmailLabels = "https://www.googleapis.com/auth/gmail.labels"

	// Send messages only. No read or modify privileges on mailbox.
	ScopeGmailSend = "https://www.googleapis.com/auth/gmail.send"

	// Read all resources and their metadata—no write operations.
	ScopeGmailReadonly = "https://www.googleapis.com/auth/gmail.readonly"

	// Create, read, update, and delete drafts. Send messages and drafts.
	ScopeGmailCompose = "https://www.googleapis.com/auth/gmail.compose"

	// Insert and import messages only.
	ScopeGmailInsert = "https://www.googleapis.com/auth/gmail.insert"

	// 	All read/write operations except immediate, permanent deletion of threads and messages, bypassing Trash.
	ScopeGmailModify = "https://www.googleapis.com/auth/gmail.modify"

	// Read resources metadata including labels, history records, and email message headers, but not the message body or attachments.
	ScopeGmailMetadata = "https://www.googleapis.com/auth/gmail.metadata"

	// Manage basic mail settings.
	ScopeGmailSettingsBasic = "https://www.googleapis.com/auth/gmail.settings.basic"

	// Manage sensitive mail settings, including forwarding rules and aliases.
	ScopeGmailSettingsSharing = "https://www.googleapis.com/auth/gmail.settings.sharing"

	// Full access to the account, including permanent deletion of threads and messages.
	ScopeGmailFull = "https://mail.google.com/"
)

func NewApp(ctl app_control.Control) api_auth.App {
	return &App{
		ctl: ctl,
		res: api_appkey.New(ctl),
	}
}

type App struct {
	ctl app_control.Control
	res api_appkey.Resource
}

func (z *App) UsePKCE() bool {
	return false
}

func (z *App) Config(scopes []string) *oauth2.Config {
	key, secret := z.res.Key(api_auth.GoogleMail)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Endpoint:     google.Endpoint,
		Scopes:       scopes,
	}
}
