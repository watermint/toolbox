package sv_namespace

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/api/api_test"
	"go.uber.org/zap"
	"testing"
)

func TestNamespaceImpl_List(t *testing.T) {
	api_test.DoTestBusinessFile(func(ctx api_context.Context) {
		svc := newTest(ctx, 3)
		namespaces, err := svc.List()
		if err != nil {
			t.Error(err)
			return
		}
		if len(namespaces) < 1 {
			t.Error("invalid")
			return
		}
		for _, n := range namespaces {
			if n.NamespaceId == "" {
				t.Error("invalid")
			}
			ctx.Log().Debug("namespace", zap.Any("namespace", n))
		}
	})
}
