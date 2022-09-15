package api_auth_basic

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"reflect"
	"testing"
	"time"
)

func TestEmbeddedImpl_Start(t *testing.T) {
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
	session := NewEmbedded(entity1)

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

	entity3, err := session.Start(api_auth.BasicSessionData{
		AppData: api_auth.BasicAppData{
			AppKeyName:      "must_fail",
			DontUseUsername: false,
			DontUsePassword: false,
		},
		PeerName: "must_fail",
	})
	if err == nil {
		t.Error(entity3)
	}
}
