package sv_usage

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

// Mock tests

func TestUsageImpl_Resolve(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx)
		_, err := sv.Resolve()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})

}
