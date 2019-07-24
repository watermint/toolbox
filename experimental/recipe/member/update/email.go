package update

import (
	"github.com/watermint/toolbox/experimental/app_file"
	"github.com/watermint/toolbox/experimental/app_kitchen"
	"github.com/watermint/toolbox/experimental/app_vo"
)

type EmailVO struct {
	File                 app_file.RowDataFile
	DontUpdateUnverified bool
}

type Email struct {
}

func (z *Email) Requirement() app_vo.ValueObject {
	panic("implement me")
}

func (z *Email) Exec(k app_kitchen.Kitchen) error {
	panic("implement me")
}
