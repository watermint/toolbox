package sv_usage

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestEndToEndUsageImpl_Resolve(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		_, err := New(ctx).Resolve()
		if err != nil {
			t.Error(err)
		}
	})
}

// Mock tests

func TestUsageImpl_Resolve(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.Resolve()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})

}
