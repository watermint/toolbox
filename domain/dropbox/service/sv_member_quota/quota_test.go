package sv_member_quota

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member_quota"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

// mock tests

func TestQuotaImpl_Remove(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewQuota(ctx)
		err := sv.Remove("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestQuotaImpl_Resolve(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewQuota(ctx)
		_, err := sv.Resolve("test")
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}

func TestQuotaImpl_Update(t *testing.T) {
	qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
		sv := NewQuota(ctx)
		_, err := sv.Update(&mo_member_quota.Quota{})
		if err != nil && err != qt_errors.ErrorMock {
			t.Error(err)
		}
	})
}
