package dbx_auth_test

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
	"time"
)

func TestCached_Auth(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		name := "test-cached-auth-" + time.Now().String()
		a := dbx_auth.NewConsoleCacheOnly(ctl, name)
		_, err := a.Auth("test-cached-auth")
		if err == nil {
			// should not exist
			t.Error("invalid")
		}

		a = dbx_auth.NewConsoleCache(ctl, dbx_auth.NewMock(name))
		_, err = a.Auth(api_auth.DropboxTokenBusinessInfo)
		if err != nil {
			// should exist
			t.Error(err)
		}

		if aa, ok := a.(*dbx_auth.Cached); !ok {
			t.Error("invalid")
		} else {
			aa.Purge(api_auth.DropboxTokenBusinessInfo)
		}

		a = dbx_auth.NewConsoleCacheOnly(ctl, name)
		_, err = a.Auth("test-cached-auth")
		if err == nil {
			// should not exist
			t.Error("invalid")
		}

		if aa, ok := a.(*dbx_auth.Cached); !ok {
			t.Error("invalid")
		} else {
			aa.Purge(api_auth.DropboxTokenBusinessInfo)
		}
	})
}
