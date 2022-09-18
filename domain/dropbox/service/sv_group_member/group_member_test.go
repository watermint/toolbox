package sv_group_member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestGroupMemberImpl_Add(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx, &mo_group.Group{})
		_, err := sv.Add(ByEmail("test@example.com"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestGroupMemberImpl_List(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx, &mo_group.Group{})
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestGroupMemberImpl_Remove(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx, &mo_group.Group{})
		_, err := sv.Remove(ByTeamMemberId("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
