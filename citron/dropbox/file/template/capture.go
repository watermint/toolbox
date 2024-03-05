package template

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/filesystem/dbx_fs"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_tag"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
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

type Capture struct {
	Peer dbx_conn.ConnScopedIndividual
	Path mo_path.DropboxPath
	Out  mo_path2.FileSystemPath
}

func (z *Capture) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeFilesMetadataRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
	)
}

func (z *Capture) findSourceLink(path es_filesystem.Path) (link string, err error) {
	svl := sv_sharedlink.New(z.Peer.Client())
	links, err := svl.ListByPath(mo_path.NewDropboxPath(path.Path()))
	if err != nil {
		return "", err
	}
	if len(links) < 1 {
		return "", nil
	}
	return links[0].LinkUrl(), nil
}

func (z *Capture) createSourceLink(path es_filesystem.Path) (link string, err error) {
	svl := sv_sharedlink.New(z.Peer.Client())
	sl, err := svl.Create(mo_path.NewDropboxPath(path.Path()))
	if err != nil {
		return "", err
	}
	return sl.LinkUrl(), nil
}

func (z *Capture) handlerSource(path es_filesystem.Path) (link string, err error) {
	link, err = z.findSourceLink(path)
	if err != nil {
		return "", err
	}
	if link == "" {
		link, err = z.createSourceLink(path)
		if err != nil {
			return "", err
		}
	}

	dlLink, err := sv_sharedlink.ToDownloadUrl(link)
	if err != nil {
		return "", err
	}
	return dlLink, nil
}

func (z *Capture) handlerTags(path es_filesystem.Path) (tags []string, err error) {
	return sv_file_tag.New(z.Peer.Client()).Resolve(mo_path.NewDropboxPath(path.Path()))
}

func (z *Capture) Exec(c app_control.Control) error {
	dfs := dbx_fs.NewFileSystem(z.Peer.Client())
	cp := es_template.NewCapture(dfs, es_template.CaptureOpts{
		HandlerSource: z.handlerSource,
		HandlerTags:   z.handlerTags,
	})

	template, err := cp.Capture(dbx_fs.NewPath("", z.Path))
	if err != nil {
		return err
	}
	tj, err := json.MarshalIndent(template, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(z.Out.Path(), tj, 0644)
}

func (z *Capture) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("capture", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &Capture{}, func(r rc_recipe.Recipe) {
		m := r.(*Capture)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("capture")
		m.Out = mo_path2.NewFileSystemPath(filepath.Join(f, "test.json"))
	})
}
