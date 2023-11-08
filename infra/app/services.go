package app

const (
	ServiceDropboxIndividual = "dropbox_individual"
	ServiceDropboxTeam       = "dropbox_team"
	ServiceDropboxSign       = "dropbox_sign"
	ServiceGithub            = "github"
	ServiceGoogleMail        = "google_mail"
	ServiceGoogleSheets      = "google_sheets"
	ServiceGoogleCalendar    = "google_calendar"
	ServiceGoogleTranslate   = "google_translate"
	ServiceAsana             = "asana"
	ServiceDeepl             = "deepl"
	ServiceSlack             = "slack"
	ServiceFigma             = "figma"
)

var (
	AllServices = []string{
		ServiceDropboxIndividual,
		ServiceDropboxTeam,
		ServiceDropboxSign,
		ServiceAsana,
		ServiceDeepl,
		ServiceFigma,
		ServiceGithub,
		ServiceGoogleCalendar,
		ServiceGoogleMail,
		ServiceGoogleSheets,
		ServiceSlack,
	}
)
