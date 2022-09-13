package api_auth_repo

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
	"reflect"
	"testing"
	"time"
)

func repoTestScenario(t *testing.T, repo api_auth.Repository) {
	ts1 := time.Now().Format(time.RFC3339)
	ts2 := time.Now().Add(-1 * time.Hour).Format(time.RFC3339)
	ts3 := time.Now().Add(1 * time.Hour).Format(time.RFC3339)
	entity1 := api_auth.Entity{
		KeyName:     "watermint",
		Scope:       "toolbox",
		PeerName:    "default",
		Credential:  "*TOP*SECRET*",
		Description: "default connection",
		Timestamp:   ts1,
	}
	entity2 := api_auth.Entity{
		KeyName:     "watermint",
		Scope:       "toolbox",
		PeerName:    "green",
		Credential:  "*TOP*SECRET*GREEN*",
		Description: "green connection",
		Timestamp:   ts2,
	}
	entity3 := api_auth.Entity{
		KeyName:     "stonemint",
		Scope:       "toolbox",
		PeerName:    "default",
		Credential:  "*SECRET*STONEMINT*",
		Description: "stonemint connection",
		Timestamp:   ts3,
	}
	entities := []api_auth.Entity{entity1, entity2, entity3}

	for _, entity := range entities {
		repo.Put(entity)
		e, found := repo.Get(entity.KeyName, entity.Scope, entity.PeerName)
		if !found {
			t.Error("not found", entity)
			return
		}
		if !reflect.DeepEqual(entity, e) {
			t.Error("mismatch")
			return
		}
	}

	elw := repo.List("watermint", "toolbox")
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

	repo.Delete(entity1.KeyName, entity1.Scope, entity1.PeerName)
	_, found := repo.Get(entity1.KeyName, entity1.Scope, entity1.PeerName)
	if found {
		t.Error("deletion failure")
		return
	}
}

func TestNewInMemory(t *testing.T) {
	repo, err := NewInMemory()
	if err != nil {
		t.Error(err)
		return
	}

	repoTestScenario(t, repo)

	repo.Close()
}

func TestNewPersistent(t *testing.T) {
	f, err := qt_file.MakeTestFolder("repo", false)
	if err != nil {
		t.Error(err)
		return
	}
	defer func() {
		_ = os.Remove(f)
	}()

	repo, err := NewPersistent(filepath.Join(f, "repo.db"))
	if err != nil {
		t.Error(err)
		return
	}

	repoTestScenario(t, repo)

	repo.Close()
}
