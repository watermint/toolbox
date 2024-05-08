package request

import (
	"github.com/watermint/toolbox/domain/dropboxsign/api/hs_conn"
	"github.com/watermint/toolbox/domain/dropboxsign/model/mo_signature"
	"github.com/watermint/toolbox/domain/dropboxsign/service/sv_signature"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Peer      hs_conn.ConnHelloSignApi
	AccountId mo_string.OptionalString
	Requests  rp_model.RowReport
}

func (z *List) Preset() {
	z.Requests.SetModel(&mo_signature.Request{},
		rp_model.HiddenColumns(
			"signatures",
			"created_at",
			"expires_at",
		),
	)
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Requests.Open(); err != nil {
		return err
	}
	return sv_signature.New(z.Peer.Client()).List(z.AccountId.Value(), func(req *mo_signature.Request) bool {
		z.Requests.Row(req)
		return true
	})
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &List{}, func(r rc_recipe.Recipe) {
		m := r.(*List)
		m.AccountId = mo_string.NewOptional("all")
	})
}
