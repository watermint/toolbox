package api_auth

const (
	DropboxTokenFull  = "user_full"
	DropboxIndividual = "dropbox_individual"
	DropboxTeam       = "dropbox_team"
	Github            = "github"
	GoogleMail        = "google_mail"
	GoogleSheets      = "google_sheets"
	GoogleCalendar    = "google_calendar"
	Asana             = "asana"
	Slack             = "slack"
)

type Auth interface {
	PeerName() string
}
