package dev

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/quality/qt_runtime"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type Quality struct {
}

func (z *Quality) Test(c app_control.Control) error {
	return nil
}

func (z *Quality) Hidden() {
}

func (z *Quality) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (z *Quality) Exec(k app_kitchen.Kitchen) error {
	qt_runtime.Suite(k.Control())
	return nil
}
