package member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_member"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type List struct {
	Peer   dbx_conn.ConnBusinessInfo
	Member rp_model.RowReport
}

func (z *List) Preset() {
	z.Member.SetModel(
		&mo_member.Member{},
		rp_model.HiddenColumns(
			"tag",
			"team_member_id",
			"member_folder_id",
			"account_id",
			"persistent_id",
			"familiar_name",
			"abbreviated_name",
			"external_id",
		),
	)
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "member", func(cols map[string]string) error {
		if _, ok := cols["email"]; !ok {
			return errors.New("email is not found")
		}
		return nil
	})
}

func (z *List) Exec(c app_control.Control) error {
	members, err := sv_member.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}
	if err := z.Member.Open(); err != nil {
		return err
	}
	for _, m := range members {
		z.Member.Row(m)
	}
	return nil
}
