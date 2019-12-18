package file

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/rc_kitchen"
	"github.com/watermint/toolbox/infra/recpie/rc_vo"
	"github.com/watermint/toolbox/recipe/team/namespace/file"
)

type List struct {
}

func (z *List) Requirement() rc_vo.ValueObject {
	return &file.ListVO{
		IncludeTeamFolder:   true,
		IncludeSharedFolder: false,
	}
}

func (z *List) Exec(k rc_kitchen.Kitchen) error {
	fl := file.List{}
	return fl.Exec(k)
}

func (z *List) Test(c app_control.Control) error {
	fl := file.List{}
	return fl.Test(c)
}
