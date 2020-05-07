package sv_sharedlink

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
	"time"
)

// Mock tests

func TestSharedLinkImpl_Create(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx)
		_, err := sv.Create(qtr_endtoend.NewTestDropboxFolderPath(), Public())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		_, err = sv.Create(qtr_endtoend.NewTestDropboxFolderPath(), TeamOnly())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
		_, err = sv.Create(qtr_endtoend.NewTestDropboxFolderPath(), Expires(time.Now()))
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_List(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_ListByPath(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx)
		_, err := sv.ListByPath(qtr_endtoend.NewTestDropboxFolderPath())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_Remove(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx)
		err := sv.Remove(&mo_sharedlink.Metadata{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_Resolve(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx)
		_, err := sv.Resolve(mo_url.NewEmptyUrl(), "test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestSharedLinkImpl_Update(t *testing.T) {
	qtr_endtoend.TestWithDbxContext(t, func(ctx dbx_context.Context) {
		sv := New(ctx)
		_, err := sv.Update(&mo_sharedlink.Metadata{}, RemoveExpiration())
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
