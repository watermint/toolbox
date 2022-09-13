package api_auth_basic_test

import (
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_basic"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestConsole(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		session := api_auth_basic.NewConsole(ctl)
		entity, err := session.Start(api_auth.BasicSessionData{
			AppData: api_auth.BasicAppData{
				AppKeyName:      "watermint",
				DontUseUsername: false,
				DontUsePassword: false,
			},
			PeerName: "default",
		})
		// should fail
		if err != app.ErrorUserCancelled {
			t.Error(entity)
		}
	})
}
