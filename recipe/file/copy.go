package file

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
)

type CopyVO struct {
	Peer app_conn.ConnUserFile
	Src  string
	Dst  string
}

type Copy struct {
}

func (z *Copy) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Copy) Console() {
}

func (z *Copy) Requirement() app_vo.ValueObject {
	return &CopyVO{}
}

func (z *Copy) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*CopyVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	uc := uc_file_relocation.New(ctx)
	return uc.Copy(mo_path.NewPath(vo.Src), mo_path.NewPath(vo.Dst))
}

func (z *Copy) Test(c app_control.Control) error {
	return qt_test.ImplementMe()
}
