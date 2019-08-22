package app_vo

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type ValueObject interface {
	Validate(t Validator)
}

type EmptyValueObject struct {
}

func (*EmptyValueObject) Validate(t Validator) {
}

type Validator interface {
	Invalid(key string, placeHolders ...app_msg.Param)
	AssertFileExists(path string)
	AssertEmailFormat(email string)
}
