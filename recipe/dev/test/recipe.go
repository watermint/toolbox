package test

import (
	"errors"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
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

	cn := strings.Join(path, "-") + "-" + name
	return app_workspace.WithFork(c.WorkBundle(), cn, func(fwb app_workspace.Bundle) error {
		cf := c.WithBundle(fwb).WithFeature(c.Feature().AsTest(false))
		if !z.Verbose {
			cf = cf.WithFeature(cf.Feature().AsQuiet()).WithUI(app_ui.NewDiscard(c.Messages(), c.Log()))
		}

		l.Debug("Testing: ")

		timeStart := time.Now()
		if err, _ := qt_errors.ErrorsForTest(l, r.Test(cf)); err != nil {
			l.Error("Error", es_log.Error(err))
			return err
		}
		timeEnd := time.Now()
		l.Info("Recipe test success", es_log.Int64("duration", timeEnd.Sub(timeStart).Milliseconds()))
		return nil
	})
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
	return qt_errors.ErrorNoTestRequired
}
