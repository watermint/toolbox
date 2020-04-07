package compare

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_compare_paths"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type Account struct {
	Left      dbx_conn.ConnUserFile
	Right     dbx_conn.ConnUserFile
	LeftPath  mo_path.DropboxPath
	RightPath mo_path.DropboxPath
	Diff      rp_model.RowReport
	ConnLeft  app_msg.Message
	ConnRight app_msg.Message
	Success   app_msg.Message
}

func (z *Account) Preset() {
	z.Diff.SetModel(&mo_file_diff.Diff{})
	z.Left.SetPeerName("left")
	z.Right.SetPeerName("right")
}

func (z *Account) Exec(c app_control.Control) error {
	ui := c.UI()

	ui.Info(z.ConnLeft)
	ctxLeft := z.Left.Context()

	ui.Info(z.ConnRight)
	ctxRight := z.Right.Context()

	err := z.Diff.Open()
	if err != nil {
		return err
	}

	diff := func(diff mo_file_diff.Diff) error {
		z.Diff.Row(&diff)
		return nil
	}

	ucc := uc_compare_paths.New(ctxLeft, ctxRight, c.UI())
	count, err := ucc.Diff(z.LeftPath, z.RightPath, diff)
	if err != nil {
		return err
	}
	ui.Info(z.Success.With("DiffCount", count))
	return nil
}

func (z *Account) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Account{}, func(r rc_recipe.Recipe) {
		m := r.(*Account)
		m.LeftPath = qt_recipe.NewTestDropboxFolderPath("left")
		m.RightPath = qt_recipe.NewTestDropboxFolderPath("right")
	})
}
