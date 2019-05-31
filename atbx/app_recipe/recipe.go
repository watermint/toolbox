package app_recipe

import (
	"github.com/watermint/toolbox/atbx/app_control"
	"github.com/watermint/toolbox/atbx/app_report"
	"github.com/watermint/toolbox/atbx/app_ui"
	"github.com/watermint/toolbox/atbx/app_vo"
	"go.uber.org/zap"
)

type Recipe interface {
	Requirement() app_vo.ValueObject
	Exec(k Kitchen) error
}

// SecretRecipe will not be listed in available commands.
type SecretRecipe interface {
	Hidden()
}

type Kitchen interface {
	Value() app_vo.ValueObject
	Control() app_control.Control
	UI() app_ui.UI
	Log() *zap.Logger
	Report(name string, row interface{}) (r app_report.Report, err error)
}

type kitchenImpl struct {
	vo  app_vo.ValueObject
	ctl app_control.Control
}

func (z *kitchenImpl) Value() app_vo.ValueObject {
	return z.vo
}

func (z *kitchenImpl) Control() app_control.Control {
	return z.ctl
}

func (z *kitchenImpl) UI() app_ui.UI {
	return z.ctl.UI()
}

func (z *kitchenImpl) Log() *zap.Logger {
	return z.ctl.Log()
}

func (z *kitchenImpl) Report(name string, row interface{}) (r app_report.Report, err error) {
	return app_report.New(name, row, z.ctl)
}

func NewKitchen(ctl app_control.Control, vo app_vo.ValueObject) Kitchen {
	return &kitchenImpl{
		ctl: ctl,
		vo:  vo,
	}
}
