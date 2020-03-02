package sv_profile

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"testing"
)

func TestProfileImpl_Current(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
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

func TestTeamImpl_Admin(t *testing.T) {
	qt_api.DoTestBusinessInfo(func(ctx api_context.Context) {
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
