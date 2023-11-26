package es_native_windows

import (
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated"
	"github.com/watermint/toolbox/essentials/go/es_idiom_deprecated/eoutcome"
	"syscall"
)

// KernelOutcome consumer need to verify procedure's output to determine weather the process failed or not.
// KernelOutcome marks as error only when the procedure is not found or could not bind to kernel module, etc.
type KernelOutcome interface {
	es_idiom_deprecated.UnconfirmedOutcome

	IsCouldNotResolveProc() bool

	Errno() syscall.Errno

	// LastError Returns the proc's last error
	LastError() error
}

const (
	kernelOutcomeNoObviousError = iota
	kernelOutcomeCouldNotResolveProc
)

func NewKernelOutcomeNoObviousError(lastErr error) KernelOutcome {
	return &kernelOutcomeImpl{
		UnconfirmedOutcomeBase: eoutcome.UnconfirmedOutcomeBase{ObviousErr: nil},
		cause:                  kernelOutcomeNoObviousError,
		lastErr:                lastErr,
	}
}

func NewKernelOutcomeCouldNotResolveProc(resolveError error) KernelOutcome {
	return &kernelOutcomeImpl{
		UnconfirmedOutcomeBase: eoutcome.UnconfirmedOutcomeBase{ObviousErr: resolveError},
		cause:                  kernelOutcomeCouldNotResolveProc,
		lastErr:                resolveError,
	}
}

type kernelOutcomeImpl struct {
	eoutcome.UnconfirmedOutcomeBase
	cause   int
	lastErr error
}

func (z kernelOutcomeImpl) Errno() syscall.Errno {
	if z.lastErr == nil {
		return 0
	}
	if en, ok := z.lastErr.(syscall.Errno); ok {
		return en
	}
	return 0
}

func (z kernelOutcomeImpl) IsCouldNotResolveProc() bool {
	return z.cause == kernelOutcomeCouldNotResolveProc
}

func (z kernelOutcomeImpl) LastError() error {
	return z.lastErr
}
