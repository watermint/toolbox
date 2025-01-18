package sharedlink

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Info struct {
	Peer       dbx_conn.ConnScopedIndividual
	Url        mo_url.Url
	Password   mo_string.OptionalString
	SharedLink rp_model.RowReport
	BasePath   mo_string.SelectString
}

func (z *Info) Preset() {
	z.Peer.SetScopes(dbx_auth.ScopeSharingRead)
	z.SharedLink.SetModel(
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"id",
			"path_display",
			"content_hash",
			"shared_folder_id",
			"parent_shared_folder_id",
		),
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Info) Exec(c app_control.Control) error {
	if err := z.SharedLink.Open(); err != nil {
		return err
	}

	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))

	link, err := sv_sharedlink.New(client).Resolve(z.Url, z.Password.Value())
	if err != nil {
		return err
	}
	z.SharedLink.Row(link.Concrete())
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Info{}, func(r rc_recipe.Recipe) {
		m := r.(*Info)
		url, _ := mo_url.NewUrl("https://www.dropbox.com/s/2sn712vy1ovegw8/Prime_Numbers.txt?dl=0")
		m.Url = url
	})
}
