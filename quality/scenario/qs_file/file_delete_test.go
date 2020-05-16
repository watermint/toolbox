package qs_file

import (
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/demo/qdm_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"github.com/watermint/toolbox/recipe/file"
	"github.com/watermint/toolbox/recipe/file/sync"
	"testing"
)

func TestFileFileDelete(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, true, func(scenario qdm_file.Scenario, dbxBase mo_path.DropboxPath) {
			// `file sync up`
			qtr_endtoend.ForkWithName(t, "file-delete1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &sync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*sync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-delete")
				})
				if err != nil {
					t.Error(err)
				}

				testContent(t, fc, scenario, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-sync-up").Path())
				testSkip(t, fc, scenario, "skipped", scenario.LocalPath)
				return nil
			})

			// `file delete` for clean up
			qtr_endtoend.ForkWithName(t, "file-delete2", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &file.Delete{}, func(r rc_recipe.Recipe) {
					rc := r.(*file.Delete)
					rc.Path = dbxBase.ChildPath("file-delete")
				})
				if err != nil {
					t.Error(err)
				}
				// TODO: verify deletion
				return nil
			})
		})
	})
}
