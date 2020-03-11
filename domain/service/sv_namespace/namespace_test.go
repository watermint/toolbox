package sv_namespace

import (
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_api"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"testing"
)

func TestEndToEndNamespaceImpl_List(t *testing.T) {
	qt_api.DoTestBusinessFile(func(ctx api_context.Context) {
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

// Mock tests

func TestNamespaceImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.Context) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
