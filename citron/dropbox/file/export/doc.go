package export

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filemove"
	"github.com/watermint/toolbox/essentials/log/esl"
	mo_path2 "github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"path/filepath"
)

type Doc struct {
	rc_recipe.RemarkExperimental
	Peer         dbx_conn.ConnScopedIndividual
	LocalPath    mo_path2.FileSystemPath
	DropboxPath  mo_path.DropboxPath
	OperationLog rp_model.RowReport
	Format       mo_string.OptionalString
	BasePath     mo_string.SelectString
}

func (z *Doc) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	client := z.Peer.Client().BaseNamespace(dbx_filesystem.AsNamespaceType(z.BasePath.Value()))
	export, path, err := sv_file_content.NewExport(client).Export(z.DropboxPath, sv_file_content.ExportFormat(z.Format.Value()))
	if err != nil {
		return err
	}
	dest := filepath.Join(z.LocalPath.Path(), export.ExportName)
	if err := es_filemove.Move(path.Path(), dest); err != nil {
		l.Debug("Unable to move file to specified path",
			esl.Error(err),
			esl.String("downloaded", path.Path()),
			esl.String("destination", dest),
		)
		return err
	}

	z.OperationLog.Row(export)

	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Doc{}, func(r rc_recipe.Recipe) {
		m := r.(*Doc)
		m.LocalPath = qtr_endtoend.NewTestFileSystemFolderPath(c, "export-doc")
		m.DropboxPath = qtr_endtoend.NewTestDropboxFolderPath("file-export-doc")
	})
}

func (z *Doc) Preset() {
	z.Peer.SetScopes(dbx_auth.ScopeFilesContentRead)
	z.OperationLog.SetModel(
		&mo_file.Export{},
		rp_model.HiddenColumns(
			"path_lower",
			"id",
			"revision",
			"content_hash",
			"export_hash",
		),
	)
	z.BasePath.SetOptions(
		dbx_filesystem.BaseNamespaceDefaultInString,
		dbx_filesystem.BaseNamespaceTypesInString...,
	)
}
