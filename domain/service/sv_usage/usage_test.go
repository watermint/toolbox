package sv_usage

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"testing"
)

func TestUsageImpl_Resolve(t *testing.T) {
	qt_api.DoTestTokenFull(func(ctx api_context.Context) {
		_, err := New(ctx).Resolve()
		if err != nil {
			t.Error(err)
		}
	})
}
