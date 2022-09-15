package api_auth_repo

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"reflect"
	"testing"
	"time"
)

func TestNewBasic(t *testing.T) {
	baseRepo, err := NewInMemory()
	if err != nil {
		t.Error(err)
		return
	}

	repo := NewBasic(baseRepo)
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
	entity2 := api_auth.BasicEntity{
		KeyName:  "watermint",
		PeerName: "green",
		Credential: api_auth.BasicCredential{
			Username: "green",
			Password: "*SECRET*GREEN*",
		},
		Description: "green user",
		Timestamp:   ts,
	}
	entity3 := api_auth.BasicEntity{
		KeyName:  "watermint",
		PeerName: "blue",
		Credential: api_auth.BasicCredential{
			Username: "blue",
			Password: "*SECRET*BLUE*",
		},
		Description: "blue user",
		Timestamp:   ts,
	}
	entities := []api_auth.BasicEntity{entity1, entity2, entity3}

	for _, entity := range entities {
		repo.Put(entity)
		e, found := repo.Get(entity.KeyName, entity.PeerName)
		if !found {
			t.Error("not found", entity)
			return
		}
		if !reflect.DeepEqual(entity, e) {
			t.Error("mismatch", e, entity)
			return
		}
	}

	listResult := repo.List("watermint")
	if len(listResult) != 3 {
		t.Error(listResult)
	}
	var match1, match2, match3 bool
	for _, e := range listResult {
		if reflect.DeepEqual(entity1, e) {
			match1 = true
		}
		if reflect.DeepEqual(entity2, e) {
			match2 = true
		}
		if reflect.DeepEqual(entity3, e) {
			match3 = true
		}
	}
	if !match1 || !match2 || !match3 {
		t.Error(match1, match2, match3)
	}

	repo.Delete(entity1.KeyName, entity1.PeerName)
	_, found := repo.Get(entity1.KeyName, entity1.PeerName)
	if found {
		t.Error("deletion failure")
		return
	}

	repo.Close()
}
