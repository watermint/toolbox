package account

import (
	"github.com/watermint/toolbox/domain/hellosign/api/hs_conn"
	"github.com/watermint/toolbox/domain/hellosign/model/mo_account"
	"github.com/watermint/toolbox/domain/hellosign/service/sv_account"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Info struct {
	rc_recipe.RemarkSecret
	Peer      hs_conn.ConnHelloSignApi
	AccountId mo_string.OptionalString
	Account   rp_model.RowReport
}

func (z *Info) Preset() {
	z.Account.SetModel(
		&mo_account.Account{},
		rp_model.HiddenColumns(
			"account_id",
		),
	)
}

func (z *Info) Exec(c app_control.Control) error {
	if err := z.Account.Open(); err != nil {
		return err
	}
	info, err := sv_account.New(z.Peer.Client()).Info(z.AccountId.Value())
	if err != nil {
		return err
	}
	z.Account.Row(&info)
	return nil
}

func (z *Info) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Info{}, rc_recipe.NoCustomValues)
}
