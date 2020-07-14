package api_auth_impl_test

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
	"time"
)

func TestCached_Auth(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		name := "test-cached-auth-" + time.Now().String()
		a := api_auth_impl.NewConsoleCacheOnly(ctl, name, dbx_auth.NewLegacyApp(ctl))
		_, err := a.Auth([]string{"test-cached-auth"})
		if err == nil {
			// should not exist
			t.Error("invalid")
		}

		a = api_auth_impl.NewConsoleCache(ctl, dbx_auth.NewMock(name), dbx_auth.NewLegacyApp(ctl))
		_, err = a.Auth([]string{api_auth.DropboxTokenBusinessInfo})
		if err != nil {
			// should exist
			t.Error(err)
		}

		if aa, ok := a.(*api_auth_impl.Cached); !ok {
			t.Error("invalid")
		} else {
			aa.Purge(api_auth.DropboxTokenBusinessInfo)
		}

		a = api_auth_impl.NewConsoleCacheOnly(ctl, name, dbx_auth.NewLegacyApp(ctl))
		_, err = a.Auth([]string{"test-cached-auth"})
		if err == nil {
			// should not exist
			t.Error("invalid")
		}

		if aa, ok := a.(*api_auth_impl.Cached); !ok {
			t.Error("invalid")
		} else {
			aa.Purge(api_auth.DropboxTokenBusinessInfo)
		}
	})
}
