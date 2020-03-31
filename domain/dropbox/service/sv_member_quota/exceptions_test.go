package sv_member_quota

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEndToEndExceptionsImpl(t *testing.T) {
	qt_api.DoTestBusinessManagement(func(ctx api_context.DropboxApiContext) {
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
		sve := NewExceptions(ctx)
		initialExceptions, err := sve.List()
		if err != nil {
			t.Error(err)
			return
		}

		isTestTargetExceptionInitially := false
		for _, ie := range initialExceptions {
			if ie.TeamMemberId == testee.TeamMemberId {
				isTestTargetExceptionInitially = true
			}
		}

		// Add
		{
			err = sve.Add(testee.TeamMemberId)
			if err != nil {
				t.Error(err)
			}
		}

		// Verify
		{
			exceptions, err := sve.List()
			if err != nil {
				t.Error(err)
			}

			found := false
			for _, e := range exceptions {
				if e.TeamMemberId == testee.TeamMemberId {
					found = true
					break
				}
			}
			if !found {
				t.Error("could not found in exceptions list")
			}
		}

		// Remove
		{
			err = sve.Remove(testee.TeamMemberId)
			if err != nil {
				t.Error(err)
			}
		}

		// Verify
		{
			exceptions, err := sve.List()
			if err != nil {
				t.Error(err)
			}

			found := false
			for _, e := range exceptions {
				if e.TeamMemberId == testee.TeamMemberId {
					found = true
					break
				}
			}
			if found {
				t.Error("the user still in exceptions list", testee.Email)
			}
		}

		// Restore
		{
			if isTestTargetExceptionInitially {
				err = sve.Add(testee.TeamMemberId)
				if err != nil {
					t.Error(err)
				}
			}
		}
	})
}

// mock tests

func TestExceptionsImpl_Add(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := NewExceptions(ctx)
		err := sv.Add("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestExceptionsImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := NewExceptions(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestExceptionsImpl_Remove(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := NewExceptions(ctx)
		err := sv.Remove("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
