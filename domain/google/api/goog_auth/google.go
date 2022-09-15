package goog_auth

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
)

// Gmail scopes
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

// Google Sheets API scopes
// https://developers.google.com/sheets/api/guides/authorizing
const (
	// Allows read-only access to the user's sheets and their properties.
	ScopeSheetsReadOnly = "https://www.googleapis.com/auth/spreadsheets.readonly"

	// Allows read/write access to the user's sheets and their properties.
	ScopeSheetsReadWrite = "https://www.googleapis.com/auth/spreadsheets"

	// Allows read-only access to the user's file metadata and file content.
	ScopeSheetsDriveReadOnly = "https://www.googleapis.com/auth/drive.readonly"

	// Per-file access to files created or opened by the app.
	ScopeSheetsDriveFile = "https://www.googleapis.com/auth/drive.file"

	// Full, permissive scope to access all of a user's files. Request this scope only when it is strictly necessary.
	ScopeSheetsFull = "https://www.googleapis.com/auth/drive"
)

// Google Calendar API scopes
// https://developers.google.com/calendar/api/guides/auth
const (
	// ScopeCalendarAllReadWrite read/write access to Calendars
	ScopeCalendarAllReadWrite = "https://www.googleapis.com/auth/calendar"
	// ScopeCalendarAllReadOnly read-only access to Calendars
	ScopeCalendarAllReadOnly = "https://www.googleapis.com/auth/calendar.readonly"
	// ScopeCalendarEventsReadWrite read/write access to Events
	ScopeCalendarEventsReadWrite = "https://www.googleapis.com/auth/calendar.events"
	// ScopeCalendarEventsReadOnly read-only access to Events
	ScopeCalendarEventsReadOnly = "https://www.googleapis.com/auth/calendar.events.readonly"
	// ScopeCalendarSettingsReadOnly read-only access to Settings
	ScopeCalendarSettingsReadOnly = "https://www.googleapis.com/auth/calendar.settings.readonly"
	// ScopeCalendarAddonsExecute run as a Calendar add-on
	ScopeCalendarAddonsExecute = "https://www.googleapis.com/auth/calendar.addons.execute"
)

var (
	Mail = api_auth.OAuthAppData{
		AppKeyName:       app.ServiceGoogleMail,
		EndpointAuthUrl:  "https://accounts.google.com/o/oauth2/auth",
		EndpointTokenUrl: "https://oauth2.googleapis.com/token",
		EndpointStyle:    api_auth.AuthStyleInParams,
		UsePKCE:          false,
		RedirectUrl:      "",
	}

	Calendar = api_auth.OAuthAppData{
		AppKeyName:       app.ServiceGoogleCalendar,
		EndpointAuthUrl:  "https://accounts.google.com/o/oauth2/auth",
		EndpointTokenUrl: "https://oauth2.googleapis.com/token",
		EndpointStyle:    api_auth.AuthStyleInParams,
		UsePKCE:          false,
		RedirectUrl:      "",
	}

	Sheets = api_auth.OAuthAppData{
		AppKeyName:       app.ServiceGoogleSheets,
		EndpointAuthUrl:  "https://accounts.google.com/o/oauth2/auth",
		EndpointTokenUrl: "https://oauth2.googleapis.com/token",
		EndpointStyle:    api_auth.AuthStyleInParams,
		UsePKCE:          false,
		RedirectUrl:      "",
	}
)
