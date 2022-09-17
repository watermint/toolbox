package capture

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_template"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"os"
	"path/filepath"
)

type Remote struct {
	Peer dbx_conn.ConnScopedIndividual
	Path mo_path.DropboxPath
	Out  mo_path2.FileSystemPath
}

func (z *Remote) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
	)
}

func (z *Remote) handlerSource(path es_filesystem.Path) (link string, err error) {
	return "", nil
}

func (z *Remote) handlerTags(path es_filesystem.Path) (tags []string, err error) {
	return []string{}, nil
}

func (z *Remote) Exec(c app_control.Control) error {
	dfs := filesystem.NewFileSystem(z.Peer.Client())
	cp := es_template.NewCapture(dfs, es_template.CaptureOpts{
		HandlerSource: z.handlerSource,
		HandlerTags:   z.handlerTags,
	})

	template, err := cp.Capture(filesystem.NewPath("", z.Path))
	if err != nil {
		return err
	}
	tj, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(z.Out.Path(), tj, 0644)
}

func (z *Remote) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("capture", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &Remote{}, func(r rc_recipe.Recipe) {
		m := r.(*Remote)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("capture")
		m.Out = mo_path2.NewFileSystemPath(filepath.Join(f, "test.json"))
	})
}
