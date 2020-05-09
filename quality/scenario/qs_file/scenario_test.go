package qs_file

import (
	"errors"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"github.com/watermint/toolbox/recipe/file"
	filecompare "github.com/watermint/toolbox/recipe/file/compare"
	filesync "github.com/watermint/toolbox/recipe/file/sync"
	filesyncpreflight "github.com/watermint/toolbox/recipe/file/sync/preflight"
	"path/filepath"
	"testing"
	"time"
)

func testContent(t *testing.T, ctl app_control.Control, scenario Scenario, reportName, localBase, dbxBase string) {
	l := ctl.Log()
	found := make(map[string]bool)
	contentErr := qtr_endtoend.TestRows(ctl, reportName, func(cols map[string]string) error {
		if cols["result.content_hash"] == "" {
			l.Debug("ignore folder")
			return nil
		}
		r, err := filepath.Rel(localBase, cols["input.file"])
		if err != nil {
			l.Debug("unable to calc rel path", esl.Error(err))
			return err
		}
		ll := l.With(esl.String("r", r))
		found[r] = true
		ch, err := dbx_util.ContentHash(cols["input.file"])
		if err != nil {
			ll.Debug("unable to calc hash", esl.Error(err))
			return err
		}
		if cols["result.content_hash"] != ch {
			ll.Error("Content hash mismatch", esl.String("hashOnServer", cols["result.content_hash"]), esl.String("hashOnLocal", ch))
			t.Error("content hash mismatch")
		}

		return nil
	})
	if contentErr != nil {
		t.Error(contentErr)
	}

	for f := range scenario.Files {
		if _, ok := found[f]; !ok {
			l.Error("File missing", esl.String("file", f))
			t.Error("missing file")
		}
	}
}

