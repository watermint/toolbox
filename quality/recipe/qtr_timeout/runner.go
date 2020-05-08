package qtr_timeout

import (
	"context"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"strings"
	"time"
)

var (
	recipeTimeout = 30 * time.Second
)

func execRecipeTest(c app_control.Control, r rc_recipe.Recipe, spec rc_recipe.Spec, timeoutEnabled, useMock bool) error {
	path, name := spec.Path()
	l := c.Log().With(es_log.Strings("path", path), es_log.String("name", name), es_log.Bool("timeoutEnabled", timeoutEnabled))
	l.Debug("Testing: ")
	execName := strings.Join(append(path, name), "-")

	ct := c.WithFeature(c.Feature().AsTest(useMock))
	return app_control_impl.WithForkedQuiet(ct, execName, func(cf app_control.Control) error {
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

func RunRecipeTestWithTimeout(c app_control.Control, r rc_recipe.Recipe, timeoutEnabled, useMock bool) (err error) {
	spec := rc_spec.New(r)
	path, name := spec.Path()
	l := c.Log().With(es_log.Strings("path", path), es_log.String("name", name))

	// Run without timeout
	if spec.IsIrreversible() || !timeoutEnabled {
		l.Debug("Run recipe without timeout")
		return execRecipeTest(c, r, spec, false, useMock)
	}

	ctx, cancel := context.WithTimeout(context.Background(), recipeTimeout)
	defer cancel()

	result := make(chan error)
	go func() {
		errRecipe := execRecipeTest(c, r, spec, true, useMock)
		result <- errRecipe
	}()

	select {
	case errRecipe := <-result:
		l.Debug("Recipe finished without timeout", es_log.Error(errRecipe))
		return errRecipe

	case <-ctx.Done():
		l.Info("Recipe test finished with timeout")
		return nil
	}
}
