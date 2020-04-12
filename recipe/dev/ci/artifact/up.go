package artifact

import (
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/infra/api/api_auth"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/file"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"github.com/watermint/toolbox/recipe/dev/ci/auth"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
	"path/filepath"
	"time"
)

type Up struct {
	PeerName    string
	LocalPath   mo_path2.FileSystemPath
	DropboxPath mo_path.DropboxPath
	Upload      *file.Upload
}

func (z *Up) Preset() {
	z.PeerName = qt_endtoend.DeployPeer
}

func (z *Up) Exec(c app_control.Control) error {
	l := c.Log()

	if err := rc_exec.Exec(c, &auth.Import{}, func(r rc_recipe.Recipe) {
		m := r.(*auth.Import)
		m.PeerName = qt_endtoend.DeployPeer
		m.EnvName = qt_endtoend.DeployEnvToken
	}); err != nil {
		l.Info("No token imported. Skip operation")
		return nil
	}

	a := api_auth_impl.NewConsoleCacheOnly(c, z.PeerName)
	ctx, err := a.Auth(api_auth.DropboxTokenFull)
	if err != nil {
		l.Info("Skip operation")
		return nil
	}
	dbxCtx, ok := ctx.(dbx_context.Context)
	if !ok {
		l.Error("Incompatible context type found", zap.Any("ctx", ctx))
		return api_context.ErrorIncompatibleContextType
	}
	err = rc_exec.Exec(c, &file.Upload{}, func(r rc_recipe.Recipe) {
		m := r.(*file.Upload)
		m.Context = dbxCtx
		m.LocalPath = z.LocalPath
		m.DropboxPath = z.DropboxPath
		m.Overwrite = true
	})
	if err != nil {
		return err
	}

	return nil
}

func (z *Up) Test(c app_control.Control) error {
	tp, err := ioutil.TempDir("", "up")
	if err != nil {
		return err
	}
	if err := ioutil.WriteFile(filepath.Join(tp, "artifact.txt"), []byte(time.Now().String()), 0644); err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(tp)
	}()

	return rc_exec.Exec(c, z, func(r rc_recipe.Recipe) {
		m := r.(*Up)
		m.LocalPath = mo_path2.NewFileSystemPath(tp)
		m.DropboxPath = qt_recipe.NewTestDropboxFolderPath("dev-ci-artifact", time.Now().Format(time.RFC3339))
	})
}
