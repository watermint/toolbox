package qs_file

import (
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
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func TestFileUploadScenario(t *testing.T) {
	l := app_root.Log()

	// path -> content
	vFile := make(map[string]string)
	vFile["123.txt"] = "123"
	vFile["abc.txt"] = "abc"
	vFile["あいう.txt"] = "あいう"
	vFile["time.txt"] = time.Now().String()
	vFile["987/654.txt"] = "654"
	vFile["zyx/wvu.txt"] = "wvu"
	vFile["アイウ/エオ.txt"] = "エオ"
	vFile["a-b-c/time.txt"] = time.Now().String()

	vIgnore := make(map[string]string)
	vIgnore[".DS_Store"] = "ignore-dsstore"
	vIgnore["987/~$abc"] = "ignore-abc"
	vIgnore["d-e-f/.~abc"] = "ignore-dot-tilde"
	vIgnore["~123.tmp"] = "ignore-123"

	// Empty folders
	vFolder := make(map[string]bool)
	vFolder["987"] = true
	vFolder["zyx"] = true
	vFolder["アイウ"] = true
	vFolder["a-b-c"] = true
	vFolder["d-e-f"] = true
	vFolder["1-2-3"] = true
	vFolder["g-h-i/j-k-l"] = true

	localBase, err := ioutil.TempDir("", "file-upload-scenario")
	if err != nil {
		l.Error("unable to create temp dir", zap.Error(err))
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

		for f := range vFile {
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
		for f := range vIgnore {
			if _, ok := found[f]; !ok {
				l.Error("File missing", zap.String("file", f))
				t.Error("missing file")
			}
		}
	}

	// Create test folders
	{
		for f := range vFolder {
			if err := os.MkdirAll(filepath.Join(localBase, f), 0755); err != nil {
				l.Error("Unable to create folder", zap.Error(err), zap.String("f", f))
				t.Error(err)
				return
			}
		}
	}

	// Create test files
	{
		for f, c := range vFile {
			if err := ioutil.WriteFile(filepath.Join(localBase, f), []byte(c), 0644); err != nil {
				l.Error("Unable to create file", zap.Error(err), zap.String("f", f))
				t.Error(err)
				return
			}
		}
		for f, c := range vIgnore {
			if err := ioutil.WriteFile(filepath.Join(localBase, f), []byte(c), 0644); err != nil {
				l.Error("Unable to create file", zap.Error(err), zap.String("f", f))
				t.Error(err)
				return
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
				LocalPath:   localBase,
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

			testContent(fc, "upload", localBase, dbxBase+"/file-upload")
			testSkip(fc, "skip", localBase)
		}

		// `file sync up`
		{
			fc, err := app_control_impl.Fork(ctl, "file-sync-up")
			if err != nil {
				return
			}
			vo := &filesync.UpVO{
				LocalPath:   localBase,
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

			testContent(fc, "upload", localBase, dbxBase+"/file-sync-up")
			testSkip(fc, "skip", localBase)
		}

		// `file compare local`
		{
			fc, err := app_control_impl.Fork(ctl, "file-compare-local")
			if err != nil {
				return
			}
			vo := &filecompare.LocalVO{
				LocalPath:   localBase,
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
				LocalPath:   localBase,
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
			testSkip(fc, "skip", localBase)
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

		// `file copy`
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

		// `file delete`
		{
			fc, err := app_control_impl.Fork(ctl, "file-delete")
			if err != nil {
				return
			}
			vo := &file.DeleteVO{
				Path: dbxBase,
			}
			r := file.Delete{}
			if !qt_recipe.ApplyTestPeers(fc, vo) {
				l.Warn("Skip: No conn resource")
				return
			}
			if err := r.Exec(app_kitchen.NewKitchen(fc, vo)); err != nil {
				t.Error(err)
			}
			// TODO: verify deletion
		}
	})
}
