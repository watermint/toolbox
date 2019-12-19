package rc_spec

import (
	"errors"
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"testing"
)

type SelfContainedTestRecipe struct {
	ProgressStart app_msg.Message
	Start         mo_time.Time
	Enabled       bool
	Limit         int
	Name          string
}

func (z *SelfContainedTestRecipe) Exec(k rc_kitchen.Kitchen) error {
	ui := k.UI()
	ui.InfoM(z.ProgressStart)

	if !z.Enabled {
		return errors.New("!enabled")
	}
	if z.Limit != 20 {
		return errors.New("limit != 20")
	}
	if z.Name != "hey" {
		return errors.New("!= hey")
	}
	if z.Start.Iso8601() != "2010-11-12T13:14:15Z" {
		return errors.New("!= 2010-11-12T13:14:15Z")
	}
	return nil
}

func (z *SelfContainedTestRecipe) Test(c app_control.Control) error {
	return qt_recipe.NoTestRequired()
}

func (z *SelfContainedTestRecipe) Init() {
	z.Limit = 10
}

func TestSpecSelfContained_ApplyValues(t *testing.T) {
	qt_recipe.TestWithControl(t, func(ctl app_control.Control) {
		scr := &SelfContainedTestRecipe{}
		spec := newSelfContained(scr)

		f := flag.NewFlagSet("test", flag.ContinueOnError)
		spec.SetFlags(f, ctl.UI())
		err := f.Parse([]string{"-enabled",
			"-limit", "20",
			"-name", "hey",
			"-start", "2010-11-12T13:14:15Z",
		})
		if err != nil {
			t.Error(err)
			return
		}

		{
			rcp, k := spec.ApplyValues(ctl)
			if err = rcp.Exec(k); err != nil {
				t.Error(err)
			}
			if err = rcp.Test(ctl); err != nil {
				switch e := err.(type) {
				case *qt_recipe.ErrorNoTestRequired:
					ctl.Log().Debug("ok")
				default:
					t.Error(e)
				}
			}
		}
	})
}
