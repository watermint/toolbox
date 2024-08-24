package rc_exec

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/network/nw_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_replay"
	"runtime"
)

var (
	ErrorPanic = errors.New("the program aborted")
)

type MsgPanic struct {
	ErrorRecipePanic                app_msg.Message
	ErrorCrashReport                app_msg.Message
	ErrorInvalidOrExpiredOAuthToken app_msg.Message
}

var (
	MPanic = app_msg.Apply(&MsgPanic{}).(*MsgPanic)
)

func Exec(ctl app_control.Control, r rc_recipe.Recipe, custom func(r rc_recipe.Recipe)) error {
	return ExecSpec(ctl, rc_spec.New(r), custom)
}

// Execute with mock test mode
func ExecMock(ctl app_control.Control, r rc_recipe.Recipe, custom func(r rc_recipe.Recipe)) error {
	cte := ctl.WithFeature(ctl.Feature().AsTest(true))
	return ExecSpec(cte, rc_spec.New(r), custom)
}

// Execute with mock test mode
func ExecReplay(ctl app_control.Control, r rc_recipe.Recipe, replayName string, custom func(r rc_recipe.Recipe)) error {
	replay, err := qt_replay.LoadReplay(replayName)
	if _, ok := err.(*json.SyntaxError); ok {
		return err
	}
	if err != nil {
		return qt_errors.ErrorNotEnoughResource
	}
	cte := ctl.WithFeature(ctl.Feature().AsSeqReplayTest(replay))
	return ExecSpec(cte, rc_spec.New(r), custom)
}

func ExecSpec(ctl app_control.Control, spec rc_recipe.Spec, custom func(r rc_recipe.Recipe)) error {
	return DoSpec(ctl, spec, custom, func(r rc_recipe.Recipe, ctl app_control.Control) error {
		return r.Exec(ctl)
	})
}

func DoSpec(ctl app_control.Control, spec rc_recipe.Spec, custom func(r rc_recipe.Recipe), do func(r rc_recipe.Recipe, ctl app_control.Control) error) error {
	l := ctl.Log()
	if spec == nil {
		l.Debug("Spec not found")
		return errors.New("no spec found")
	}
	l = l.With(esl.String("cliPath", spec.CliPath()))
	scr, err := spec.SpinUp(ctl, custom)
	if err != nil {
		return err
	}

	rcpErr := doSpecInternal(spec, scr, ctl, do)

	if rcpErr != nil {
		for _, handler := range spec.ErrorHandlers() {
			if handler.Handle(ctl.UI(), rcpErr) {
				break
			}
		}
		switch rcpErr {
		case nw_auth.ErrorInvalidOrExpiredRefreshToken:
			ctl.UI().Error(MPanic.ErrorInvalidOrExpiredOAuthToken)
		}
	}

	// mark indicators as done before spin down
	ea_indicator.Global().Done()

	if err = spec.SpinDown(ctl); err != nil {
		l.Debug("Spin down error", esl.Error(err))
	}
	return rcpErr
}

type ErrorTrace struct {
	Depth int    `json:"depth"`
	File  string `json:"file"`
	Line  int    `json:"line"`
}

func doSpecInternal(spec rc_recipe.Spec, scr rc_recipe.Recipe, ctl app_control.Control, do func(r rc_recipe.Recipe, ctl app_control.Control) error) (rcpErr error) {
	l := ctl.Log()
	ui := ctl.UI()

	defer func() {
		var rErr interface{}
		if ctl.Feature().IsProduction() {
			rErr = recover()
		}
		if rErr != nil { // && ctl.Feature().IsProduction() {
			l.Debug("Recovery from panic", esl.Any("err", rErr))
			traces := make([]ErrorTrace, 0)
			for depth := 0; ; depth++ {
				_, file, line, ok := runtime.Caller(depth)
				if !ok {
					break
				}
				l.Debug("Trace",
					esl.Int("Depth", depth),
					esl.String("File", file),
					esl.Int("Line", line),
				)
				traces = append(traces, ErrorTrace{
					Depth: depth,
					File:  file,
					Line:  line,
				})
			}
			traceLog, _ := json.Marshal(traces)

			ui.Error(MPanic.ErrorRecipePanic)
			ui.Error(MPanic.ErrorCrashReport.
				With("Version", app_definitions.Version.String()).
				With("OS", runtime.GOOS).
				With("Recipe", spec.CliPath()).
				With("Reason", rErr).
				With("Trace", string(traceLog)))

			rcpErr = ErrorPanic
		}
	}()

	rcpErr = do(scr, ctl)

	return rcpErr
}
