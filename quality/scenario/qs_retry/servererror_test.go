package qs_retry

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_team"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"testing"
)

func TestInternalServerError(t *testing.T) {
	qtr_endtoend.TestWithReplayDbxContext(t, "qs_retry-500.json", func(ctx dbx_client.Client) {
		info, err := sv_team.New(ctx).Info()
		if err != nil {
			t.Error(err)
			return
		}
		if info.Name != "xxxxxxxxx xxx" || info.TeamId != "dbtid:xxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxxx" {
			t.Error(info)
		}
	})
}
