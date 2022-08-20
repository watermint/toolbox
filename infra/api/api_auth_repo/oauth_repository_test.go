package api_auth_repo

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"reflect"
	"testing"
	"time"
)

func TestNewOAuth(t *testing.T) {
	baseRepo, err := NewInMemory()
	if err != nil {
		t.Error(err)
		return
	}

	expiry := time.Now().Truncate(time.Second).UTC()

	repo := NewOAuth(baseRepo)

	entity1 := api_auth.OAuthEntity{
		KeyName:  "watermint",
		Scopes:   []string{"toolbox:read", "toolbox:write"},
		PeerName: "default",
		Token: api_auth.OAuthTokenData{
			AccessToken:  "*SECRET*ACCESS*",
			RefreshToken: "*SECRET*REFRESH*",
			Expiry:       expiry.Add(1 * time.Hour),
		},
		Description: "default connection",
	}
	entity1b := api_auth.OAuthEntity{
		KeyName:  "watermint",
		Scopes:   []string{"toolbox:write", "toolbox:read"},
		PeerName: "default",
		Token: api_auth.OAuthTokenData{
			AccessToken:  "*SECRET*ACCESS*",
			RefreshToken: "*SECRET*REFRESH*",
			Expiry:       expiry.Add(1 * time.Hour),
		},
		Description: "default connection",
	}
	entity2 := api_auth.OAuthEntity{
		KeyName:  "watermint",
		Scopes:   []string{"toolbox:read", "toolbox:write"},
		PeerName: "green",
		Token: api_auth.OAuthTokenData{
			AccessToken:  "*SECRET*ACCESS*GREEN*",
			RefreshToken: "*SECRET*REFRESH*GREEN*",
			Expiry:       expiry.Add(1 * time.Hour),
		},
		Description: "green connection",
	}
	entity3 := api_auth.OAuthEntity{
		KeyName:  "watermint",
		Scopes:   []string{"toolbox:read"},
		PeerName: "default",
		Token: api_auth.OAuthTokenData{
			AccessToken:  "*SECRET*ACCESS*READONLY*",
			RefreshToken: "*SECRET*REFRESH*READONLY*",
			Expiry:       expiry.Add(1 * time.Hour),
		},
		Description: "green connection",
	}
	entities := []api_auth.OAuthEntity{entity1, entity2, entity3}

	for _, entity := range entities {
		repo.Put(entity)
		e, found := repo.Get(entity.KeyName, entity.Scopes, entity.PeerName)
		if !found {
			t.Error("not found", entity)
			return
		}
		if !reflect.DeepEqual(entity, e) {
			t.Error("mismatch", e, entity)
			return
		}
	}

	// overwrite with entity1b (different scope order)
	{
		repo.Put(entity1b)
		e, found := repo.Get(entity1.KeyName, entity1.Scopes, entity1.PeerName)
		if !found {
			t.Error("not found", entity1)
			return
		}
		if !reflect.DeepEqual(entity1, e) {
			t.Error("mismatch")
			return
		}
	}

	elw := repo.List("watermint", []string{"toolbox:read", "toolbox:write"})
	if len(elw) != 2 {
		t.Error(elw)
	}
	match1 := false
	match2 := false
	for _, e := range elw {
		if reflect.DeepEqual(entity1, e) {
			match1 = true
		}
		if reflect.DeepEqual(entity2, e) {
			match2 = true
		}
	}
	if !match1 {
		t.Error(match1)
	}
	if !match2 {
		t.Error(match2)
	}

	repo.Delete(entity1.KeyName, entity1.Scopes, entity1.PeerName)
	_, found := repo.Get(entity1.KeyName, entity1.Scopes, entity1.PeerName)
	if found {
		t.Error("deletion failure")
		return
	}

	repo.Close()
}
