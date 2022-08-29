package sv_sharedfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

// Mock tests

func TestSharedFolderImpl_Create(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Create(qtr_endtoend.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_Leave(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		err := sv.Leave(&mo_sharedfolder.SharedFolder{}, LeaveACopy(true))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_Remove(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		err := sv.Remove(&mo_sharedfolder.SharedFolder{}, LeaveACopy(true))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_Resolve(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Resolve("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_Transfer(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		err := sv.Transfer(&mo_sharedfolder.SharedFolder{}, ToProfile(&mo_profile.Profile{}))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		err = sv.Transfer(&mo_sharedfolder.SharedFolder{}, ToAccountId("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		err = sv.Transfer(&mo_sharedfolder.SharedFolder{}, ToTeamMemberId("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedFolderImpl_UpdatePolicy(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.UpdatePolicy("test", MemberPolicy("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
