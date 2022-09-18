package sv_profile

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

// Mock tests

func TestProfileImpl_Current(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewProfile(ctx)
		_, err := sv.Current()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestTeamImpl_Admin(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewTeam(ctx)
		_, err := sv.Admin()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
