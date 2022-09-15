package api_auth_oauth

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"os"
	"reflect"
	"testing"
	"time"
)

func TestDeployEnv(t *testing.T) {
	hash := sc_random.MustGetSecureRandomString(4)
	envName := "API_AUTH_OAUTH_DEPLOY_KEY_" + hash
	expiry := time.Now().Truncate(time.Second).UTC().Add(1 * time.Hour)
	entity1 := api_auth.OAuthEntity{
		KeyName:  "watermint",
		Scopes:   []string{"toolbox:read", "toolbox:write"},
		PeerName: "default",
		Token: api_auth.OAuthTokenData{
			AccessToken:  "*SECRET*ACCESS*",
			RefreshToken: "*SECRET*REFRESH*",
			Expiry:       expiry,
		},
		Description: "default connection",
	}
	entitySerialize, err := json.Marshal(entity1)
	if err != nil {
		t.Error(err)
		return
	}
	if err := os.Setenv(envName, string(entitySerialize)); err != nil {
		t.Error(err)
	}
	defer func() {
		_ = os.Unsetenv(envName)
	}()

	session := NewSessionDeployEnv(envName)
	entity2, err := session.Start(api_auth.OAuthSessionData{
		AppData: api_auth.OAuthAppData{
			AppKeyName:       "watermint",
			EndpointAuthUrl:  "https://example.com/auth",
			EndpointTokenUrl: "https://example.com/token",
			EndpointStyle:    api_auth.AuthStyleAutoDetect,
			UsePKCE:          false,
			RedirectUrl:      "",
		},
		PeerName: "default",
		Scopes: []string{
			"toolbox:write", "toolbox:read",
		},
	})

	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(entity1, entity2) {
		t.Error(entity1, entity2)
	}
}
