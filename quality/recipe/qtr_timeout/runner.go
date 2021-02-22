package qtr_timeout

import (
	"context"
	"github.com/watermint/toolbox/essentials/log/esl"
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

type RecipeTestResult struct {
	Path           string `json:"path"`
	Name           string `json:"name"`
	Skip           bool   `json:"skip"`
	TimeoutEnabled bool   `json:"timeout_enabled"`
	UseMock        bool   `json:"use_mock"`
	Timeout        bool   `json:"timeout"`
	Duration       int64  `json:"duration"`
	NoError        bool   `json:"no_error"`
	Error          string `json:"error"`
}

func execRecipeTest(c app_control.Control, r rc_recipe.Recipe, spec rc_recipe.Spec, timeoutEnabled, useMock bool) (rtr RecipeTestResult, err error) {
	path, name := spec.Path()
	l := c.Log().With(esl.Strings("path", path), esl.String("name", name), esl.Bool("timeoutEnabled", timeoutEnabled))
	l.Debug("Testing: ")
	execName := strings.Join(append(path, name), "-")

	ct := c.WithFeature(c.Feature().AsTest(useMock))
	rtr.Path = strings.Join(path, " ")
	rtr.Name = name
	rtr.TimeoutEnabled = timeoutEnabled
	rtr.UseMock = useMock

	lastErr := app_control_impl.WithForkedQuiet(ct, execName, func(cf app_control.Control) error {
		timeStart := time.Now()
		if err, _ := qt_errors.ErrorsForTest(l, r.Test(cf)); err != nil {
			l.Error("Error", esl.Error(err))
			rtr.Error = err.Error()
			return err
		}
		timeEnd := time.Now()
		testDur := timeEnd.Sub(timeStart).Milliseconds()
		l.Info("Recipe test success", esl.Int64("duration", testDur))
		rtr.Duration = testDur
		return nil
	})
	if lastErr != nil {
		rtr.Error = lastErr.Error()
		rtr.NoError = false
	} else {
		rtr.NoError = true
	}
	return rtr, lastErr
}

func RunRecipeTestWithTimeout(c app_control.Control, r rc_recipe.Recipe, timeoutEnabled, useMock bool) (rtr RecipeTestResult, err error) {
	spec := rc_spec.New(r)
	path, name := spec.Path()
	l := c.Log().With(esl.Strings("path", path), esl.String("name", name))

	rtr.Path = strings.Join(path, " ")
	rtr.Name = name
	rtr.TimeoutEnabled = timeoutEnabled
	rtr.UseMock = useMock

	// Run without timeout
	if spec.IsIrreversible() || !timeoutEnabled {
		l.Debug("Run recipe without timeout")
		return execRecipeTest(c, r, spec, false, useMock)
	}

	ctx, cancel := context.WithTimeout(context.Background(), recipeTimeout)
	defer cancel()

	result := make(chan error)
	var errRecipe error
	go func() {
		rtr, errRecipe = execRecipeTest(c, r, spec, true, useMock)
		result <- errRecipe
	}()

	select {
	case errRecipe := <-result:
		l.Debug("Recipe finished without timeout", esl.Error(errRecipe))
		return rtr, errRecipe

	case <-ctx.Done():
		l.Info("Recipe test finished with timeout")
		rtr.Timeout = true
		rtr.NoError = true
		return rtr, nil
	}
}
