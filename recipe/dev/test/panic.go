package test

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Panic struct {
	rc_recipe.RemarkSecret
	PanicType int
}

func (z *Panic) Preset() {
}

func (z *Panic) Exec(c app_control.Control) error {
	switch z.PanicType {
	case 0:
		c.Log().Info("Result", esl.Int("value", 10/z.PanicType))

	case 1:
		var dic map[string]string
		dic["hello"] = "world"
		c.Log().Info("Result", esl.Any("dic", dic))

	default:
		panic("other panic")
	}
	return nil
}

func (z *Panic) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
