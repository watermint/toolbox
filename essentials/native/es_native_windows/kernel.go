//go:build windows
// +build windows

package es_native_windows

import "syscall"

type Kernel interface {
	Call(procName string, args ...uintptr) (r1, r2 uintptr, err error)
}

var (
	Kernel32 Kernel = &kernelWrapper{}
)

type kernelWrapper struct {
}

func (z kernelWrapper) Call(procName string, args ...uintptr) (r1, r2 uintptr, err error) {
	k32, resolveErr := syscall.LoadDLL("kernel32")
	if resolveErr != nil {
		return 0, 0, resolveErr
	}
	proc, resolveErr := k32.FindProc(procName)
	if resolveErr != nil {
		return 0, 0, resolveErr
	}

	r1, r2, lastErr := proc.Call(args...)
	if lastErr != nil {
		return r1, r2, lastErr
	}
	return r1, r2, nil
}
