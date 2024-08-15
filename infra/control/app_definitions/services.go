package app_definitions

// App service keys
const (
	ServiceKeyDropboxIndividual = "dropbox_individual"
	ServiceKeyDropboxTeam       = "dropbox_team"
	ServiceKeyDropboxSign       = "dropbox_sign"
	ServiceKeyGithub            = "github"
	ServiceKeyAsana             = "asana"
	ServiceKeyDeepl             = "deepl"
	ServiceKeySlack             = "slack"
	ServiceKeyFigma             = "figma"
)

var (
	AllServiceKeys = []string{
		ServiceKeyDropboxIndividual,
		ServiceKeyDropboxTeam,
		ServiceKeyDropboxSign,
		ServiceKeyAsana,
		ServiceKeyDeepl,
		ServiceKeyFigma,
		ServiceKeyGithub,
		ServiceKeySlack,
	}
)
