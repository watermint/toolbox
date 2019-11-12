package compare

import (
	"github.com/watermint/toolbox/domain/model/mo_file_diff"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_compare_paths"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/report/rp_spec_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type AccountVO struct {
	Left      app_conn.ConnUserFile
	Right     app_conn.ConnUserFile
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

func (z *Account) Requirement() app_vo.ValueObject {
	return &AccountVO{}
}

func (z *Account) Exec(k app_kitchen.Kitchen) error {
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

	ucc := uc_compare_paths.New(ctxLeft, ctxRight)
	count, err := ucc.Diff(mo_path.NewPath(vo.LeftPath), mo_path.NewPath(vo.RightPath), diff)
	if err != nil {
		return err
	}
	ui.Info("recipe.file.compare.account.success", app_msg.P{
		"DiffCount": count,
	})
	return nil
}

func (z *Account) Test(c app_control.Control) error {
	return qt_test.ImplementMe()
}
