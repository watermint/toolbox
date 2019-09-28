package test

import (
	"errors"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"go.uber.org/zap"
	"strings"
)

type RecipeVO struct {
	All    bool
	Recipe string
}

type Recipe struct {
}

func (z *Recipe) Console() {
}

func (z *Recipe) Hidden() {
}

func (z *Recipe) Requirement() app_vo.ValueObject {
	return &RecipeVO{}
}

func (z *Recipe) Exec(k app_kitchen.Kitchen) error {
	cl := k.Control().(app_control_launcher.ControlLauncher)
	cat := cl.Catalogue()
	l := k.Log()
	vo := k.Value().(*RecipeVO)

	switch {
	case vo.All:
		for _, r := range cat {
			path, name := app_recipe.Path(r)
			ll := l.With(zap.Strings("path", path), zap.String("name", name))
			ll.Info("Testing: ")

			if err := r.Test(k.Control()); err != nil {
				ll.Error("Error", zap.Error(err))
				return err
			}
		}

	case vo.Recipe != "":
		for _, r := range cat {
			path, name := app_recipe.Path(r)
			p := strings.Join(append(path, name), ".")
			if p != vo.Recipe {

			}
			ll := l.With(zap.Strings("path", path), zap.String("name", name))
			ll.Info("Testing: ")

			if err := r.Test(k.Control()); err != nil {
				ll.Error("Error", zap.Error(err))
				return err
			} else {
				return nil
			}
		}
		l.Error("recipe not found", zap.String("vo.Recipe", vo.Recipe))
		return errors.New("recipe not found")

	default:
		l.Error("require -all or -recipe option")
		return errors.New("missing option")
	}
	return nil
}

func (z *Recipe) Test(c app_control.Control) error {
	return nil
}
