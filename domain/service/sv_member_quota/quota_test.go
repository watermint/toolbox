package sv_member_quota

import (
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_test"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/model/mo_member_quota"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"testing"
)

func TestQuotaImpl(t *testing.T) {
	api_test.DoTestBusinessManagement(func(ctx api_context.Context) {
		svm := sv_member.New(ctx)
		members, err := svm.List()
		if err != nil {
			t.Error(err)
			return
		}
		var testee *mo_member.Member
		testee = members[0]
		for _, m := range members {
			if m.Role == "member_only" {
				testee = m
			}
		}

		// Preserve initial state
		svq := NewQuota(ctx)
		initialQuota, err := svq.Resolve(testee.TeamMemberId)
		if err != nil {
			t.Error(err)
			return
		}

		// Set
		{
			q, err := svq.Update(&mo_member_quota.Quota{
				TeamMemberId: testee.TeamMemberId,
				Quota:        123,
			})
			if err != nil {
				t.Error(err)
			}
			if q.Quota != 123 {
				t.Error("invalid")
			}
		}

		// Get
		{
			q, err := svq.Resolve(testee.TeamMemberId)
			if err != nil {
				t.Error(err)
			}
			if q.Quota != 123 {
				t.Error("invalid")
			}
		}

		// Remove
		{
			err := svq.Remove(testee.TeamMemberId)
			if err != nil {
				t.Error(err)
			}
		}

		// Get
		{
			q, err := svq.Resolve(testee.TeamMemberId)
			if err != nil {
				t.Error(err)
			}
			if q.Quota != 0 || !q.IsUnlimited() {
				t.Error("invalid")
			}
		}

		// Restore
		if !initialQuota.IsUnlimited() {
			q, err := svq.Update(initialQuota)
			if err != nil {
				t.Error(err)
			}
			if q.Quota != initialQuota.Quota {
				t.Error("unable to restore")
			}
		}
	})
}
