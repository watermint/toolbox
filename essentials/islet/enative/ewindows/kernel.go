//go:build windows
// +build windows

package ewindows

import "syscall"

type Kernel interface {
	Call(procName string, args ...uintptr) (r1, r2 uintptr, outcome KernelOutcome)
}

var (
	Kernel32 Kernel = &kernelWrapper{}
)

type kernelWrapper struct {
}

func (z kernelWrapper) Call(procName string, args ...uintptr) (r1, r2 uintptr, outcome KernelOutcome) {
	k32, resolveErr := syscall.LoadDLL("kernel32")
	if resolveErr != nil {
		return 0, 0, NewKernelOutcomeCouldNotResolveProc(resolveErr)
	}
	proc, resolveErr := k32.FindProc(procName)
	if resolveErr != nil {
		return 0, 0, NewKernelOutcomeCouldNotResolveProc(resolveErr)
	}

	r1, r2, lastErr := proc.Call(args...)
	return r1, r2, NewKernelOutcomeNoObviousError(lastErr)
}
