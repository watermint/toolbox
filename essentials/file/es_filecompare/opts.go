package es_filecompare

import (
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type Opt func(o Opts) Opts

type Opts struct {
	dontCompareTime      bool
	dontCompareContent   bool
	handlerFileDiff      FileDiff
	handlerTypeDiff      TypeDiff
	handlerSameFile      SameFile
	handlerMissingSource MissingSource
	handlerMissingTarget MissingTarget
	logger               esl.Logger
}

func (z Opts) Log() esl.Logger {
	if z.logger != nil {
		return z.logger
	} else {
		return esl.Default()
	}
}

func (z Opts) ReportMissingSource(base PathPair, source es_filesystem.Entry) {
	if z.handlerMissingSource != nil {
		z.handlerMissingSource(base, source)
	}
}

func (z Opts) ReportMissingTarget(base PathPair, target es_filesystem.Entry) {
	if z.handlerMissingTarget != nil {
		z.handlerMissingTarget(base, target)
	}
}

func (z Opts) ReportSameFile(base PathPair, source, target es_filesystem.Entry) {
	if z.handlerSameFile != nil {
		z.handlerSameFile(base, source, target)
	}
}

func (z Opts) ReportTypeDiff(base PathPair, source, target es_filesystem.Entry) {
	if z.handlerTypeDiff != nil {
		z.handlerTypeDiff(base, source, target)
	}
}

func (z Opts) ReportFileDiff(base PathPair, source, target es_filesystem.Entry) {
	if z.handlerFileDiff != nil {
		z.handlerFileDiff(base, source, target)
	}
}

func (z Opts) Apply(opts []Opt) Opts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:])
	}
}

func DontCompareTime(enabled bool) Opt {
	return func(o Opts) Opts {
		o.dontCompareTime = enabled
		return o
	}
}

func DontCompareContent(enabled bool) Opt {
	return func(o Opts) Opts {
		o.dontCompareContent = enabled
		return o
	}
}

func HandlerMissingSource(h MissingSource) Opt {
	return func(o Opts) Opts {
		o.handlerMissingSource = h
		return o
	}
}

func HandlerMissingTarget(h MissingTarget) Opt {
	return func(o Opts) Opts {
		o.handlerMissingTarget = h
		return o
	}
}

func HandlerFileDiff(h FileDiff) Opt {
	return func(o Opts) Opts {
		o.handlerFileDiff = h
		return o
	}
}

func HandlerTypeDiff(h TypeDiff) Opt {
	return func(o Opts) Opts {
		o.handlerTypeDiff = h
		return o
	}
}

func HandlerSameFile(h SameFile) Opt {
	return func(o Opts) Opts {
		o.handlerSameFile = h
		return o
	}
}
