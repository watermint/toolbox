package recipe

import (
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type License struct {
}

func (*License) Console() {
}

func (*License) Requirement() app_vo.ValueObject {
	return &app_vo.EmptyValueObject{}
}

func (*License) Exec(k app_kitchen.Kitchen) error {
	return nil
}
