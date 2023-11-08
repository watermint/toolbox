package api_auth_basic

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_repo"
	"reflect"
	"testing"
	"time"
)

func TestRepoImpl_Start(t *testing.T) {
	ts := time.Now().Truncate(time.Second).Format(time.RFC3339)
	entity1 := api_auth.BasicEntity{
		KeyName:  "watermint",
		PeerName: "default",
		Credential: api_auth.BasicCredential{
			Username: "toolbox",
			Password: "*SECRET*",
		},
		Description: "default user",
		Timestamp:   ts,
	}

	baseRepo, err := api_auth_repo.NewInMemory()
	if err != nil {
		t.Error(err)
		return
	}

	session := NewSession(
		NewEmbedded(entity1),
		baseRepo,
	)
	entity2, err := session.Start(api_auth.BasicSessionData{
		AppData: api_auth.BasicAppData{
			AppKeyName:      "watermint",
			DontUseUsername: false,
			DontUsePassword: false,
		},
		PeerName: "default",
	})
	if err != nil {
		t.Error(err)
	}
	if !reflect.DeepEqual(entity1, entity2) {
		t.Error(entity1, entity2)
	}

	// must be registered in the repo
	repo := api_auth_repo.NewBasic(baseRepo)
	entity3, found := repo.Get(entity1.KeyName, entity1.PeerName)
	if !found {
		t.Error(found)
	}
	if !reflect.DeepEqual(entity1, entity3) {
		t.Error(entity1, entity3)
	}
}
