package app_report

import (
	"github.com/watermint/toolbox/app86/app_msg"
)

type Report interface {
	Write(row interface{}, opts ...WriteOpt)
	Result(kind ResultKind, in interface{}, result interface{})
}

type WriteOpt func()

type ResultKind func() Result

type Result struct {
	Kind   string
	Reason app_msg.Message
}

func Success() ResultKind {
	return func() Result {
		return Result{
			Kind: "success",
		}
	}
}
func Failure(err error) ResultKind {
	return func() Result {
		return Result{
			Kind: "failure",
		}
	}
}
func Skip(reason string) ResultKind {
	return func() Result {
		return Result{
			Kind: "skip",
		}
	}
}
