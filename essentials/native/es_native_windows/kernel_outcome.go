package es_native_windows

import (
	"errors"
	"syscall"
)

// KernelOutcome は特定のエラー状態を表します
type KernelError struct {
	error
	cause   int
	lastErr error
}

// エラーのタイプを識別するための定数
const (
	kernelOutcomeNoObviousError = iota
	kernelOutcomeCouldNotResolveProc
)

// エラー識別用の関数
func IsCouldNotResolveProcError(err error) bool {
	var ke *KernelError
	if errors.As(err, &ke) {
		return ke.cause == kernelOutcomeCouldNotResolveProc
	}
	return false
}

// エラーのErrnoを取得する関数
func GetErrno(err error) syscall.Errno {
	var ke *KernelError
	if errors.As(err, &ke) {
		if ke.lastErr == nil {
			return 0
		}
		if en, ok := ke.lastErr.(syscall.Errno); ok {
			return en
		}
	}
	return 0
}

// エラーの元になったLastErrorを取得する関数
func GetLastError(err error) error {
	var ke *KernelError
	if errors.As(err, &ke) {
		return ke.lastErr
	}
	return nil
}

// エラーメッセージの実装
func (z *KernelError) Error() string {
	if z.error != nil {
		return z.error.Error()
	}
	return "kernel operation error"
}

// Unwrap()メソッドの実装
func (z *KernelError) Unwrap() error {
	return z.error
}

// エラーなしの状態を表すファクトリー関数
func NewKernelOutcomeNoObviousError(lastErr error) error {
	if lastErr == nil {
		return nil
	}
	return &KernelError{
		error:   lastErr,
		cause:   kernelOutcomeNoObviousError,
		lastErr: lastErr,
	}
}

// プロシージャの解決失敗を表すファクトリー関数
func NewKernelOutcomeCouldNotResolveProc(resolveError error) error {
	return &KernelError{
		error:   resolveError,
		cause:   kernelOutcomeCouldNotResolveProc,
		lastErr: resolveError,
	}
}
