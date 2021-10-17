package sharedfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Unshare struct {
	Peer      dbx_conn.ConnScopedIndividual
	Path      mo_path.DropboxPath
	LeaveCopy bool
}

func (z *Unshare) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
	)
}

func (z *Unshare) Exec(c app_control.Control) error {
	f1, err := sv_file.NewFiles(z.Peer.Context()).Resolve(z.Path)
	if err != nil {
		return err
	}
	f2, ok := f1.Folder()
	if !ok {
		return errors.New("shared folder not found")
	}
	if f2.EntrySharedFolderId == "" {
		return errors.New("the folder is not a shared folder")
	}
	sf, err := sv_sharedfolder.New(z.Peer.Context()).Resolve(f2.EntrySharedFolderId)
	if err != nil {
		return err
	}

	return sv_sharedfolder.New(z.Peer.Context()).Remove(sf, sv_sharedfolder.LeaveACopy(z.LeaveCopy))
}

func (z *Unshare) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Unshare{}, func(r rc_recipe.Recipe) {
		m := r.(*Unshare)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("test")
	})
}
