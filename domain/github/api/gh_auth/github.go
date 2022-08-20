package gh_auth

import (
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const (
	// Read-only access to public information.
	ScopeNoScope = ""

	// full access to private and public repositories.
	ScopeRepo = "repo"

	// Read write access to public and private repository commit statuses.
	ScopeRepoStatus = "repo:status"

	// Deployment statuses for public and private repositories.
	ScopeRepoDeployment = "repo_deployment"

	// Read/write access to code, commit statuses, repository etc for public repositories.
	ScopePublicRepo = "public_repo"
)

const (
	ScopeLabelPublic = "github_public"
	ScopeLabelRepo   = "github_repo"
)

var (
	Github = api_auth.OAuthAppData{
		AppKeyName:       api_auth.Github,
		EndpointAuthUrl:  "https://github.com/login/oauth/authorize",
		EndpointTokenUrl: "https://github.com/login/oauth/access_token",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          false,
		RedirectUrl:      "",
	}
)

func NewApp(ctl app_control.Control) api_auth.OAuthApp {
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
	key, secret := z.res.Key(api_auth.Github)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Endpoint:     github.Endpoint,
		Scopes:       scopes,
	}
}
