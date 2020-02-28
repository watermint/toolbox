package test

import (
	"errors"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"io/ioutil"
)

type Recipe struct {
	All      bool
	Recipe   string
	Resource string
}

func (z *Recipe) Preset() {
}

func (z *Recipe) Exec(c app_control.Control) error {
	cl := c.(app_control_launcher.ControlLauncher)
	cat := cl.Catalogue()
	l := c.Log()

	testResource := gjson.Parse("{}")

	if z.Resource != "" {
		ll := l.With(zap.String("resource", z.Resource))
		b, err := ioutil.ReadFile(z.Resource)
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

	tc, err := c.(*app_control_impl.Single).NewTestControl(testResource)
	if err != nil {
		l.Error("Unable to create test control", zap.Error(err))
		return err
	}

	switch {
	case z.All:
		for _, r := range cat.Recipes() {
			rs := rc_spec.New(r)
			path, name := rs.Path()
			ll := l.With(zap.Strings("path", path), zap.String("name", name))
			if rs.IsSecret() {
				ll.Info("Skip secret recipe")
				continue
			}
			ll.Info("Testing: ")

			if err, _ := qt_recipe.RecipeError(l, r.Test(tc)); err != nil {
				ll.Error("Error", zap.Error(err))
				return err
			}
			ll.Info("Recipe test success")
		}
		l.Info("All tests passed without error")

	case z.Recipe != "":
		for _, r := range cat.Recipes() {
			p := rc_recipe.Key(r)
			if p != z.Recipe {
				continue
			}
			ll := l.With(zap.String("recipeKey", p))
			ll.Info("Testing: ")

			if err, _ := qt_recipe.RecipeError(l, r.Test(tc)); err != nil {
				ll.Error("Error", zap.Error(err))
				return err
			} else {
				ll.Info("Recipe test success")
				return nil
			}
		}
		l.Error("recipe not found", zap.String("vo.Recipe", z.Recipe))
		return errors.New("recipe not found")

	default:
		l.Error("require -all or -recipe option")
		return errors.New("missing option")
	}
	return nil
}

func (z *Recipe) Test(c app_control.Control) error {
	return qt_endtoend.NoTestRequired()
}
