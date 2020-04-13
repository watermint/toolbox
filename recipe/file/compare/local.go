package compare

import (
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/usecase/uc_compare_local"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type Local struct {
	Peer        dbx_conn.ConnUserFile
	LocalPath   mo_path2.FileSystemPath
	DropboxPath mo_path.DropboxPath
	Diff        rp_model.RowReport
	Skip        rp_model.RowReport
	Success     app_msg.Message
}

func (z *Local) Preset() {
	z.Diff.SetModel(&mo_file_diff.Diff{})
	z.Skip.SetModel(&mo_file_diff.Diff{})
}

func (z *Local) Exec(c app_control.Control) error {
	ui := c.UI()
	ctx := z.Peer.Context()

	if err := z.Diff.Open(rp_model.NoConsoleOutput()); err != nil {
		return err
	}
	if err := z.Skip.Open(rp_model.NoConsoleOutput()); err != nil {
		return err
	}

	diff := func(diff mo_file_diff.Diff) error {
		app_ui.ShowProgress(c.UI())
		switch diff.DiffType {
		case mo_file_diff.DiffSkipped:
			z.Skip.Row(&diff)
		default:
			z.Diff.Row(&diff)
		}
		return nil
	}

	ucl := uc_compare_local.New(ctx, c.UI())
	count, err := ucl.Diff(z.LocalPath, z.DropboxPath, diff)
	if err != nil {
		return err
	}
	ui.Info(z.Success.With("DiffCount", count))
	return nil
}

func (z *Local) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Local{}, func(r rc_recipe.Recipe) {
		m := r.(*Local)
		m.LocalPath = qt_recipe.NewTestFileSystemFolderPath(c, "compare")
		m.DropboxPath = qt_recipe.NewTestDropboxFolderPath("compare")
	})
	if err, _ = qt_recipe.RecipeError(c.Log(), err); err != nil {
		return err
	}
	return qt_errors.ErrorScenarioTest
}
