package share

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharing"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Info struct {
	Peer     dbx_conn.ConnScopedIndividual
	Path     mo_path.DropboxPath
	Metadata rp_model.RowReport
	BasePath mo_string.SelectString
}

func (z *Info) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentRead,
		dbx_auth.ScopeSharingRead,
	)
	z.Metadata.SetModel(&mo_file.ConcreteEntry{})
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Info) Exec(c app_control.Control) error {
	if err := z.Metadata.Open(); err != nil {
		return err
	}

	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	info, err := sv_sharing.New(client).Resolve(z.Path.Path())
	if err != nil {
		return err
	}

	z.Metadata.Row(info)
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Info{}, func(r rc_recipe.Recipe) {
		m := r.(*Info)
		m.Path = mo_path.NewDropboxPath("id:3kmLmQFnf1AAAAAAAAAAAw")
	})
}
