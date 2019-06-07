package update

import (
	"github.com/watermint/toolbox/experimental/app_flow"
	"github.com/watermint/toolbox/experimental/app_recipe"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type EmailVO struct {
	File                 app_flow.RowDataFile
	DontUpdateUnverified bool
}

type Email struct {
}

func (z *Email) Requirement() app_vo.ValueObject {
	panic("implement me")
}

func (z *Email) Exec(k app_recipe.Kitchen) error {
	panic("implement me")
}
