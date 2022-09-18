package sv_sharedfolder_mount

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

// mock test

func TestMountImpl_List(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.List()
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMountImpl_Mount(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		_, err := sv.Mount(&mo_sharedfolder.SharedFolder{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestMountImpl_Unmount(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := New(ctx)
		err := sv.Unmount(&mo_sharedfolder.SharedFolder{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
