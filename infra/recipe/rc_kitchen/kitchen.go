package rc_kitchen

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/recipe/rc_worker"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"go.uber.org/zap"
)

type Kitchen interface {
	// Deprecated:
	Value() rc_vo.ValueObject
	Control() app_control.Control
	UI() app_ui.UI
	Log() *zap.Logger
	NewQueue() rc_worker.Queue
}

type kitchenImpl struct {
	vo  rc_vo.ValueObject
	ctl app_control.Control
}

func (z *kitchenImpl) NewQueue() rc_worker.Queue {
	return z.ctl.NewQueue()
}

func (z *kitchenImpl) Value() rc_vo.ValueObject {
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

func NewKitchen(ctl app_control.Control, vo rc_vo.ValueObject) Kitchen {
	return &kitchenImpl{
		ctl: ctl,
		vo:  vo,
	}
}
