package app_definitions

// App service keys
// These keys are used for identifying services in the app key repository.
const (
	AppKeyAsana             = "asana"
	AppKeyDeepl             = "deepl"
	AppKeyDropboxIndividual = "dropbox_individual"
	AppKeyDropboxSign       = "dropbox_sign"
	AppKeyDropboxTeam       = "dropbox_team"
	AppKeyFigma             = "figma"
	AppKeyGithubPublic      = "github_public"
	AppKeyGithubRepo        = "github_repo"
	AppKeySlack             = "slack"
)

var (
	AllAppKeys = []string{
		AppKeyAsana,
		AppKeyDeepl,
		AppKeyDropboxIndividual,
		AppKeyDropboxSign,
		AppKeyDropboxTeam,
		AppKeyFigma,
		AppKeyGithubPublic,
		AppKeyGithubRepo,
		AppKeySlack,
	}
)

// Scope labels for services
// These labels are used for categorizing services in the UI, documentation, etc.
const (
	ScopeLabelAsana             = "asana"
	ScopeLabelDeepl             = "deepl"
	ScopeLabelDropboxIndividual = "dropbox_individual"
	ScopeLabelDropboxSign       = "dropbox_sign"
	ScopeLabelDropboxTeam       = "dropbox_team"
	ScopeLabelFigma             = "figma"
	ScopeLabelGithub            = "github"
	ScopeLabelSlack             = "slack"
	ScopeLabelUtility           = ""
)

var (
	Services = []string{
		ScopeLabelAsana,
		ScopeLabelDeepl,
		ScopeLabelDropboxIndividual,
		ScopeLabelDropboxSign,
		ScopeLabelDropboxTeam,
		ScopeLabelFigma,
		ScopeLabelGithub,
		ScopeLabelSlack,
		ScopeLabelUtility,
	}
)
