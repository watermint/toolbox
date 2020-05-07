package sv_member_quota

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

// mock tests

func TestExceptionsImpl_Add(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewExceptions(ctx)
		err := sv.Add("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestExceptionsImpl_List(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewExceptions(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestExceptionsImpl_Remove(t *testing.T) {
	qt_recipe.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := NewExceptions(ctx)
		err := sv.Remove("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
