package app_workflow

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/watermint/toolbox/essentials/encoding/es_unicode"
	"github.com/watermint/toolbox/essentials/go/es_goroutine"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

const (
	RunBookName     = "tbx.runbook"
	RunBookTestName = "tbx.runbook.test"
	RootWorkerName  = "root"
)

type MsgRunBook struct {
	ErrorFailedParseArgs           app_msg.Message
	ErrorInvalidVersion            app_msg.Message
	ErrorMissingInvalidCommand     app_msg.Message
	ErrorMissingReservedWorkerName app_msg.Message
	ErrorMissingStepName           app_msg.Message
	ErrorMissingStepsOrWorkers     app_msg.Message
	ErrorMissingWorkerName         app_msg.Message
	ErrorRecipeFailed              app_msg.Message
	ErrorUnableStart               app_msg.Message
	ProgressRecipeFinished         app_msg.Message
	ProgressRecipeStart            app_msg.Message
}

var (
	MRunBook = app_msg.Apply(&MsgRunBook{}).(*MsgRunBook)
)

type RunStep struct {
	Name string
	Args []string `json:"args"`
}

type RunWorker struct {
	Name  string     `json:"name"`
	Steps []*RunStep `json:"steps"`
}

type RunBook struct {
	Version int          `json:"version"`
	Steps   []*RunStep   `json:"steps"`
	Workers []*RunWorker `json:"workers"`
}

func (z *RunBook) Verify(c app_control.Control) (lastErr error) {
	cat := app_catalogue.Current()
	rg := cat.RootGroup()
	ui := c.UI()
	lastErr = nil
	if z.Version != 1 {
		ui.Error(MRunBook.ErrorInvalidVersion.With("Version", z.Version))
		lastErr = errors.New("invalid version number")
	}
	if (z.Steps == nil || len(z.Steps) < 1) && (z.Workers == nil || len(z.Workers) < 1) {
		ui.Error(MRunBook.ErrorMissingStepsOrWorkers)
		lastErr = errors.New("missing `steps` or `workers` element")
	}
	scanSteps := func(steps []*RunStep) {
		for _, s := range z.Steps {
			if s.Name == "" {
				ui.Error(MRunBook.ErrorMissingStepName)
				lastErr = errors.New("`name` is missing in the step")
			}
			if _, r, _, err := rg.Select(s.Args); r == nil || err != nil {
				ui.Error(MRunBook.ErrorMissingInvalidCommand.With("Args", strings.Join(s.Args, " ")))
				lastErr = errors.New("invalid command")
			}
		}
	}
	if z.Steps != nil {
		scanSteps(z.Steps)
	}
	if z.Workers != nil {
		for _, w := range z.Workers {
			switch w.Name {
			case "":
				ui.Error(MRunBook.ErrorMissingWorkerName)
				lastErr = errors.New("missing worker name")

			case RootWorkerName:
				ui.Error(MRunBook.ErrorMissingReservedWorkerName)
				lastErr = errors.New("reserved worker name")
			}
			scanSteps(w.Steps)
		}
	}
	return lastErr
}

func (z *RunBook) runRecipe(workerName string, step *RunStep, rg rc_group.Group, c app_control.Control) error {
	l := c.Log().With(es_log.String("workerName", workerName), es_log.String("stepName", step.Name))
	ui := c.UI()
	ui.Info(MRunBook.ProgressRecipeStart.With("Worker", workerName).With("Args", strings.Join(step.Args, " ")))
	_, rOrig, args, err := rg.Select(step.Args)
	if err != nil || rOrig == nil {
		return errors.New("invalid command")
	}
	r := rOrig.New()
	f := flag.NewFlagSet(r.CliPath(), flag.ContinueOnError)
	r.SetFlags(f, c.UI())
	if err := f.Parse(args); err != nil {
		ui.Error(MRunBook.ErrorFailedParseArgs.With("Error", err))
		return err
	}

	// fork control
	runName := workerName + "-" + step.Name
	return app_workspace.WithFork(c.WorkBundle(), runName, func(fwb app_workspace.Bundle) error {
		cf := c.WithBundle(fwb)

		l.Debug("Run step", es_log.Any("vo", r.Debug()))
		if err = rc_exec.ExecSpec(cf, r, rc_recipe.NoCustomValues); err != nil {
			ui.Error(MRunBook.ErrorRecipeFailed.With("Error", err))
			return err
		}
		ui.Info(MRunBook.ProgressRecipeFinished)
		return nil
	})
}

func (z *RunBook) runWorker(wg *sync.WaitGroup, workerErrors []error, workerName string, steps []*RunStep, rg rc_group.Group, c app_control.Control) error {
	wg.Add(1)
	defer wg.Done()

	l := c.Log().With(es_log.String("worker", workerName), es_log.String("goroutine", es_goroutine.GetGoRoutineName()))
	l.Info("Worker start")
	for _, step := range steps {
		if err := z.runRecipe(workerName, step, rg, c); err != nil {
			l.Error("Failed exec recipe", es_log.Error(err))
			workerErrors = append(workerErrors, err)
			return err
		}
	}
	l.Info("Worker finished")
	return nil
}

func (z *RunBook) Run(c app_control.Control) error {
	l := c.Log()

	cat := app_catalogue.Current()
	rg := cat.RootGroup()
	wg := sync.WaitGroup{}
	workerErrors := make([]error, 0)

	l.Debug("Run runbook")
	if z.Steps != nil && len(z.Steps) > 0 {
		go z.runWorker(&wg, workerErrors, RootWorkerName, z.Steps, rg, c)
	}
	if z.Workers != nil && len(z.Workers) > 0 {
		for _, worker := range z.Workers {
			go z.runWorker(&wg, workerErrors, worker.Name, worker.Steps, rg, c)
		}
	}

	l.Info("Waiting for workers")
	wg.Wait()
	l.Info("Workers finished")
	if len(workerErrors) > 0 {
		l.Error("One or more errors from workers", es_log.Errors("worker", workerErrors))
		return workerErrors[0]
	}
	return nil
}

func NewRunBook(path string) (rb *RunBook, found bool) {
	l := es_log.Default()
	_, err := os.Lstat(path)
	if err != nil {
		return nil, false
	}

	content, err := es_unicode.BomAwareReadBytes(path)
	if err != nil {
		l.Error("Unable to read the runbook file", es_log.Error(err))
		return nil, false
	}

	rb = &RunBook{}
	if err = json.Unmarshal(content, rb); err != nil {
		l.Error("Unable to unmarshal the runbook file", es_log.Error(err))
		return nil, false
	}

	return rb, true
}

func DefaultRunBook(forTest bool) (path string, rb *RunBook, found bool) {
	name := RunBookName
	if forTest {
		name = RunBookTestName
	}
	path = filepath.Join(filepath.Dir(os.Args[0]), name)
	rb, found = NewRunBook(path)
	return path, rb, found
}
