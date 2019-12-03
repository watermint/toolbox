package qs_file

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
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
	scenario := &Scenario{}
	if err := scenario.Create(); err != nil {
		t.Error(err)
		return
	}

	dbxBase := "/" + qt_recipe.TestTeamFolderName + "/" + time.Now().Format("2006-01-02T15-04-05")

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
			ch, err := api_util.ContentHash(cols["input.file"])
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
		// `file upload`
		{
			fc, err := app_control_impl.Fork(ctl, "file-upload")
			if err != nil {
				return
			}
			vo := &file.UploadVO{
				LocalPath:   scenario.LocalPath,
				DropboxPath: dbxBase + "/file-upload",
				Overwrite:   false,
			}
			r := file.Upload{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
				t.Error(err)
			}

			testContent(fc, "upload", scenario.LocalPath, dbxBase+"/file-upload")
			testSkip(fc, "skip", scenario.LocalPath)
		}

		// `file sync up`
		{
			fc, err := app_control_impl.Fork(ctl, "file-sync-up")
			if err != nil {
				return
			}
			vo := &filesync.UpVO{
				LocalPath:   scenario.LocalPath,
				DropboxPath: dbxBase + "/file-sync-up",
			}
			r := filesync.Up{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
				t.Error(err)
			}

			testContent(fc, "upload", scenario.LocalPath, dbxBase+"/file-sync-up")
			testSkip(fc, "skip", scenario.LocalPath)
		}

		// `file sync up`: should skip uploading for all files
		{
			fc, err := app_control_impl.Fork(ctl, "file-sync-up-dup")
			if err != nil {
				return
			}
			vo := &filesync.UpVO{
				LocalPath:   scenario.LocalPath,
				DropboxPath: dbxBase + "/file-sync-up",
			}
			r := filesync.Up{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
				t.Error(err)
			}

			qt_recipe.TestRows(fc, "upload", func(cols map[string]string) error {
				t.Error("upload should not contain on 2nd run")
				return errors.New("upload should not contain rows")
			})

			testSkip(fc, "skip", scenario.LocalPath)
		}

		// `file compare local`
		{
			fc, err := app_control_impl.Fork(ctl, "file-compare-local")
			if err != nil {
				return
			}
			vo := &filecompare.LocalVO{
				LocalPath:   scenario.LocalPath,
				DropboxPath: dbxBase + "/file-sync-up",
			}
			r := filecompare.Local{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
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
			vo := &filesyncpreflight.UpVO{
				LocalPath:   scenario.LocalPath,
				DropboxPath: dbxBase + "/file-sync-preflight-up",
			}
			r := filesyncpreflight.Up{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
				t.Error(err)
			}
			testSkip(fc, "skip", scenario.LocalPath)
		}

		// `file list`
		{
			fc, err := app_control_impl.Fork(ctl, "file-list")
			if err != nil {
				return
			}
			vo := &file.ListVO{
				Path:      dbxBase + "/file-sync-up",
				Recursive: true,
			}
			r := file.List{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
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
			vo := &file.CopyVO{
				Src: dbxBase + "/file-sync-up",
				Dst: dbxBase + "/file-copy",
			}
			r := file.Copy{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
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
			vo := &file.MoveVO{
				Src: dbxBase + "/file-copy",
				Dst: dbxBase + "/file-move",
			}
			r := file.Move{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
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
			vo := &file.MergeVO{
				From:   dbxBase + "/file-sync-up",
				To:     dbxBase + "/file-move",
				DryRun: false,
			}
			r := file.Merge{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
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
			vo := &file.DeleteVO{
				Path: dbxBase,
			}
			//r := file.Delete{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			//if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
			//	t.Error(err)
			//}
			// TODO: verify deletion
		}
	})
}
