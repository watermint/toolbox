package api_auth_oauth_test

import (
	api_auth2 "github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestNewSessionAlwaysFail(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		session := api_auth_oauth.NewSessionAlwaysFail(ctl)
		entity, err := session.Start(api_auth2.OAuthSessionData{
			AppData: api_auth2.OAuthAppData{
				AppKeyName:       "test",
				EndpointAuthUrl:  "https://example.com/auth",
				EndpointTokenUrl: "https://example.com/token",
				EndpointStyle:    api_auth2.AuthStyleAutoDetect,
				UsePKCE:          false,
				RedirectUrl:      "",
			},
			PeerName: "default",
			Scopes: []string{
				"test:write", "test:read",
			},
		})
		// should fail
		if err == nil {
			t.Error(entity)
		}
	})
}
