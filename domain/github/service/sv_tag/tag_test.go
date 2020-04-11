package sv_tag

import (
	"github.com/watermint/toolbox/domain/github/api/gh_context_impl"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"testing"
)

func TestTagImpl_List(t *testing.T) {
	mc := &gh_context_impl.Mock{}
	sv := New(mc, "watermint", "toolbox")
	if _, err := sv.List(); err != qt_errors.ErrorMock {
		t.Error(err)
	}
}

func TestTagImpl_Create(t *testing.T) {
	mc := &gh_context_impl.Mock{}
	sv := New(mc, "watermint", "toolbox_sandbox")
	if _, err := sv.Create("v1.1.1", "testing", "4e1243bd22c66e76c2ba9eddc1f91394e57f9f83"); err != qt_errors.ErrorMock {
		t.Error(err)
	}
}
