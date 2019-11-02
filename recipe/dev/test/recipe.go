package test

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/quality/qt_test"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_recipe"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"go.uber.org/zap"
	"io/ioutil"
)

type RecipeVO struct {
	All      bool
	Recipe   string
	Resource string
}

type Recipe struct {
}

func (z *Recipe) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
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

	testResource := gjson.Parse("{}")

	if vo.Resource != "" {
		ll := l.With(zap.String("resource", vo.Resource))
		b, err := ioutil.ReadFile(vo.Resource)
		if err != nil {
			ll.Error("Unable to read resource file", zap.Error(err))
			return err
		}
		if !gjson.ValidBytes(b) {
			ll.Error("Invalid JSON format of resource file")
			return err
		}
		testResource = gjson.ParseBytes(b)
	}

	tc, err := k.Control().(*app_control_impl.Single).NewTestControl(testResource)
	if err != nil {
		l.Error("Unable to create test control", zap.Error(err))
		return err
	}

	switch {
	case vo.All:
		for _, r := range cat {
			path, name := app_recipe.Path(r)
			ll := l.With(zap.Strings("path", path), zap.String("name", name))
			if _, ok := r.(app_recipe.SecretRecipe); ok {
				ll.Info("Skip secret recipe")
				continue
			}
			ll.Info("Testing: ")

			if err := r.Test(tc); err != nil {
				ll.Error("Error", zap.Error(err))
				return err
			}
			ll.Info("Recipe test success")
		}
		l.Info("All tests passed without error")

	case vo.Recipe != "":
		for _, r := range cat {
			p := app_recipe.Key(r)
			if p != vo.Recipe {
				continue
			}
			ll := l.With(zap.String("recipeKey", p))
			ll.Info("Testing: ")

			if err := r.Test(tc); err != nil {
				ll.Error("Error", zap.Error(err))
				return err
			} else {
				ll.Info("Recipe test success")
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
	return qt_test.NoTestRequired()
}
