package test

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_runtime"
)

type Resources struct {
}

func (z *Resources) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Resources) Test(c app_control.Control) error {
	return qt_recipe.NoTestRequired()
}

func (z *Resources) Hidden() {
}

func (z *Resources) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (z *Resources) Exec(k app_kitchen.Kitchen) error {
	qt_runtime.Suite(k.Control())
	return nil
}