func testSkip(t *testing.T, ctl app_control.Control, scenario Scenario, reportName, localBase string) {
	l := ctl.Log()
	found := make(map[string]bool)
	skipErr := qtr_endtoend.TestRows(ctl, reportName, func(cols map[string]string) error {
		r, err := filepath.Rel(localBase, cols["input.file"])
		if err != nil {
			l.Debug("unable to calc rel path", esl.Error(err))
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
			l.Error("File missing", esl.String("file", f))
			t.Error("missing file")
		}
	}
}

func execScenario(t *testing.T, ctl app_control.Control, short bool, f func(scenario Scenario, dbxBase mo_path.DropboxPath)) {
	l := ctl.Log()
	if qt_endtoend.IsSkipEndToEndTest() {
		l.Info("Skip end to end test")
		return
	}
	if _, err := dbx_conn_impl.ConnectTest(api_auth.DropboxTokenFull, app.PeerEndToEndTest, ctl); err != nil {
		l.Info("Skip: no end to end test resource found")
		return
	}

	if sc, err := NewScenario(short); err != nil {
		t.Error(err)
		return
	} else {
		dbxBase := qtr_endtoend.NewTestDropboxFolderPath(time.Now().Format("2006-01-02T15-04-05"))
		f(sc, dbxBase)
	}
}

func TestFileUpload(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, false, func(scenario Scenario, dbxBase mo_path.DropboxPath) {
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

func TestFileSyncUp(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, false, func(scenario Scenario, dbxBase mo_path.DropboxPath) {
			qtr_endtoend.ForkWithName(t, "file-sync-up", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
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
		execScenario(t, ctl, false, func(scenario Scenario, dbxBase mo_path.DropboxPath) {
			// `file sync up`
			qtr_endtoend.ForkWithName(t, "file-sync-up-dup1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
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
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
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

func TestFileCompareLocal(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, false, func(scenario Scenario, dbxBase mo_path.DropboxPath) {
			// `file sync up`
			qtr_endtoend.ForkWithName(t, "file-compare-local1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-compare-local")
				})
				if err != nil {
					t.Error(err)
				}

				testContent(t, fc, scenario, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-compare-local").Path())
				testSkip(t, fc, scenario, "skipped", scenario.LocalPath)
				return nil
			})

			// `file compare local`
			qtr_endtoend.ForkWithName(t, "file-compare-local2", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filecompare.Local{}, func(r rc_recipe.Recipe) {
					rc := r.(*filecompare.Local)
					rc.DropboxPath = dbxBase.ChildPath("file-compare-local")
					rc.LocalPath = mo_path2.NewFileSystemPath(scenario.LocalPath)
				})
				if err != nil {
					t.Error(err)
				}
				// TODO: verify result
				return err
			})
		})
	})
}

func TestFileSyncPreflightUp(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, true, func(scenario Scenario, dbxBase mo_path.DropboxPath) {
			qtr_endtoend.ForkWithName(t, "file-sync-preflight-up", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filesyncpreflight.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesyncpreflight.Up)
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

func TestFileFileList(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, true, func(scenario Scenario, dbxBase mo_path.DropboxPath) {
			// `file sync up`
			qtr_endtoend.ForkWithName(t, "file-list1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-list")
				})
				if err != nil {
					t.Error(err)
				}

				testContent(t, fc, scenario, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-list").Path())
				testSkip(t, fc, scenario, "skipped", scenario.LocalPath)
				return nil
			})

			// `file list`
			qtr_endtoend.ForkWithName(t, "file-list2", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &file.List{}, func(r rc_recipe.Recipe) {
					rc := r.(*file.List)
					rc.Path = dbxBase.ChildPath("file-list")
					rc.Recursive = true
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

func TestFileFileCopy(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, true, func(scenario Scenario, dbxBase mo_path.DropboxPath) {
			// `file sync up`
			qtr_endtoend.ForkWithName(t, "file-copy1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-copy")
				})
				if err != nil {
					t.Error(err)
				}

				testContent(t, fc, scenario, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-copy").Path())
				testSkip(t, fc, scenario, "skipped", scenario.LocalPath)
				return nil
			})

			// `file copy`
			qtr_endtoend.ForkWithName(t, "file-copy2", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &file.Copy{}, func(r rc_recipe.Recipe) {
					rc := r.(*file.Copy)
					rc.Src = dbxBase.ChildPath("file-copy")
					rc.Dst = dbxBase.ChildPath("file-copy2")
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

func TestFileFileMove(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, true, func(scenario Scenario, dbxBase mo_path.DropboxPath) {
			// `file sync up`
			qtr_endtoend.ForkWithName(t, "file-move1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
					ru.LocalPath = mo_path2.NewExistingFileSystemPath(scenario.LocalPath)
					ru.DropboxPath = dbxBase.ChildPath("file-move")
				})
				if err != nil {
					t.Error(err)
				}

				testContent(t, fc, scenario, "uploaded", scenario.LocalPath, dbxBase.ChildPath("file-move").Path())
				testSkip(t, fc, scenario, "skipped", scenario.LocalPath)
				return nil
			})

			// `file move`
			qtr_endtoend.ForkWithName(t, "file-move2", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &file.Move{}, func(r rc_recipe.Recipe) {
					rc := r.(*file.Move)
					rc.Src = dbxBase.ChildPath("file-move")
					rc.Dst = dbxBase.ChildPath("file-move2")
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

func TestFileFileMerge(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, false, func(scFrom Scenario, dbxBase mo_path.DropboxPath) {
			scTo, err := NewScenario(true)
			if err != nil {
				t.Error(err)
				return
			}

			// `file sync up`: scFrom
			qtr_endtoend.ForkWithName(t, "file-merge1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
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
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
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

func TestFileFileDelete(t *testing.T) {
	qtr_endtoend.TestWithControl(t, func(ctl app_control.Control) {
		execScenario(t, ctl, true, func(scenario Scenario, dbxBase mo_path.DropboxPath) {
			// `file sync up`
			qtr_endtoend.ForkWithName(t, "file-delete1", ctl, func(fc app_control.Control) error {
				err := rc_exec.Exec(fc, &filesync.Up{}, func(r rc_recipe.Recipe) {
					ru := r.(*filesync.Up)
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
