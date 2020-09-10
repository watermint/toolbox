package qs_file

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/demo/qdm_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"github.com/watermint/toolbox/recipe/file/sync"
	"testing"
)

func TestFileSyncUp(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, false, func(scenario qdm_file.Scenario, dbxBase mo_path.DropboxPath) {
			qtr_endtoend.ForkWithName(t, "file-sync-up", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &sync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*sync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-sync-up")
				})
				if err != nil {
					t.Error(err)
				}

				testContent(t, fc, scenario, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-sync-up").Path())
				testSkip(t, fc, scenario, "skipped", scenario.LocalPath)
				return nil
			})
		})
	})
}

func TestFileSyncUpDup(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, false, func(scenario qdm_file.Scenario, dbxBase mo_path.DropboxPath) {
			// `file sync up`
			qtr_endtoend.ForkWithName(t, "file-sync-up-dup1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &sync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*sync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-sync-up-dup")
				})
				if err != nil {
					t.Error(err)
				}

				testContent(t, fc, scenario, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-sync-up-dup").Path())
				testSkip(t, fc, scenario, "skipped", scenario.LocalPath)
				return nil
			})

			// `file sync up`: should skip uploading for all files
			qtr_endtoend.ForkWithName(t, "file-sync-up-dup2", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &sync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*sync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-sync-up-dup")
				})
				if err != nil {
					t.Error(err)
				}
				qtr_endtoend.TestRows(fc, "upload", func(cols map[string]string) error {
					t.Error("upload should not contain on 2nd run")
					return errors.New("upload should not contain rows")
				})
				testSkip(t, fc, scenario, "skipped", scenario.LocalPath)
				return nil
			})
		})
	})
}
