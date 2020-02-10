package app_workflow

import (
	"encoding/json"
	"errors"
	"flag"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/control/app_root"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_group"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/util/ut_encoding"
	"github.com/watermint/toolbox/infra/util/ut_runtime"
	"go.uber.org/zap"
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
	cl, ok := c.(app_control_launcher.ControlLauncher)
	if !ok {
		c.Log().Debug("Skip verification")
		return nil
	}
	cat := cl.Catalogue()
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

func (z *RunBook) runRecipe(workerName string, step *RunStep, rg rc_group.Group, cp app_control.Control, cf app_control_launcher.ControlFork) error {
	l := cp.Log().With(zap.String("workerName", workerName), zap.String("stepName", step.Name))
	ui := cp.UI()
	ui.Info(MRunBook.ProgressRecipeStart.With("Worker", workerName).With("Args", strings.Join(step.Args, " ")))
	_, rOrig, args, err := rg.Select(step.Args)
	if err != nil || rOrig == nil {
		return errors.New("invalid command")
	}
	r := rOrig.New()
	f := flag.NewFlagSet(r.CliPath(), flag.ContinueOnError)
	r.SetFlags(f, cp.UI())
	if err := f.Parse(args); err != nil {
		ui.Error(MRunBook.ErrorFailedParseArgs.With("Error", err))
		return err
	}
	so := make([]app_control.UpOpt, 0)
	so = append(so, app_control.RecipeName(r.CliPath()))
	so = append(so, app_control.RecipeOptions(r.Debug()))

	fc, err := cf.Fork(workerName+"-"+step.Name, so...)
	if err != nil {
		ui.Error(MRunBook.ErrorUnableStart.With("Error", err))
		return err
	}
	defer fc.Down()
	l.Debug("Run step", zap.Any("vo", r.Debug()))
	if err = rc_exec.ExecSpec(fc, r, rc_recipe.NoCustomValues); err != nil {
		ui.Error(MRunBook.ErrorRecipeFailed.With("Error", err))
		return err
	}
	ui.Info(MRunBook.ProgressRecipeFinished)
	return nil
}

func (z *RunBook) runWorker(wg *sync.WaitGroup, workerErrors []error, workerName string, steps []*RunStep, rg rc_group.Group, cp app_control.Control, cf app_control_launcher.ControlFork) error {
	wg.Add(1)
	defer wg.Done()

	l := cp.Log().With(zap.String("worker", workerName), zap.String("goroutine", ut_runtime.GetGoRoutineName()))
	l.Info("Worker start")
	for _, step := range steps {
		if err := z.runRecipe(workerName, step, rg, cp, cf); err != nil {
			l.Error("Failed exec recipe", zap.Error(err))
			workerErrors = append(workerErrors, err)
			return err
		}
	}
	l.Info("Worker finished")
	return nil
}

func (z *RunBook) Run(c app_control.Control) error {
	l := c.Log()
	cl, ok := c.(app_control_launcher.ControlLauncher)
	if !ok {
		l.Debug("Skip run")
		return nil
	}
	cf, ok := c.(app_control_launcher.ControlFork)
	if !ok {
		l.Debug("Skip run")
		return nil
	}

	cat := cl.Catalogue()
	rg := cat.RootGroup()
	wg := sync.WaitGroup{}
	workerErrors := make([]error, 0)

	l.Debug("Run runbook")
	if z.Steps != nil && len(z.Steps) > 0 {
		go z.runWorker(&wg, workerErrors, RootWorkerName, z.Steps, rg, c, cf)
	}
	if z.Workers != nil && len(z.Workers) > 0 {
		for _, worker := range z.Workers {
			go z.runWorker(&wg, workerErrors, worker.Name, worker.Steps, rg, c, cf)
		}
	}

	l.Info("Waiting for workers")
	wg.Wait()
	l.Info("Workers finished")
	if len(workerErrors) > 0 {
		l.Error("One or more errors from workers", zap.Errors("worker", workerErrors))
		return workerErrors[0]
	}
	return nil
}

func NewRunBook(path string) (rb *RunBook, found bool) {
	l := app_root.Log()
	_, err := os.Lstat(path)
	if err != nil {
		return nil, false
	}

	content, err := ut_encoding.BomAwareReadBytes(path)
	if err != nil {
		l.Error("Unable to read the runbook file", zap.Error(err))
		return nil, false
	}

	rb = &RunBook{}
	if err = json.Unmarshal(content, rb); err != nil {
		l.Error("Unable to unmarshal the runbook file", zap.Error(err))
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
