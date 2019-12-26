package file

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
)

type MoveVO struct {
	Peer rc_conn.OldConnUserFile
	Src  string
	Dst  string
}

type Move struct {
}

func (z *Move) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Move) Console() {
}

func (z *Move) Requirement() rc_vo.ValueObject {
	return &MoveVO{}
}

func (z *Move) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*MoveVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	uc := uc_file_relocation.New(ctx)
	return uc.Move(mo_path.NewDropboxPath(vo.Src), mo_path.NewDropboxPath(vo.Dst))
}

func (z *Move) Test(c app_control.Control) error {
	return qt_endtoend.ScenarioTest()
}
