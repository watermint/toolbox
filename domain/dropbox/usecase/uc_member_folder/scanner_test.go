package uc_member_folder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestScanImpl_Scan(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		qtr_endtoend.TestWithDbxClient(t, func(ctx dbx_client.Client) {
			s := New(ctl, ctx)
			_, err := s.Scan(mo_filter.New(""))
			if err != qt_errors.ErrorMock && err != nil {
				t.Error(err)
			}
		})
	})
}
