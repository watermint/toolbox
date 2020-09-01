package eq_queue

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/queue/eq_bundle"
	"github.com/watermint/toolbox/essentials/queue/eq_mould"
	"github.com/watermint/toolbox/essentials/queue/eq_pipe"
	"github.com/watermint/toolbox/essentials/queue/eq_progress"
)

type Opts struct {
	logger       esl.Logger
	numWorker    int
	factory      eq_pipe.Factory
	progress     eq_progress.Progress
	errorHandler eq_mould.ErrorHandler
	policy       eq_bundle.FetchPolicy
}

func (z Opts) Apply(opts ...Opt) Opts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:]...)
	}
}

func defaultOpts() Opts {
	return Opts{
		logger:    esl.Default(),
		numWorker: 1,
		progress:  nil,
		factory:   eq_pipe.NewTransientSimple(esl.Default()),
		policy:    eq_bundle.FetchSequential,
	}
}

type Opt func(o Opts) Opts

func Logger(l esl.Logger) Opt {
	return func(o Opts) Opts {
		o.logger = l
		return o
	}
}

func FetchPolicy(p eq_bundle.FetchPolicy) Opt {
	return func(o Opts) Opts {
		o.policy = p
		return o
	}
}

func NumWorker(n int) Opt {
	return func(o Opts) Opts {
		o.numWorker = n
		return o
	}
}

func Progress(p eq_progress.Progress) Opt {
	return func(o Opts) Opts {
		o.progress = p
		return o
	}
}

func Factory(f eq_pipe.Factory) Opt {
	return func(o Opts) Opts {
		o.factory = f
		return o
	}
}

func ErrorHandler(eh eq_mould.ErrorHandler) Opt {
	return func(o Opts) Opts {
		o.errorHandler = eh
		return o
	}
}
