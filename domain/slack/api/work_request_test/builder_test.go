package work_request_test

import (
	"github.com/watermint/toolbox/domain/slack/api/work_request"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestBuilderImpl_Build(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		builder := work_request.New(ctl, nil)
		_, err := builder.Build()
		if err != nil {
			t.Error(err)
		}
	})
}
