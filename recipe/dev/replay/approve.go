package replay

import (
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/network/nw_replay"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_replay"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"os"
	"path/filepath"
	"strings"
)

type Approve struct {
	rc_recipe.RemarkSecret
	Id            string
	WorkspacePath mo_string.OptionalString
	ReplayPath    mo_string.OptionalString
	Name          mo_string.OptionalString
}

func (z *Approve) Preset() {
}

func (z *Approve) Exec(c app_control.Control) error {
	home := ""
	if z.WorkspacePath.IsExists() {
		home = z.WorkspacePath.Value()
	}
	l := c.Log().With(esl.String("JobId", z.Id))

	// default non transient workspace
	ws, err := app_workspace.NewWorkspace(home, false)
	if err != nil {
		return err
	}

	replayPath, err := rc_replay.ReplayPath(z.ReplayPath)
	if err != nil {
		return err
	}

	replayFolderInfo, err := os.Lstat(replayPath)
	switch {
	case os.IsNotExist(err):
		l.Debug("Not found, try create the folder")
		if err := os.MkdirAll(replayPath, 0755); err != nil {
			l.Debug("Unable to create the folder", esl.Error(err))
			return err
		}
	case !replayFolderInfo.IsDir():
		l.Debug("The path is not a folder")
		return errors.New("replay path is not a folder")
	}

	replay := rc_replay.New(c.Log())

	historian := app_job_impl.NewHistorian(ws)
	histories, err := historian.Histories()
	for _, history := range histories {
		if history.JobId() != z.Id {
			continue
		}

		archiveName := strings.ReplaceAll(es_filepath.Escape(history.RecipeName()), " ", "-")
		if z.Name.IsExists() {
			archiveName += "_" + z.Name.Value()
		}
		archiveName = archiveName + ".zip"

		l.Debug("Replay the recipe",
			esl.String("jobId", history.JobId()),
			esl.String("recipe", history.RecipeName()),
			esl.String("archiveName", archiveName))

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
		if err := replay.Replay(history.Job(), forkCtl); err == nw_replay.ErrorNoReplayFound {
			l.Warn("No replay data found. Could not approve it.")
			return err
		}

		return replay.Preserve(history.Job(), filepath.Join(replayPath, archiveName))
	}

	return ErrorJobNotFound
}

func (z *Approve) Test(c app_control.Control) error {
	return qt_errors.ErrorHumanInteractionRequired
}
