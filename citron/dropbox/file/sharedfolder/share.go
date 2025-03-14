package sharedfolder

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedfolder"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type Share struct {
	Peer             dbx_conn.ConnScopedIndividual
	Path             mo_path.DropboxPath
	AclUpdatePolicy  mo_string.SelectString
	MemberPolicy     mo_string.SelectString
	SharedLinkPolicy mo_string.SelectString
	Shared           rp_model.RowReport
	BasePath         mo_string.SelectString
}

func (z *Share) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeSharingWrite,
	)
	z.AclUpdatePolicy.SetOptions(
		"owner",
		"owner", "editor",
	)
	z.MemberPolicy.SetOptions(
		"anyone",
		"team", "anyone",
	)
	z.SharedLinkPolicy.SetOptions(
		"anyone",
		"anyone", "members",
	)
	z.Shared.SetModel(&mo_sharedfolder.SharedFolder{})
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *Share) Exec(c app_control.Control) error {
	if err := z.Shared.Open(); err != nil {
		return err
	}
	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	sf, err := sv_sharedfolder.New(client).Create(
		z.Path,
		sv_sharedfolder.AclUpdatePolicy(z.AclUpdatePolicy.Value()),
		sv_sharedfolder.MemberPolicy(z.MemberPolicy.Value()),
		sv_sharedfolder.SharedLinkPolicy(z.SharedLinkPolicy.Value()),
	)
	if err != nil {
		return err
	}
	z.Shared.Row(sf)
	return nil
}

func (z *Share) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Share{}, func(r rc_recipe.Recipe) {
		m := r.(*Share)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("test")
	})
}
