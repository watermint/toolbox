package compare

import (
	"github.com/watermint/toolbox/domain/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_compare_paths"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type AccountVO struct {
	Left      rc_conn.ConnUserFile
	Right     rc_conn.ConnUserFile
	LeftPath  string
	RightPath string
}

const (
	reportAccount = "diff"
)

type Account struct {
}

func (z *Account) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{
		rp_spec_impl.Spec(reportAccount, &mo_file_diff.Diff{}),
	}
}

func (z *Account) Console() {
}

func (z *Account) Requirement() rc_vo.ValueObject {
	return &AccountVO{}
}

func (z *Account) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*AccountVO)
	ui := k.UI()

	ui.Info("recipe.file.compare.account.conn_left")
	ctxLeft, err := vo.Left.Connect(k.Control())
	if err != nil {
		return err
	}

	ui.Info("recipe.file.compare.account.conn_right")
	ctxRight, err := vo.Right.Connect(k.Control())
	if err != nil {
		return err
	}

	rep, err := rp_spec_impl.New(z, k.Control()).Open(reportAccount)
	if err != nil {
		return err
	}
	defer rep.Close()

	diff := func(diff mo_file_diff.Diff) error {
		rep.Row(&diff)
		return nil
	}

	ucc := uc_compare_paths.New(ctxLeft, ctxRight, k.UI())
	count, err := ucc.Diff(mo_path.NewDropboxPath(vo.LeftPath), mo_path.NewDropboxPath(vo.RightPath), diff)
	if err != nil {
		return err
	}
	ui.Info("recipe.file.compare.account.success", app_msg.P{
		"DiffCount": count,
	})
	return nil
}

func (z *Account) Test(c app_control.Control) error {
	return qt_recipe.ImplementMe()
}
