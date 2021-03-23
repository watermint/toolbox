package paper

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_paper"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_paper"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"os"
)

type Create struct {
	Peer    dbx_conn.ConnScopedIndividual
	Content da_text.TextInput
	Path    mo_path.DropboxPath
	Format  mo_string.SelectString
	Created rp_model.RowReport
}

func (z *Create) Preset() {
	z.Format.SetOptions("markdown", "markdown", "plain_text", "html")
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
	z.Created.SetModel(
		&mo_paper.PaperUpdate{},
		rp_model.HiddenColumns(
			"file_id",
		),
	)
}

func (z *Create) Exec(c app_control.Control) error {
	if err := z.Created.Open(); err != nil {
		return err
	}
	content, err := z.Content.Content()
	if err != nil {
		return err
	}

	paper, err := sv_paper.New(z.Peer.Context()).Create(z.Path, z.Format.Value(), content)
	if err != nil {
		return err
	}
	z.Created.Row(&paper)
	return nil
}

func (z *Create) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("paper", "# Header\n* Dropbox\n* Paper")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Create{}, func(r rc_recipe.Recipe) {
		m := r.(*Create)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("paper", "test.paper")
		m.Content.SetFilePath(f)
	})
}
