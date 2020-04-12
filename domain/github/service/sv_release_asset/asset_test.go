package sv_release_asset

import (
	"github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"testing"
)

func TestAssetImpl_Upload(t *testing.T) {
	mc := &gh_context_impl.Mock{}
	sv := New(mc, "watermint", "toolbox", "25040282")
	fp, err := qt_file.MakeTestFile("test.txt", "hello this is test")
	if err != nil {
		t.Error(err)
		return
	}
	ef := mo_path.NewExistingFileSystemPath(fp)
	if _, err := sv.Upload(ef); err != qt_errors.ErrorMock {
		t.Error(err)
	}
}

func TestAssetImpl_List(t *testing.T) {
	mc := &gh_context_impl.Mock{}
	sv := New(mc, "watermint", "toolbox", "25040282")
	if _, err := sv.List(); err != qt_errors.ErrorMock {
		t.Error(err)
	}
}
