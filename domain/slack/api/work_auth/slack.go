package work_auth

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
)

const (
	// https://api.slack.com/scopes/channels:read
	ScopeChannelsRead = "channels:read"

	// https://api.slack.com/scopes/channels:history
	ScopeChannelsHistory = "channels:history"

	// https://api.slack.com/scopes/users:read
	ScopeUsersRead = "users:read"
)

var (
	Slack = api_auth.OAuthAppData{
		AppKeyName:       app.ServiceSlack,
		EndpointAuthUrl:  "https://slack.com/oauth/v2/authorize",
		EndpointTokenUrl: "https://slack.com/api/oauth.v2.access",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          false,
		RedirectUrl:      "",
	}
)
