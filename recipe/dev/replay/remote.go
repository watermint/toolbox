package replay

import (
	"context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/concurrency/es_timeout"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/io/es_zip"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/recipe/dev/ci/auth"
	"io/ioutil"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

type Remote struct {
	rc_recipe.RemarkSecret
	ReplayUrl   mo_string.OptionalString
	ResultsPath mo_path.DropboxPath
	PeerName    string
	Timeout     int
}

func (z *Remote) Preset() {
	z.Timeout = 60
	z.PeerName = app.PeerDeploy
	z.ResultsPath = mo_path.NewDropboxPath("/watermint-toolbox-logs/{{.Date}}-{{.Time}}/{{.Random}}")
}

func (z *Remote) Exec(c app_control.Control) error {
	url := os.Getenv(app.EnvNameReplayUrl)
	if z.ReplayUrl.IsExists() {
		url = z.ReplayUrl.Value()
	}
	l := c.Log()
	if url == "" {
		l.Warn("No replay url. Skip")
		return nil
	}

	url = regexp.MustCompile(`\?.*$`).ReplaceAllString(url, "") + "?raw=1"
	archivePath := filepath.Join(c.Workspace().Job(), "replay.zip")
	l.Debug("Downloading replay data", esl.String("url", url), esl.String("path", archivePath))
	err := es_download.Download(l, url, archivePath)
	if err != nil {
		l.Debug("Unable to download", esl.Error(err))
		return err
	}

	replayPath := filepath.Join(c.Workspace().Job(), "replay")
	l.Debug("Extract archive", esl.String("archivePath", archivePath), esl.String("replayPath", replayPath))
	err = es_zip.Extract(l, archivePath, replayPath)
	if err != nil {
		l.Debug("Unable to extract", esl.Error(err))
		return err
	}

	entries, err := ioutil.ReadDir(replayPath)
	if err != nil {
		l.Debug("Unable to read replay path", esl.Error(err))
		return err
	}
	for _, entry := range entries {
		if entry.IsDir() || !strings.HasSuffix(strings.ToLower(entry.Name()), ".zip") {
			continue
		}
		l.Info("Replay", esl.String("Entry", entry.Name()), esl.Int64("Size", entry.Size()))
	}

	l.Debug("Run replay bundle", esl.String("replayPath", replayPath))
	replayErr := rc_exec.Exec(c, &Bundle{}, func(r rc_recipe.Recipe) {
		m := r.(*Bundle)
		m.ReplayPath = mo_string.NewOptional(replayPath)
	})

	if replayErr == nil {
		return nil
	}

	l.Warn("One or more tests failed. Backup logs", esl.String("backupPath", z.ResultsPath.Path()))
	if err := rc_exec.Exec(c, &auth.Import{}, func(r rc_recipe.Recipe) {
		m := r.(*auth.Import)
		m.PeerName = z.PeerName
		m.EnvName = app.EnvNameDeployToken
	}); err != nil {
		l.Info("No token imported. Skip operation")
		return replayErr
	}
	a := api_auth_impl.NewConsoleCacheOnly(c, z.PeerName, dbx_auth.NewLegacyApp(c))
	ctx, err := a.Auth([]string{api_auth.DropboxTokenFull})
	if err != nil {
		l.Info("Skip operation")
		return replayErr
	}
	dbxCtx := dbx_context_impl.New(ctx.PeerName(), c, ctx)

	to := es_timeout.DoWithTimeout(time.Duration(z.Timeout)*time.Second, func(ctx context.Context) {
		err = rc_exec.Exec(c, &file.Upload{}, func(r rc_recipe.Recipe) {
			m := r.(*file.Upload)
			m.Context = dbxCtx
			m.LocalPath = mo_path2.NewFileSystemPath(c.Workspace().Job())
			m.DropboxPath = z.ResultsPath
			m.Overwrite = true
		})
	})
	if to {
		l.Warn("Operation timeout")
	}

	return replayErr
}

func (z *Remote) Test(c app_control.Control) error {
	return qt_errors.ErrorScenarioTest
}
