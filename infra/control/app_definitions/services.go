package app_definitions

const (
	ServiceDropboxIndividual = "dropbox_individual"
	ServiceDropboxTeam       = "dropbox_team"
	ServiceDropboxSign       = "dropbox_sign"
	ServiceGithub            = "github"
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
		ServiceSlack,
	}
)
