package sv_sharedfolder_member

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_profile"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEndToEndMemberImpl_List(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		sf := sv_sharedfolder.New(ctx)
		folders, err := sf.List()
		if err != nil {
			t.Error(err)
			return
		}

		for _, folder := range folders {
			sm := New(ctx, folder)
			members, err := sm.List()
			if err != nil {
				t.Error(err)
				return
			}

			for _, m := range members {
				if m.MemberType() == "" {
					t.Error("invalid")
				}
			}
		}
	})
}

// Mock tests

func TestMemberImpl_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx, &mo_sharedfolder.SharedFolder{})
		err := sv.Remove(RemoveByEmail("test@example.com"), LeaveACopy())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		err = sv.Remove(RemoveByGroup(&mo_group.Group{}), LeaveACopy())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		err = sv.Remove(RemoveByGroupId("test"), LeaveACopy())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		err = sv.Remove(RemoveByProfile(&mo_profile.Profile{}), LeaveACopy())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		err = sv.Remove(RemoveByTeamMemberId("test"), LeaveACopy())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewByTeamFolder(ctx, &mo_teamfolder.TeamFolder{})
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMemberImpl_Add(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx, &mo_sharedfolder.SharedFolder{})
		err := sv.Add(AddByEmail("test@example.com", LevelEditor),
			AddQuiet(),
			AddCustomMessage("test"),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}

		err = sv.Add(AddByGroup(&mo_group.Group{}, LevelViewer),
			AddQuiet(),
			AddCustomMessage("test"),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}

		err = sv.Add(AddByGroupId("test", LevelViewerNoComment),
			AddQuiet(),
			AddCustomMessage("test"),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}

		err = sv.Add(AddByProfile(&mo_profile.Profile{}, LevelEditor),
			AddQuiet(),
			AddCustomMessage("test"),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}

		err = sv.Add(AddByTeamMemberId("test", LevelOwner),
			AddQuiet(),
			AddCustomMessage("test"),
		)
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

// Mock test : cached

func TestCachedMember_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx, "test")
		err := sv.Remove(RemoveByTeamMemberId("test"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx, "test")
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedMember_Add(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx, "test")
		err := sv.Add(AddByEmail("test", LevelEditor))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
