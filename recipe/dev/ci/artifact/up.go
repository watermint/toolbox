package artifact

import (
	"context"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context_impl"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/concurrency/es_timeout"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/file"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"github.com/watermint/toolbox/recipe/dev/ci/auth"
	"os"
	"time"
)

type Up struct {
	rc_recipe.RemarkSecret
	PeerName    string
	Timeout     int
	LocalPath   mo_path2.FileSystemPath
	DropboxPath mo_path.DropboxPath
	Upload      *file.Upload
}

func (z *Up) Preset() {
	z.PeerName = app.PeerDeploy
	z.Timeout = 30
}

func (z *Up) Exec(c app_control.Control) error {
	l := c.Log()

	if err := rc_exec.Exec(c, &auth.Import{}, func(r rc_recipe.Recipe) {
		m := r.(*auth.Import)
		m.PeerName = app.PeerDeploy
		m.EnvName = app.EnvNameDeployToken
	}); err != nil {
		l.Info("No token imported. Skip operation")
		return nil
	}

	a := api_auth_impl.NewConsoleCacheOnly(c, z.PeerName, dbx_auth.NewLegacyApp(c))
	ctx, err := a.Auth([]string{api_auth.DropboxTokenFull})
	if err != nil {
		l.Info("Skip operation")
		return nil
	}
	dbxCtx := dbx_context_impl.New(c, ctx)
	to := es_timeout.DoWithTimeout(time.Duration(z.Timeout)*time.Second, func(ctx context.Context) {
		err = rc_exec.Exec(c, &file.Upload{}, func(r rc_recipe.Recipe) {
			m := r.(*file.Upload)
			m.Context = dbxCtx
			m.LocalPath = z.LocalPath
			m.DropboxPath = z.DropboxPath
			m.Overwrite = true
		})
	})
	if to {
		l.Warn("Operation timeout")
		return nil
	}
	return err
}

func (z *Up) Test(c app_control.Control) error {
	if qt_endtoend.IsSkipEndToEndTest() {
		return qt_errors.ErrorSkipEndToEndTest
	}

	tp, err := qt_file.MakeTestFolder("up", true)
	if err != nil {
		return qt_errors.ErrorNotEnoughResource
	}
	defer func() {
		os.RemoveAll(tp)
	}()

	return rc_exec.Exec(c, z, func(r rc_recipe.Recipe) {
		m := r.(*Up)
		m.LocalPath = mo_path2.NewFileSystemPath(tp)
		m.DropboxPath = qtr_endtoend.NewTestDropboxFolderPath("dev-ci-artifact", time.Now().Format("2006-01-02-15-04-05"))
	})
}
