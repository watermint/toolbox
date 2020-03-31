package sv_profile

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEndToEndProfileImpl_Current(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.DropboxApiContext) {
		svc := NewProfile(ctx)
		prof, err := svc.Current()
		if err != nil {
			t.Error(err)
		}
		if prof.Email == "" || prof.AccountId == "" {
			t.Error("invalid")
		}
	})
}

func TestEndToEndTeamImpl_Admin(t *testing.T) {
	qt_api.DoTestBusinessInfo(func(ctx api_context.DropboxApiContext) {
		svc := NewTeam(ctx)
		prof, err := svc.Admin()
		if err != nil {
			t.Error(err)
		}
		if prof.Email == "" || prof.AccountId == "" {
			t.Error("invalid")
		}
	})
}

// Mock tests

func TestProfileImpl_Current(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := NewProfile(ctx)
		_, err := sv.Current()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestTeamImpl_Admin(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := NewTeam(ctx)
		_, err := sv.Admin()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
