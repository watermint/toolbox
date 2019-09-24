package filerequest

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_file"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
)

type CloneVO struct {
	File     app_file.ColDataFile
	PeerName app_conn.ConnBusinessFile
}

type Clone struct {
}

func (z *Clone) Requirement() app_vo.ValueObject {
	return &CloneVO{}
}

func (z *Clone) Exec(k app_kitchen.Kitchen) error {
	panic("implement me")
}

func (z *Clone) Test(c app_control.Control) error {
	panic("implement me")
}
