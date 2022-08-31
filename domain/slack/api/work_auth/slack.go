package work_auth

import (
	"github.com/watermint/toolbox/infra/api/api_appkey"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"golang.org/x/oauth2"
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
		AppKeyName:       api_auth.Slack,
		EndpointAuthUrl:  "https://slack.com/oauth/v2/authorize",
		EndpointTokenUrl: "https://slack.com/api/oauth.v2.access",
		EndpointStyle:    api_auth.AuthStyleAutoDetect,
		UsePKCE:          false,
		RedirectUrl:      "",
	}
)

func New(ctl app_control.Control) api_auth.OAuthAppLegacy {
	return &App{
		ctl: ctl,
		res: api_appkey.New(ctl),
	}
}

type App struct {
	ctl app_control.Control
	res api_appkey.Resource
}

func (z App) Config(scope []string) *oauth2.Config {
	key, secret := z.res.Key(api_auth.Slack)
	return &oauth2.Config{
		ClientID:     key,
		ClientSecret: secret,
		Endpoint: oauth2.Endpoint{
			AuthURL:  "https://slack.com/oauth/v2/authorize",
			TokenURL: "https://slack.com/api/oauth.v2.access",
		},
		Scopes: scope,
	}
}

func (z App) UsePKCE() bool {
	return false
}
