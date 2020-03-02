package sv_group

import (
	"fmt"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
	"time"
)

func TestEndToEndImplGroup_CreateRemove(t *testing.T) {
	qt_api.DoTestBusinessManagement(func(ctx api_context.Context) {
		svc := New(ctx)
		name := fmt.Sprintf("toolbox-test-%x", time.Now().Unix())
		createdGroup, err := svc.Create(name, CompanyManaged())
		if err != nil {
			t.Error(err)
			return
		}

		resolvedGroup, err := svc.Resolve(createdGroup.GroupId)
		if err != nil {
			t.Error(err)
			return
		}
		if resolvedGroup.GroupId != createdGroup.GroupId ||
			resolvedGroup.GroupName != createdGroup.GroupName ||
			resolvedGroup.GroupManagementType != createdGroup.GroupManagementType ||
			resolvedGroup.GroupExternalId != createdGroup.GroupExternalId {
			t.Error("invalid")
		}

		err = svc.Remove(createdGroup.GroupId)
		if err != nil {
			t.Error(err)
			return
		}

		if rg, err := svc.Resolve(createdGroup.GroupId); err == nil || rg != nil {
			t.Error("invalid")
			return
		}
	})
}

func TestEndToEndImplGroup_List(t *testing.T) {
	qt_api.DoTestBusinessManagement(func(ctx api_context.Context) {
		svc := New(ctx)
		groups, err := svc.List()
		if err != nil {
			t.Error(err)
			return
		}

		for _, g := range groups {
			resolvedGroup, err := svc.Resolve(g.GroupId)
			if err != nil {
				t.Error(err)
			}
			if resolvedGroup.GroupId != g.GroupId ||
				resolvedGroup.GroupName != g.GroupName ||
				resolvedGroup.GroupManagementType != g.GroupManagementType ||
				resolvedGroup.GroupExternalId != g.GroupExternalId {
				t.Error("invalid", g)
			}
		}
	})
}

// ---- Mock tests for Cache

func TestCachedGroup_Create(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		_, err := sv.Create("test", CompanyManaged())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedGroup_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedGroup_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		err := sv.Remove("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedGroup_Resolve(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		_, err := sv.Resolve("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedGroup_ResolveByName(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		_, err := sv.ResolveByName("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestCachedGroup_Update(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := NewCached(ctx)
		_, err := sv.Update(&mo_group.Group{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

/// ----- Mock tests for impl

func TestImplGroup_Create(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Create("test", ManagementType("company_managed"))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestImplGroup_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})

}

func TestImplGroup_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		err := sv.Remove("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestImplGroup_Resolve(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Resolve("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestImplGroup_ResolveByName(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.ResolveByName("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestImplGroup_Update(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Update(&mo_group.Group{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
