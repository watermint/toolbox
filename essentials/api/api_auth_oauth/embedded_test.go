package api_auth_oauth

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"reflect"
	"testing"
	"time"
)

func TestNewSessionEmbedded(t *testing.T) {
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

	session := NewSessionEmbedded(entity1)
	entity2, err := session.Start(api_auth.OAuthSessionData{
		AppData: api_auth.OAuthAppData{
			AppKeyName:       "must_fail",
			EndpointAuthUrl:  "https://example.com/auth",
			EndpointTokenUrl: "https://example.com/token",
			EndpointStyle:    api_auth.AuthStyleAutoDetect,
			UsePKCE:          false,
			RedirectUrl:      "",
		},
		PeerName: "default",
		Scopes: []string{
			"must_fail:write", "must_fail:read",
		},
	})

	// should fail
	if err == nil {
		t.Error(entity2)
	}

	entity3, err := session.Start(api_auth.OAuthSessionData{
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
	if !reflect.DeepEqual(entity1, entity3) {
		t.Error(entity1, entity3)
	}
}
