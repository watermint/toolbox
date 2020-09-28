package replay

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_replay"
	"strings"
)

var (
	ErrorJobNotFound = errors.New("job id not found")
)

type Recipe struct {
	rc_recipe.RemarkSecret
	JobId string
	Path  mo_string.OptionalString
}

func (z *Recipe) Preset() {
}

func (z *Recipe) Exec(c app_control.Control) error {
	l := c.Log()

	home := ""
	if z.Path.IsExists() {
		home = z.Path.Value()
	}

	// default non transient workspace
	ws, err := app_workspace.NewWorkspace(home, false)
	if err != nil {
		return err
	}

	replay := rc_replay.New(c.Log())

	historian := app_job_impl.NewHistorian(ws)
	histories, err := historian.Histories()
	for _, history := range histories {
		if history.JobId() != z.JobId {
			continue
		}

		forkName := strings.ReplaceAll(es_filepath.Escape(history.RecipeName()), " ", "-")
		l.Debug("Replay the recipe",
			esl.String("jobId", history.JobId()),
			esl.String("recipe", history.RecipeName()),
			esl.String("forkName", forkName))
		forkBundle, err := app_workspace.ForkBundle(c.WorkBundle(), forkName)
		if err != nil {
			l.Debug("Unable to fork the bundle", esl.Error(err))
			return err
		}
		forkCtl := c.WithBundle(forkBundle)
		return replay.Replay(history.Job(), forkCtl)
	}

	return ErrorJobNotFound
}

func (z *Recipe) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Recipe{}, func(r rc_recipe.Recipe) {
		m := r.(*Recipe)
		m.JobId = "1234"
	})
	if err != ErrorJobNotFound {
		return err
	}
	return nil
}
