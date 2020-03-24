package sv_activity

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_activity"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

func TestActivityImpl_List(t *testing.T) {
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
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
	qt_recipe.TestWithApiContext(t, func(ctx api_context.DropboxApiContext) {
		sv := New(ctx)
		err := sv.All(func(event *mo_activity.Event) error {
			return nil
		})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
