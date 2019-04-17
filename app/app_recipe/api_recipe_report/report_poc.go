package api_recipe_report

import "github.com/watermint/toolbox/app/app_recipe/api_recipe_msg"

type Report interface {
	Write(row interface{}, opts ...WriteOpt)
	Result(kind ResultKind, in interface{}, result interface{}, reasons ...Reason)
}

type WriteOpt func()

type ResultKind int
type Reason api_recipe_msg.Message

func DueToError(err error) Reason {
	return nil
}

const (
	Success = iota
	Failure
	FailurePartially
	Skip
)
