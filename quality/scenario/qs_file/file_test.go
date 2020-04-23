package qs_file

import (
	"errors"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"github.com/watermint/toolbox/recipe/file"
	filecompare "github.com/watermint/toolbox/recipe/file/compare"
	filesync "github.com/watermint/toolbox/recipe/file/sync"
	filesyncpreflight "github.com/watermint/toolbox/recipe/file/sync/preflight"
	"go.uber.org/zap"
	"path/filepath"
	"testing"
	"time"
)

func TestFileUploadScenario(t *testing.T) {
	l := app_root.Log()

	if qt_endtoend.IsSkipEndToEndTest() {
		l.Info("Skip end to end test")
		return
	}

	scenario := &Scenario{}
	if err := scenario.Create(); err != nil {
		t.Error(err)
		return
	}

	dbxBase := qt_recipe.NewTestDropboxFolderPath(time.Now().Format("2006-01-02T15-04-05"))

	testContent := func(ctl app_control.Control, reportName, localBase, dbxBase string) {
		found := make(map[string]bool)
		contentErr := qt_recipe.TestRows(ctl, reportName, func(cols map[string]string) error {
			if cols["result.content_hash"] == "" {
				l.Debug("ignore folder")
				return nil
			}
			r, err := filepath.Rel(localBase, cols["input.file"])
			if err != nil {
				l.Debug("unable to calc rel path", zap.Error(err))
				return err
			}
			ll := l.With(zap.String("r", r))
			found[r] = true
			ch, err := dbx_util.ContentHash(cols["input.file"])
			if err != nil {
				ll.Debug("unable to calc hash", zap.Error(err))
				return err
			}
			if cols["result.content_hash"] != ch {
				ll.Error("Content hash mismatch", zap.String("hashOnServer", cols["result.content_hash"]), zap.String("hashOnLocal", ch))
				t.Error("content hash mismatch")
			}

			return nil
		})
		if contentErr != nil {
			t.Error(contentErr)
		}

		for f := range scenario.Files {
			if _, ok := found[f]; !ok {
				l.Error("File missing", zap.String("file", f))
				t.Error("missing file")
			}
		}
	}
	testSkip := func(ctl app_control.Control, reportName, localBase string) {
		found := make(map[string]bool)
		skipErr := qt_recipe.TestRows(ctl, reportName, func(cols map[string]string) error {
			r, err := filepath.Rel(localBase, cols["input.file"])
			if err != nil {
				l.Debug("unable to calc rel path", zap.Error(err))
				return err
			}
			found[r] = true
			return nil
		})
		if skipErr != nil {
			t.Error(skipErr)
		}
		for f := range scenario.Ignore {
			if _, ok := found[f]; !ok {
				l.Error("File missing", zap.String("file", f))
				t.Error("missing file")
			}
		}
	}

	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		if _, err := dbx_conn_impl.ConnectTest(api_auth.DropboxTokenFull, app.PeerEndToEndTest, ctl); err != nil {
			l.Info("Skip: no end to end test resource found")
			return
		}

		// `file upload`
		{
			fc, err := app_control_impl.Fork(ctl, "file-upload")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &file.Upload{}, func(r rc_recipe.Recipe) {
				ru := r.(*file.Upload)
				ru.LocalPath = mo_path2.NewFileSystemPath(scenario.LocalPath)
				ru.DropboxPath = dbxBase.ChildPath("file-upload")
				ru.Overwrite = false
			})
			if err != nil {
				t.Error(err)
				return
			}

			testContent(fc, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-upload").Path())
			testSkip(fc, "skipped", scenario.LocalPath)
		}

		// `file sync up`
		{
			fc, err := app_control_impl.Fork(ctl, "file-sync-up")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
				ru := r.(*filesync.Up)
				ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
				ru.DropboxPath = dbxBase.ChildPath("file-sync-up")
			})
			if err != nil {
				t.Error(err)
			}

			testContent(fc, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-sync-up").Path())
			testSkip(fc, "skipped", scenario.LocalPath)
		}

		// `file sync up`: should skip uploading for all files
		{
			fc, err := app_control_impl.Fork(ctl, "file-sync-up-dup")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
				ru := r.(*filesync.Up)
				ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
				ru.DropboxPath = dbxBase.ChildPath("file-sync-up")
			})
			if err != nil {
				t.Error(err)
			}
			qt_recipe.TestRows(fc, "upload", func(cols map[string]string) error {
				t.Error("upload should not contain on 2nd run")
				return errors.New("upload should not contain rows")
			})

			testSkip(fc, "skipped", scenario.LocalPath)
		}

		// `file compare local`
		{
			fc, err := app_control_impl.Fork(ctl, "file-compare-local")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &filecompare.Local{}, func(r rc_recipe.Recipe) {
				rc := r.(*filecompare.Local)
				rc.DropboxPath = dbxBase.ChildPath("file-sync-up")
				rc.LocalPath = mo_path2.NewFileSystemPath(scenario.LocalPath)
			})
			if err != nil {
				t.Error(err)
			}
			// TODO: verify result
		}

		// `file sync preflight up`
		{
			fc, err := app_control_impl.Fork(ctl, "file-sync-preflight-up")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &filesyncpreflight.Up{}, func(r rc_recipe.Recipe) {
				ru := r.(*filesyncpreflight.Up)
				ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
				ru.DropboxPath = dbxBase.ChildPath("file-sync-preflight-up")
			})
			if err != nil {
				t.Error(err)
			}
			testSkip(fc, "skipped", scenario.LocalPath)
		}

		// `file list`
		{
			fc, err := app_control_impl.Fork(ctl, "file-list")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &file.List{}, func(r rc_recipe.Recipe) {
				rc := r.(*file.List)
				rc.Path = dbxBase.ChildPath("file-sync-up")
				rc.Recursive = true
			})
			if err != nil {
				t.Error(err)
			}
			//TODO: verify content
		}

		// `file copy`
		{
			fc, err := app_control_impl.Fork(ctl, "file-copy")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &file.Copy{}, func(r rc_recipe.Recipe) {
				rc := r.(*file.Copy)
				rc.Src = dbxBase.ChildPath("file-sync-up")
				rc.Dst = dbxBase.ChildPath("/file-copy")
			})
			if err != nil {
				t.Error(err)
			}
			//TODO: verify content
		}

		// `file move`
		{
			fc, err := app_control_impl.Fork(ctl, "file-move")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &file.Move{}, func(r rc_recipe.Recipe) {
				rc := r.(*file.Move)
				rc.Src = dbxBase.ChildPath("file-copy")
				rc.Dst = dbxBase.ChildPath("file-move")
			})
			if err != nil {
				t.Error(err)
			}
			//TODO: verify content
		}

		// `file merge`
		{
			fc, err := app_control_impl.Fork(ctl, "file-merge")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &file.Merge{}, func(r rc_recipe.Recipe) {
				rc := r.(*file.Merge)
				rc.From = dbxBase.ChildPath("file-sync-up")
				rc.To = dbxBase.ChildPath("file-move")
				rc.DryRun = false
			})
			if err != nil {
				t.Error(err)
			}
			//TODO: verify content
		}

		// `file delete` for clean up
		{
			fc, err := app_control_impl.Fork(ctl, "file-delete-clean-up")
			if err != nil {
				return
			}
			err = rc_exec.Exec(fc, &file.Delete{}, func(r rc_recipe.Recipe) {
				rc := r.(*file.Delete)
				rc.Path = dbxBase
			})
			if err != nil {
				t.Error(err)
			}
			// TODO: verify deletion
		}
	})
}
