package log

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/control/app_workspace"
)

func getHistories(path mo_string.OptionalString) (histories []app_job.History, err error) {
	l := esl.Default()

	home := ""
	if path.IsExists() {
		home = path.Value()
	}

	// default non transient workspace
	ws, err := app_workspace.NewWorkspace(home, false)
	if err != nil {
		return nil, err
	}

	historian := app_job_impl.NewHistorian(ws)
	histories, err = historian.Histories()
	if err != nil {
		l.Debug("Unable to retrieve histories", esl.Error(err))
		return nil, err
	}
	if len(histories) < 1 {
		l.Debug("No log found", esl.Any("histories", histories))
	}
	return
}
