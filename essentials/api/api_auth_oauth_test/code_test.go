package api_auth_oauth_test

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestCode(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		session := api_auth_oauth.NewSessionCodeAuth(ctl)
		entity, err := session.Start(api_auth.OAuthSessionData{
			AppData: api_auth.OAuthAppData{
				AppKeyName:       "test",
				EndpointAuthUrl:  "https://example.com/auth",
				EndpointTokenUrl: "https://example.com/token",
				EndpointStyle:    api_auth.AuthStyleAutoDetect,
				UsePKCE:          false,
				RedirectUrl:      "",
			},
			PeerName: "default",
			Scopes: []string{
				"test:write", "test:read",
			},
		})
		// should fail with user cancellation
		if err != app_definitions.ErrorUserCancelled {
			t.Error(entity)
		}
	})
}
