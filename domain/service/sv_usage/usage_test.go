package sv_usage

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_test"
	"testing"
)

func TestUsageImpl_Resolve(t *testing.T) {
	api_test.DoTestTokenFull(func(ctx api_context.Context) {
		_, err := New(ctx).Resolve()
		if err != nil {
			t.Error(err)
		}
	})
}
