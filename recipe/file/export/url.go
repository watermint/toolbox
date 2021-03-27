package export

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	mo_path2 "github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_url"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"os"
	"path/filepath"
)

type Url struct {
	Peer                  dbx_conn.ConnScopedIndividual
	Url                   mo_url.Url
	Password              mo_string.OptionalString
	LocalPath             mo_path.FileSystemPath
	OperationLog          rp_model.RowReport
	Format                mo_string.OptionalString
	ErrorNotInYourDropbox app_msg.Message
}

func (z *Url) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeSharingRead,
		dbx_auth.ScopeFilesContentRead,
	)
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
}

func (z *Url) Exec(c app_control.Control) error {
	l := c.Log()
	if err := z.OperationLog.Open(); err != nil {
		return err
	}

	link, err := sv_sharedlink.New(z.Peer.Context()).Resolve(z.Url, z.Password.Value())
	if err != nil {
		l.Debug("Unable to retrieve shared link")
		return err
	}

	if link.PathLower() == "" {
		c.UI().Error(z.ErrorNotInYourDropbox)
		return errors.New("the document is not in your dropbox")
	}

	export, path, err := sv_file_content.NewExport(z.Peer.Context()).Export(mo_path2.NewDropboxPath(link.PathLower()),
		sv_file_content.ExportFormat(z.Format.Value()))
	if err != nil {
		return err
	}
	dest := filepath.Join(z.LocalPath.Path(), export.ExportName)
	if err := os.Rename(path.Path(), dest); err != nil {
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

func (z *Url) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Url{}, func(r rc_recipe.Recipe) {
		m := r.(*Url)
		m.LocalPath = qtr_endtoend.NewTestFileSystemFolderPath(c, "export-doc")
		url, _ := mo_url.NewUrl("https://www.dropbox.com/s/2sn712vy1ovegw8/Prime_Numbers.txt?dl=0")
		m.Url = url
	})
}
