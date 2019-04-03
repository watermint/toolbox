package sv_group

import (
	"fmt"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"testing"
	"time"
)

func TestImplGroup_CreateRemove(t *testing.T) {
	api_test.DoTestBusinessManagement(func(ctx api_context.Context) {
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

func TestImplGroup_List(t *testing.T) {
	api_test.DoTestBusinessManagement(func(ctx api_context.Context) {
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
