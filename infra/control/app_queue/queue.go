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

func selectBatchFetchPolicy(fe app_feature.Feature) eq_bundle.FetchPolicy {
	if fe.Experiment(app.ExperimentBatchRandom) {
		return eq_bundle.FetchRandom
	}
	if fe.Experiment(app.ExperimentBatchSequential) {
		return eq_bundle.FetchSequential
	}
	if fe.Experiment(app.ExperimentBatchBalance) {
		return eq_bundle.FetchBalance
	}
	return eq_bundle.FetchBalance
}

func NewSequence(lg esl.Logger, fe app_feature.Feature, ui app_ui.UI, wb app_workspace.Bundle) (seq eq_sequence.Sequence, er app_error.ErrorReport) {
	preservePath := wb.Workspace().KVS()
	preserve := eq_pipe_preserve.NewFactory(lg, preservePath)
	factory := eq_pipe.NewSimple(lg, preserve)
	progress := eq_progress.NewProgress(ea_indicator.Global())

	er = app_error.NewErrorReport(lg, wb, ui)

	seq = eq_sequence.New(
		eq_queue.AddErrorListener(er.ErrorListener),
		eq_queue.Factory(factory),
		eq_queue.FetchPolicy(selectBatchFetchPolicy(fe)),
		eq_queue.Logger(lg),
		eq_queue.NumWorker(fe.Concurrency()),
		eq_queue.Progress(progress),
		eq_queue.Verbose(fe.IsVerbose()),
	)
	return
}

func NewQueue(lg esl.Logger, fe app_feature.Feature, wb app_workspace.Bundle) (q eq_queue.Definition) {
	preservePath := wb.Workspace().KVS()
	preserve := eq_pipe_preserve.NewFactory(lg, preservePath)
	factory := eq_pipe.NewSimple(lg, preserve)
	progress := eq_progress.NewProgress(ea_indicator.Global())

	return eq_queue.New(
		eq_queue.Factory(factory),
		eq_queue.FetchPolicy(selectBatchFetchPolicy(fe)),
		eq_queue.Logger(lg),
		eq_queue.NumWorker(fe.Concurrency()),
		eq_queue.Progress(progress),
		eq_queue.Verbose(fe.IsVerbose()),
	)
}
