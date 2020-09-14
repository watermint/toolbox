package qs_file

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/demo/qdm_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"github.com/watermint/toolbox/recipe/file"
	"github.com/watermint/toolbox/recipe/file/sync"
	"testing"
)

func TestFileFileMerge(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, false, func(scFrom qdm_file.Scenario, dbxBase mo_path.DropboxPath) {
			scTo, err := qdm_file.NewScenario(true)
			if err != nil {
				t.Error(err)
				return
			}

			// `file sync up`: scFrom
			qtr_endtoend.ForkWithName(t, "file-merge1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &sync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*sync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scFrom.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-merge-from")
				})
				if err != nil {
					t.Error(err)
				}

				testContent(t, fc, scFrom, "uploaded", scFrom.LocalPath, dbxBase.ChildPath("file-merge-from").Path())
				testSkip(t, fc, scFrom, "skipped", scFrom.LocalPath)
				return nil
			})

			// `file sync up`: scTo
			qtr_endtoend.ForkWithName(t, "file-merge2", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &sync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*sync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scTo.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-merge-to")
				})
				if err != nil {
					t.Error(err)
				}

				testContent(t, fc, scTo, "uploaded", scTo.LocalPath, dbxBase.ChildPath("file-merge-to").Path())
				testSkip(t, fc, scTo, "skipped", scTo.LocalPath)
				return nil
			})

			// `file merge`
			qtr_endtoend.ForkWithName(t, "file-merge2", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &file.Merge{}, func(r rc_recipe.Recipe) {
					rc := r.(*file.Merge)
					rc.From = dbxBase.ChildPath("file-merge-from")
					rc.To = dbxBase.ChildPath("file-merge-to")
					rc.DryRun = false
				})
				if err != nil {
					t.Error(err)
				}
				//TODO: verify content
				return nil
			})
		})
	})
}
