package qs_file

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn_impl"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/demo/qdm_file"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"path/filepath"
	"testing"
	"time"
)

func testContent(t *testing.T, ctl app_control.Control, scenario qdm_file.Scenario, reportName, localBase, dbxBase string) {
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

func testSkip(t *testing.T, ctl app_control.Control, scenario qdm_file.Scenario, reportName, localBase string) {
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

func execScenario(t *testing.T, ctl app_control.Control, short bool, f func(scenario qdm_file.Scenario, dbxBase mo_path.DropboxPath)) {
	l := ctl.Log()
	if qt_endtoend.IsSkipEndToEndTest() {
		l.Info("Skip end to end test")
		return
	}
	if _, err := dbx_conn_impl.ConnectTest(api_auth.DropboxTokenFull, app.PeerEndToEndTest, ctl); err != nil {
		l.Info("Skip: no end to end test resource found")
		return
	}

	if sc, err := qdm_file.NewScenario(short); err != nil {
		t.Error(err)
		return
	} else {
		dbxBase := qtr_endtoend.NewTestDropboxFolderPath(time.Now().Format("2006-01-02T15-04-05"))
		f(sc, dbxBase)
	}
}
