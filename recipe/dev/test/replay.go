package test

import (
	"errors"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
)

var (
	ErrorJobNotFound = errors.New("job id not found")
)

type Replay struct {
	rc_recipe.RemarkSecret
	rc_recipe.RemarkTransient
	JobId string
	Path  mo_string.OptionalString
}

func (z *Replay) Preset() {
}

func (z *Replay) Exec(c app_control.Control) error {
	home := ""
	if z.Path.IsExists() {
		home = z.Path.Value()
	}

	// default non transient workspace
	ws, err := app_workspace.NewWorkspace(home, false)
	if err != nil {
		return err
	}

	historian := app_job_impl.NewHistorian(ws)
	histories, err := historian.Histories()
	for _, history := range histories {
		if history.JobId() != z.JobId {
			continue
		}

		return nil
	}

	return ErrorJobNotFound
}

func (z *Replay) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Replay{}, func(r rc_recipe.Recipe) {
		m := r.(*Replay)
		m.JobId = "1234"
	})
	if err != ErrorJobNotFound {
		return err
	}
	return nil
}
