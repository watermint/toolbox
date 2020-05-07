package test

import (
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/recipe/dev"
	"strings"
	"time"
)

type Recipe struct {
	rc_recipe.RemarkSecret
	All     bool
	Recipe  mo_string.OptionalString
	Verbose bool
}

func (z *Recipe) Preset() {
}

func (z *Recipe) runSingle(c app_control.Control, r rc_recipe.Recipe) error {
	rs := rc_spec.New(r)
	path, name := rs.Path()
	l := c.Log().With(es_log.Strings("path", path), es_log.String("name", name))

	if rs.IsSecret() {
		l.Info("Skip secret recipe")
		return nil
	}

	cn := strings.Join(append(path, name), "-")
	cf, err := app_control_impl.ForkQuiet(c.WithFeature(c.Feature().AsTest(false)), cn)
	if err != nil {
		return err
	}
	defer cf.WorkBundle().Close()
	l.Debug("Testing: ")

	timeStart := time.Now()
	if err, _ := qt_errors.ErrorsForTest(l, r.Test(cf)); err != nil {
		l.Error("Error", es_log.Error(err))
		return err
	}
	timeEnd := time.Now()
	l.Info("Recipe test success", es_log.Int64("duration", timeEnd.Sub(timeStart).Milliseconds()))
	return nil
}

func (z *Recipe) runAll(c app_control.Control) error {
	cat := app_catalogue.Current()
	l := c.Log()

	for _, r := range cat.Recipes() {
		if err := z.runSingle(c, r); err != nil {
			return err
		}
	}
	l.Info("All tests passed without error")
	return nil
}

func (z *Recipe) Exec(c app_control.Control) error {
	cat := app_catalogue.Current()
	l := c.Log()

	switch {
	case z.All:
		if err := z.runAll(c); err != nil {
			return err
		}

	case z.Recipe.IsExists():
		for _, r := range cat.Recipes() {
			p := rc_recipe.Key(r)
			if p != z.Recipe.Value() {
				continue
			}
			if err := z.runSingle(c, r); err != nil {
				return err
			}
			l.Info("Recipe test success")
			return nil
		}
		l.Error("recipe not found", es_log.String("vo.Recipe", z.Recipe.Value()))
		return errors.New("recipe not found")

	default:
		l.Error("require -all or -recipe option")
		return errors.New("missing option")
	}
	return nil
}

func (z *Recipe) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Recipe{}, func(r rc_recipe.Recipe) {
		m := r.(*Recipe)
		m.Recipe = mo_string.NewOptional(rc_recipe.Key(&dev.Echo{}))
	})
}
