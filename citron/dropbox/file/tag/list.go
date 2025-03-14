package tag

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_tag"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type FileTag struct {
	Path string `json:"path"`
	Tag  string `json:"tag"`
}

type List struct {
	Peer     dbx_conn.ConnScopedIndividual
	Path     mo_path.DropboxPath
	Tags     rp_model.RowReport
	BasePath mo_string.SelectString
}

func (z *List) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesMetadataRead,
	)
	z.Tags.SetModel(&FileTag{})
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Tags.Open(); err != nil {
		return err
	}
	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))

	tags, err := sv_file_tag.New(client).Resolve(z.Path)
	if err != nil {
		return err
	}

	for _, tag := range tags {
		z.Tags.Row(&FileTag{
			Path: z.Path.Path(),
			Tag:  tag,
		})
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("list")
	})
}
