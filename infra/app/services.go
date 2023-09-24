package app

const (
	ServiceDropboxIndividual = "dropbox_individual"
	ServiceDropboxTeam       = "dropbox_team"
	ServiceGithub            = "github"
	ServiceGoogleMail        = "google_mail"
	ServiceGoogleSheets      = "google_sheets"
	ServiceGoogleCalendar    = "google_calendar"
	ServiceGoogleTranslate   = "google_translate"
	ServiceAsana             = "asana"
	ServiceDropboxSign       = "dropbox_sign"
	ServiceSlack             = "slack"
	ServiceFigma             = "figma"
)

var (
	AllServices = []string{
		ServiceDropboxIndividual,
		ServiceDropboxTeam,
		ServiceGithub,
		ServiceGoogleMail,
		ServiceGoogleSheets,
		ServiceGoogleCalendar,
		ServiceAsana,
		ServiceDropboxSign,
		ServiceSlack,
		ServiceFigma,
	}
)
