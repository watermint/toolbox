package gh_auth

import (
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
	"strings"
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

type App struct {
	ctl app_control.Control
	res api_appkey.Resource
}

func (z *App) AppKey(scope string) (key, secret string) {
	return z.res.Key(scope)
}

func (z *App) Config(scope string) *oauth2.Config {
	key, secret := z.AppKey(scope)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Endpoint:     github.Endpoint,
		Scopes:       strings.Split(scope, ","),
	}
}
