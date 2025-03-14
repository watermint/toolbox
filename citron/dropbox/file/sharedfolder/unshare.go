package sharedfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_sharedfolder"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Unshare struct {
	Peer      dbx_conn.ConnScopedIndividual
	Path      mo_path.DropboxPath
	LeaveCopy bool
	BasePath  mo_string.SelectString
}

func (z *Unshare) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeSharingWrite,
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Unshare) Exec(c app_control.Control) error {
	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	sf, err := uc_sharedfolder.NewResolver(client).Resolve(z.Path)
	if err != nil {
		return err
	}

	return sv_sharedfolder.New(client).Remove(sf, sv_sharedfolder.LeaveACopy(z.LeaveCopy))
}

func (z *Unshare) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Unshare{}, func(r rc_recipe.Recipe) {
		m := r.(*Unshare)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("test")
	})
}
