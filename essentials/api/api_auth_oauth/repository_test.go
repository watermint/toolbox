package api_auth_oauth

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"reflect"
	"testing"
	"time"
)

func TestRepository(t *testing.T) {
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

	baseRepo, err := api_auth_repo.NewInMemory()
	if err != nil {
		t.Error(err)
		return
	}

	session := NewSessionRepository(
		NewSessionEmbedded(entity1),
		baseRepo,
	)

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

	// must be registered in the repo
	repo := api_auth_repo.NewOAuth(baseRepo)
	entity3, found := repo.Get(entity1.KeyName, entity1.Scopes, entity1.PeerName)
	if !found {
		t.Error(found)
	}
	if !reflect.DeepEqual(entity1, entity3) {
		t.Error(entity1, entity3)
	}
}
