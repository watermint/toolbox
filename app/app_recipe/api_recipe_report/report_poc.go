package api_recipe_report

import "github.com/watermint/toolbox/app/app_recipe/api_recipe_msg"

type Report interface {
	Write(row interface{}, opts ...WriteOpt)
	Result(kind ResultKind, in interface{}, result interface{})
}

type WriteOpt func()

type ResultKind func() Result

type Result struct {
	Kind   string
	Reason api_recipe_msg.Message
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
