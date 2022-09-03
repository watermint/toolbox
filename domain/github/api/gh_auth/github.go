package gh_auth

import (
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
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
	Github = api_auth2.OAuthAppData{
		AppKeyName:       app.ServiceGithub,
		EndpointAuthUrl:  "https://github.com/login/oauth/authorize",
		EndpointTokenUrl: "https://github.com/login/oauth/access_token",
		EndpointStyle:    api_auth2.AuthStyleAutoDetect,
		UsePKCE:          false,
		RedirectUrl:      "",
	}
)
