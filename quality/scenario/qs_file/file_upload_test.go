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
	"testing"
)

func TestFileUpload(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, false, func(scenario qdm_file.Scenario, dbxBase mo_path.DropboxPath) {
			qtr_endtoend.ForkWithName(t, "file-upload", ctl, func(c app_control.Control) error {
				err := rc_exec.Exec(c, &file.Upload{}, func(r rc_recipe.Recipe) {
					ru := r.(*file.Upload)
					ru.LocalPath = mo_path2.NewFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-upload")
					ru.Overwrite = false
				})
				if err != nil {
					t.Error(err)
					return nil
				}

				testContent(t, c, scenario, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-upload").Path())
				testSkip(t, c, scenario, "skipped", scenario.LocalPath)
				return nil
			})
		})
	})
}
