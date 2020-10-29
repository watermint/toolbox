package app_queue

import (
	"github.com/watermint/toolbox/essentials/ambient/ea_indicator"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe_preserve"
	"github.com/watermint/toolbox/essentials/queue/eq_progress"
	"github.com/watermint/toolbox/essentials/queue/eq_queue"
	"github.com/watermint/toolbox/essentials/queue/eq_sequence"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_error"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

func NewQueue(lg esl.Logger, fe app_feature.Feature, ui app_ui.UI, wb app_workspace.Bundle) (seq eq_sequence.Sequence, er app_error.ErrorReport) {
	preservePath := wb.Workspace().KVS()
	preserve := eq_pipe_preserve.NewFactory(lg, preservePath)
	factory := eq_pipe.NewSimple(lg, preserve)
	progress := eq_progress.NewProgress(ea_indicator.Global())

	er = app_error.NewErrorReport(lg, wb, ui)

	batchPolicy := eq_bundle.FetchSequential
	if fe.Experiment(app.ExperimentBatchRandom) {
		batchPolicy = eq_bundle.FetchRandom
	}
	if fe.Experiment(app.ExperimentBatchSequential) {
		batchPolicy = eq_bundle.FetchSequential
	}
	lg.Debug("Queue execution policy", esl.Any("policy", batchPolicy))

	seq = eq_sequence.New(
		eq_queue.Logger(lg),
		eq_queue.FetchPolicy(batchPolicy),
		eq_queue.Progress(progress),
		eq_queue.NumWorker(fe.Concurrency()),
		eq_queue.Factory(factory),
		eq_queue.ErrorHandler(er.ErrorHandler),
	)
	return
}
