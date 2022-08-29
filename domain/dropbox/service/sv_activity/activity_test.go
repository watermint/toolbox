package sv_activity

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestActivityImpl_List(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		err := sv.List(func(event *mo_activity.Event) error {
			return nil
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestActivityImpl_All(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		err := sv.All(func(event *mo_activity.Event) error {
			return nil
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
