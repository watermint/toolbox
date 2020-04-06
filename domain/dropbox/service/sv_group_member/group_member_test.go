package sv_group_member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEndToEndGroupMemberImpl_List(t *testing.T) {
	qt_api.DoTestBusinessManagement(func(ctx dbx_context.Context) {
		gsv := sv_group.New(ctx)
		groups, err := gsv.List()
		if err != nil {
			t.Error(err)
			return
		}

		for i, group := range groups {
			if i > 10 {
				break
			}
			msv := New(ctx, group)
			members, err := msv.List()
			if err != nil {
				t.Error(err)
			}
			for _, m := range members {
				if m.TeamMemberId == "" || m.AccessType == "" {
					t.Error("invalid")
				}
			}
		}
	})
}

func TestGroupMemberImpl_Add(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx, &mo_group.Group{})
		_, err := sv.Add(ByEmail("test@example.com"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestGroupMemberImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx, &mo_group.Group{})
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestGroupMemberImpl_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx, &mo_group.Group{})
		_, err := sv.Remove(ByTeamMemberId("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
