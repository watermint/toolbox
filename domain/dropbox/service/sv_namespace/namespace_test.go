package sv_namespace

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

// Mock tests

func TestNamespaceImpl_List(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
