package unixtime

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"time"
)

const (
	PrecisionSecond      = "second"
	PrecisionMilliSecond = "ms"
	PrecisionNanoSecond  = "ns"
)

type Now struct {
	rc_recipe.RemarkTransient
	Precision mo_string.SelectString
}

func (z *Now) Preset() {
	z.Precision.SetOptions(
		PrecisionSecond,
		PrecisionSecond,
		PrecisionMilliSecond,
		PrecisionNanoSecond,
	)
}

func (z *Now) Exec(c app_control.Control) error {
	switch z.Precision.Value() {
	case PrecisionSecond:
		ui_out.TextOut(c, fmt.Sprintf("%d", time.Now().Unix()))
	case PrecisionMilliSecond:
		ui_out.TextOut(c, fmt.Sprintf("%d", time.Now().UnixNano()/1_000_000))
	case PrecisionNanoSecond:
		ui_out.TextOut(c, fmt.Sprintf("%d", time.Now().UnixNano()))
	default:
		c.Log().Error("Undefined precision", esl.String("precision", z.Precision.Value()))
		return errors.New("undefined precision")
	}
	return nil
}

func (z *Now) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Now{}, rc_recipe.NoCustomValues)
}
