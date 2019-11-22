package file

import (
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_file"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type DeleteVO struct {
	Peer app_conn.ConnUserFile
	Path string
}

type Delete struct {
}

func (z *Delete) Console() {
}

func (z *Delete) Requirement() app_vo.ValueObject {
	return &DeleteVO{}
}

func (z *Delete) Exec(k app_kitchen.Kitchen) error {
	vo := k.Value().(*DeleteVO)
	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	_, err = sv_file.NewFiles(ctx).Remove(mo_path.NewPath(vo.Path))
	if err != nil {
		return err
	}
	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return qt_recipe.ScenarioTest()
}

func (z *Delete) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}
