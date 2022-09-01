package api_auth

const (
	// DropboxIndividual Dropbox for individual user (personal or end user of business team)
	DropboxIndividual = "dropbox_individual"

	// DropboxTeam Admin scope
	DropboxTeam    = "dropbox_team"
	Github         = "github"
	GoogleMail     = "google_mail"
	GoogleSheets   = "google_sheets"
	GoogleCalendar = "google_calendar"
	Asana          = "asana"
	Slack          = "slack"
)

type Auth interface {
	PeerName() string
}
