package sync

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_filter"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"github.com/watermint/toolbox/ingredient/ig_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Up struct {
	rc_recipe.RemarkIrreversible
	Peer        dbx_conn.ConnScopedIndividual
	LocalPath   mo_path2.ExistingFileSystemPath
	DropboxPath mo_path.DropboxPath
	Upload      *ig_file.Upload
	Overwrite   bool
	Delete      bool
	BatchSize   mo_int.RangeInt
	Name        mo_filter.Filter
}

func (z *Up) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesContentWrite,
	)
	z.BatchSize.SetRange(1, 1000, 50)
	z.Name.SetOptions(
		mo_filter.NewNameFilter(),
		mo_filter.NewNameSuffixFilter(),
		mo_filter.NewNamePrefixFilter(),
		mo_file_filter.NewIgnoreFileFilter(),
	)
}

func (z *Up) Exec(c app_control.Control) error {
	return rc_exec.Exec(c, z.Upload, func(r rc_recipe.Recipe) {
		ru := r.(*ig_file.Upload)
		ru.LocalPath = z.LocalPath
		ru.DropboxPath = z.DropboxPath
		ru.Overwrite = z.Overwrite
		ru.Name = z.Name
		ru.Context = z.Peer.Client()
		ru.BatchSize = z.BatchSize.Value()
		ru.Delete = z.Delete
	})
}

func (z *Up) Test(c app_control.Control) error {
	// #470 : -delete option may not work
	{
		l := c.Log()

		sourceTree, err := qt_file.MakeTestFolder("source", true)
		if err != nil {
			return err
		}
		defer func() {
			_ = os.RemoveAll(sourceTree)
		}()
		targetTree, err := qt_file.MakeTestFolder("target", false)
		if err != nil {
			return err
		}
		defer func() {
			_ = os.RemoveAll(targetTree)
		}()

		dbxPath := mo_path.NewDropboxPath(qtr_endtoend.TestTeamFolderName).ChildPath("file-upload_" + sc_random.MustGetSecureRandomString(4))

		cs, err := app_control_impl.ForkQuiet(c, "source")
		if err != nil {
			return err
		}
		ct, err := app_control_impl.ForkQuiet(c, "target")
		if err != nil {
			return err
		}

		// upload source state (with test file)
		err = rc_exec.Exec(cs, &Up{}, func(r rc_recipe.Recipe) {
			m := r.(*Up)
			m.LocalPath = mo_path2.NewExistingFileSystemPath(sourceTree)
			m.DropboxPath = dbxPath
			m.Delete = true
		})
		if resolvedErr, cont := qt_errors.ErrorsForTest(c.Log(), err); !cont {
			l.Debug("Error on source copy", esl.Error(err))
			return resolvedErr
		}

		// upload target state
		err = rc_exec.Exec(ct, &Up{}, func(r rc_recipe.Recipe) {
			m := r.(*Up)
			m.LocalPath = mo_path2.NewExistingFileSystemPath(targetTree)
			m.DropboxPath = dbxPath
			m.Delete = true
		})
		if resolvedErr, cont := qt_errors.ErrorsForTest(c.Log(), err); !cont {
			l.Debug("Error on target copy", esl.Error(err))
			return resolvedErr
		}

		targetReportPath := filepath.Join(ct.Workspace().Report(), "summary.json")
		targetReportFile, err := ioutil.ReadFile(targetReportPath)
		if err != nil {
			l.Debug("Unable to retrieve report", esl.Error(err))
			return err
		}

		summary := ig_file.Summary{}
		if err := json.Unmarshal(targetReportFile, &summary); err != nil {
			l.Debug("Unable to unmarshal report", esl.Error(err))
			return err
		}

		if summary.NumDeleted != 1 {
			l.Debug("Delete failed", esl.Any("report", summary))
			return errors.New("failed sync deletes")
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
