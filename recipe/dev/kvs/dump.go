package kvs

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
)

type DumpVO struct {
	Path string
}

type Dump struct {
}

func (z *Dump) Hidden() {
	panic("implement me")
}

func (z *Dump) Exec(k app_kitchen.Kitchen) error {
	panic("implement me")
}

func (z *Dump) Test(c app_control.Control) error {
	panic("implement me")
}

func (z *Dump) Requirement() app_vo.ValueObject {
	panic("implement me")
}

func (z *Dump) Reports() []rp_spec.ReportSpec {
	panic("implement me")
}
