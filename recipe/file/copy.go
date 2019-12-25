package file

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type CopyVO struct {
	Peer rc_conn.OldConnUserFile
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

func (z *Copy) Requirement() rc_vo.ValueObject {
	return &CopyVO{}
}

func (z *Copy) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*CopyVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	uc := uc_file_relocation.New(ctx)
	return uc.Copy(mo_path.NewDropboxPath(vo.Src), mo_path.NewDropboxPath(vo.Dst))
}

func (z *Copy) Test(c app_control.Control) error {
	return qt_recipe.ScenarioTest()
}
