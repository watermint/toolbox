package sync

import (
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/file"
	"github.com/watermint/toolbox/quality/demo/qdm_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Up struct {
	rc_recipe.RemarkIrreversible
	Peer        dbx_conn.ConnUserFile
	LocalPath   mo_path2.ExistingFileSystemPath
	DropboxPath mo_path.DropboxPath
	ChunkSizeKb mo_int.RangeInt
	Upload      *file.Upload
	FailOnError bool
}

func (z *Up) Preset() {
	z.ChunkSizeKb.SetRange(1, 150*1024, 64*1024)
}

func (z *Up) Exec(c app_control.Control) error {
	l := c.Log()
	err := rc_exec.Exec(c, z.Upload, func(r rc_recipe.Recipe) {
		ru := r.(*file.Upload)
		ru.EstimateOnly = false
		ru.LocalPath = z.LocalPath
		ru.DropboxPath = z.DropboxPath
		ru.Overwrite = true
		ru.CreateFolder = true
		ru.Context = z.Peer.Context()
		ru.ChunkSizeKb = z.ChunkSizeKb.Value()
	})
	if z.FailOnError && err != nil {
		l.Debug("Return error", esl.Error(err))
		return err
	}
	return nil
}

func (z *Up) Test(c app_control.Control) error {
	// replay test
	{
		sc, err := qdm_file.NewScenario(false)
		if err != nil {
			return err
		}
		err = rc_exec.ExecReplay(c, &Up{}, "recipe-file-sync-up.json.gz", func(r rc_recipe.Recipe) {
			m := r.(*Up)
			m.LocalPath = mo_path2.NewExistingFileSystemPath(sc.LocalPath)
			m.DropboxPath = qtr_endtoend.NewTestDropboxFolderPath("file-sync-up")
		})
		if err != nil {
			return err
		}
	}

	err := rc_exec.ExecMock(c, &Up{}, func(r rc_recipe.Recipe) {
		m := r.(*Up)
		m.LocalPath = qtr_endtoend.NewTestExistingFileSystemFolderPath(c, "up")
		m.DropboxPath = qtr_endtoend.NewTestDropboxFolderPath("up")
	})
	if err, _ = qt_errors.ErrorsForTest(c.Log(), err); err != nil {
		return err
	}

	return qt_errors.ErrorScenarioTest
}
