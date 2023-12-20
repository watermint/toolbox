package replay

import (
	"context"
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/essentials/api/api_auth_oauth"
	"github.com/watermint/toolbox/essentials/concurrency/es_timeout"
	"github.com/watermint/toolbox/essentials/io/es_zip"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_replay"
	"github.com/watermint/toolbox/ingredient/ig_file"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"io/ioutil"
	"path/filepath"
	"strings"
	"time"
)

type Bundle struct {
	rc_recipe.RemarkSecret
	ReplayPath  mo_string.OptionalString
	ResultsPath mo_path.DropboxPath
	PeerName    string
	Timeout     int
}

func (z *Bundle) Preset() {
	z.Timeout = 60
	z.PeerName = app_definitions.PeerDeploy
	z.ResultsPath = mo_path.NewDropboxPath("/watermint-toolbox-logs/{{.Date}}-{{.Time}}/{{.Random}}")
}

func (z *Bundle) deployDbxContext(c app_control.Control) (client dbx_client.Client, err error) {
	l := c.Log()
	sd := api_auth.OAuthSessionData{
		AppData:  dbx_auth.DropboxIndividual,
		PeerName: z.PeerName,
		Scopes: []string{
			dbx_auth.ScopeFilesContentRead,
			dbx_auth.ScopeFilesContentWrite,
		},
	}
	session := api_auth_oauth.NewSessionDeployEnv(app_definitions.EnvNameDeployToken)
	entity, err := session.Start(sd)
	if err != nil {
		l.Info("No token found. Skip operation")
		return nil, errors.New("skip")
	}
	client = dbx_client_impl.New(c, dbx_auth.DropboxIndividual, entity)
	return
}

func (z *Bundle) execReplay(l esl.Logger, entryName string, replay rc_replay.Replay, dbxCtx dbx_client.Client, c, forkCtl app_control.Control) (err error) {
	defer func() {
		if rescue := recover(); rescue != nil {
			var ok bool
			if err, ok = rescue.(error); ok {
				l.Warn("Warn: panic", esl.Error(err), esl.String("entry", entryName))
			} else {
				l.Warn("Warn: panic", esl.Any("error", rescue), esl.String("entry", entryName))
				err = errors.New("panic")
			}
		}
	}()

	start := time.Now()
	l.Debug("Running", esl.String("entryName", entryName))
	err = replay.Replay(forkCtl.Workspace(), forkCtl)
	duration := time.Now().Sub(start).Truncate(time.Millisecond)

	if err != nil {
		l.Warn("Error on replay", esl.Error(err))
		l.Info("Uploading logs")
		to := es_timeout.DoWithTimeout(time.Duration(z.Timeout)*time.Second, func(ctx context.Context) {
			err = rc_exec.Exec(c, &ig_file.Upload{}, func(r rc_recipe.Recipe) {
				m := r.(*ig_file.Upload)
				m.Context = dbxCtx
				m.LocalPath = mo_path2.NewFileSystemPath(forkCtl.Workspace().Job())
				m.DropboxPath = z.ResultsPath
				m.Overwrite = true
			})
		})
		if to {
			l.Warn("Operation timeout")
		}
		return err
	}

	l.Debug("Success", esl.Duration("duration", duration))
	return nil
}

func (z *Bundle) Exec(c app_control.Control) error {
	l := c.Log()
	replayPath, err := rc_replay.ReplayPath(z.ReplayPath)
	if err != nil {
		l.Warn("Unable to find replay path, skip run replay bundle", esl.Error(err))
		return nil
	}

	entries, err := ioutil.ReadDir(replayPath)
	if err != nil {
		return err
	}

	ea_indicator.SuppressIndicatorForce()

	dbxCtx, err := z.deployDbxContext(c)
	if err != nil {
		l.Warn("No deploy token found. Skip uploading logs on failure")
	}

	var recipeErr error

	for _, entry := range entries {
		entryLower := strings.ToLower(entry.Name())
		l := c.Log().With(esl.String("entry", entryLower))
		replay := rc_replay.New(c.Log())
		if entry.IsDir() || !strings.HasSuffix(entryLower, ".zip") {
			l.Debug("Skip entry", esl.String("entry", entry.Name()))
			continue
		}

		entryName := strings.ReplaceAll(entryLower, ".zip", "")
		if entryName == "" {
			l.Debug("Skip")
			continue
		}

		forkCtl, err := app_control_impl.ForkQuiet(c, entryName)
		if err != nil {
			l.Debug("Unable to fork bundle", esl.Error(err))
			return err
		}

		err = es_zip.Extract(l, filepath.Join(replayPath, entry.Name()), forkCtl.Workspace().Job())
		if err != nil {
			l.Debug("Unable to extract", esl.Error(err))
			return err
		}
		recipeErr = z.execReplay(l, entryName, replay, dbxCtx, c, forkCtl)
	}
	return recipeErr
}

func (z *Bundle) Test(c app_control.Control) error {
	return qt_errors.ErrorScenarioTest
}
