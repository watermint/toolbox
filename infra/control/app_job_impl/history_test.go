package app_job_impl

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_budget"
	"github.com/watermint/toolbox/infra/control/app_opt"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"testing"
)

func TestHistory(t *testing.T) {
	qt_file.TestWithTestFolder(t, "launch", false, func(path string) {
		wb, err := app_workspace.NewBundle(path, app_budget.BudgetUnlimited, esl.LevelInfo, false)
		if err != nil {
			t.Error(err)
			return
		}
		mc := app_msg_container_impl.NewSingleWithMessages(map[string]string{})
		ui := app_ui.NewDiscard(mc, wb.Logger().Logger())
		spec := rc_spec.New(&AppJobTestRecipe{})
		launcher := NewLauncher(ui, wb, app_opt.Default(), spec)
		ctl, err := launcher.Up()
		if err != nil {
			t.Error(err)
		}
		ctl.Log().Debug("Hello")
		launcher.Down(nil, ctl)

		his := NewHistorian(wb.Workspace())
		jobs, err := his.Histories()
		if err != nil {
			t.Error(err)
		}
		if len(jobs) != 1 {
			t.Error(jobs)
		}

		job := jobs[0]
		if job.JobId() == "" {
			t.Error(job.JobId())
		}

		// should not found, because the recipe is not in the catalogue
		if r, found := job.Recipe(); found {
			t.Error(r, found)
		}

		if job.AppName() == "" {
			t.Error(job.AppName())
		}
		if job.AppVersion() == "" {
			t.Error(job.AppVersion())
		}
		if job.JobPath() == "" {
			t.Error(job.JobPath())
		}
		if job.RecipeName() != "github.com watermint toolbox infra control app_job_impl app_job_test_recipe" {
			t.Error(job.RecipeName())
		}
		if v, found := job.TimeStart(); !found {
			t.Error(v, found)
		}
		if v, found := job.TimeFinish(); !found {
			t.Error(v, found)
		}
	})
}

func TestHistory_Archive(t *testing.T) {
	qt_file.TestWithTestFolder(t, "launch", false, func(path string) {
		wb, err := app_workspace.NewBundle(path, app_budget.BudgetUnlimited, esl.LevelInfo, false)
		if err != nil {
			t.Error(err)
			return
		}
		mc := app_msg_container_impl.NewSingleWithMessages(map[string]string{})
		ui := app_ui.NewDiscard(mc, wb.Logger().Logger())
		spec := rc_spec.New(&AppJobTestRecipe{})
		launcher := NewLauncher(ui, wb, app_opt.Default(), spec)
		ctl, err := launcher.Up()
		if err != nil {
			t.Error(err)
		}
		launcher.Down(nil, ctl)

		his := NewHistorian(wb.Workspace())
		jobs, err := his.Histories()
		if err != nil {
			t.Error(err)
		}
		if len(jobs) != 1 {
			t.Error(jobs)
		}

		job := jobs[0]

		// the path should exist
		if f, err := os.Lstat(job.JobPath()); err != nil && f.IsDir() {
			t.Error(err)
		}

		if arcPath, err := job.Archive(); err != nil {
			t.Error(path, err)
		} else if f, err := os.Lstat(arcPath); err != nil && !f.IsDir() {
			t.Error(f, err)
		}

		// the path should not exist
		if _, err := os.Lstat(job.JobPath()); err == nil {
			t.Error(err)
		}
	})
}

func TestHistory_Delete(t *testing.T) {
	qt_file.TestWithTestFolder(t, "launch", false, func(path string) {
		wb, err := app_workspace.NewBundle(path, app_budget.BudgetUnlimited, esl.LevelInfo, false)
		if err != nil {
			t.Error(err)
			return
		}
		mc := app_msg_container_impl.NewSingleWithMessages(map[string]string{})
		ui := app_ui.NewDiscard(mc, wb.Logger().Logger())
		spec := rc_spec.New(&AppJobTestRecipe{})
		launcher := NewLauncher(ui, wb, app_opt.Default(), spec)
		ctl, err := launcher.Up()
		if err != nil {
			t.Error(err)
		}
		launcher.Down(nil, ctl)

		his := NewHistorian(wb.Workspace())
		jobs, err := his.Histories()
		if err != nil {
			t.Error(err)
		}
		if len(jobs) != 1 {
			t.Error(jobs)
		}

		job := jobs[0]

		// the path should exist
		if f, err := os.Lstat(job.JobPath()); err != nil && f.IsDir() {
			t.Error(err)
		}

		if err := job.Delete(); err != nil {
			t.Error(path, err)
		}

		// the path should not exist
		if _, err := os.Lstat(job.JobPath()); err == nil {
			t.Error(err)
		}
	})
}
