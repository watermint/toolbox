package qs_file

import (
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/demo/qdm_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"github.com/watermint/toolbox/recipe/file/sync/preflight"
	"testing"
)

func TestFileSyncPreflightUp(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, true, func(scenario qdm_file.Scenario, dbxBase mo_path.DropboxPath) {
			qtr_endtoend.ForkWithName(t, "file-sync-preflight-up", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &preflight.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*preflight.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-sync-preflight-up")
				})
				if err != nil {
					t.Error(err)
				}
				testSkip(t, fc, scenario, "skipped", scenario.LocalPath)
				return nil
			})
		})
	})
}
