package app_definitions

const (
	ServiceDropboxIndividual = "dropbox_individual"
	ServiceDropboxTeam       = "dropbox_team"
	ServiceDropboxSign       = "dropbox_sign"
	ServiceGithub            = "github"
	ServiceGoogleMail        = "google_mail2024"     // Adding 2024 to avoid conflict with existing auth data #777
	ServiceGoogleSheets      = "google_sheets2024"   // Adding 2024 to avoid conflict with existing auth data #777
	ServiceGoogleCalendar    = "google_calendar2024" // Adding 2024 to avoid conflict with existing auth data #777
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
