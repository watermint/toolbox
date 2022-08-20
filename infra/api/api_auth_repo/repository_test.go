package api_auth_repo

import (
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
	"reflect"
	"testing"
)

func repoTestScenario(t *testing.T, repo api_auth.Repository) {
	entity1 := api_auth.Entity{
		KeyName:     "watermint",
		Scope:       "toolbox",
		PeerName:    "default",
		Credential:  "*TOP*SECRET*",
		Description: "default connection",
	}
	entity2 := api_auth.Entity{
		KeyName:     "watermint",
		Scope:       "toolbox",
		PeerName:    "green",
		Credential:  "*TOP*SECRET*GREEN*",
		Description: "green connection",
	}
	entity3 := api_auth.Entity{
		KeyName:     "stonemint",
		Scope:       "toolbox",
		PeerName:    "default",
		Credential:  "*SECRET*STONEMINT*",
		Description: "stonemint connection",
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
