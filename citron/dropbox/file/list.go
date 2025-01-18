package file

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
)

type List struct {
	Peer                         dbx_conn.ConnScopedIndividual
	Path                         mo_path.DropboxPath
	Recursive                    bool
	IncludeDeleted               bool
	IncludeMountedFolders        bool
	IncludeExplicitSharedMembers bool
	FileList                     rp_model.RowReport
	BasePath                     mo_string.SelectString
}

func (z *List) Preset() {
	z.Peer.SetScopes(dbx_auth.ScopeFilesContentRead)
	z.FileList.SetModel(
		&mo_file.ConcreteEntry{},
		rp_model.HiddenColumns(
			"id",
			"path_lower",
			"revision",
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

func (z *List) Exec(c app_control.Control) error {
	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))

	opts := make([]sv_file.ListOpt, 0)
	opts = append(opts, sv_file.IncludeDeleted(z.IncludeDeleted))
	opts = append(opts, sv_file.Recursive(z.Recursive))
	opts = append(opts, sv_file.IncludeHasExplicitSharedMembers(z.IncludeExplicitSharedMembers))
	opts = append(opts, sv_file.IncludeMountedFolders(z.IncludeMountedFolders))

	if err := z.FileList.Open(); err != nil {
		return err
	}

	err := sv_file.NewFiles(client).ListEach(z.Path, func(entry mo_file.Entry) {
		z.FileList.Row(entry.Concrete())
	}, opts...)
	if err != nil {
		c.Log().Debug("Failed to list files")
		return err
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	err := rc_exec.Exec(c, &List{}, func(r rc_recipe.Recipe) {
		r0 := r.(*List)
		r0.Path = qtr_endtoend.NewTestDropboxFolderPath()
		r0.Recursive = false
	})
	if err != nil {
		return err
	}
	return qtr_endtoend.TestRows(c, "file_list", func(cols map[string]string) error {
		if _, ok := cols["name"]; !ok {
			return errors.New("`name` is not found")
		}
		if _, ok := cols["path_display"]; !ok {
			return errors.New("`path_display` is not found")
		}
		return nil
	})
}
