package update

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Type struct {
	Peer  dbx_conn.ConnScopedTeam
	Name  string
	Type  mo_string.SelectString
	Group rp_model.RowReport
}

func (z *Type) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsRead,
		dbx_auth.ScopeGroupsWrite,
	)
	z.Type.SetOptions("company_managed", "user_managed", "company_managed")
	z.Group.SetModel(&mo_group.Group{})
}

func (z *Type) Exec(c app_control.Control) error {
	if err := z.Group.Open(); err != nil {
		return err
	}
	g, err := sv_group.New(z.Peer.Client()).ResolveByName(z.Name)
	if err != nil {
		return err
	}
	update, err := sv_group.New(z.Peer.Client()).Update(&mo_group.Group{
		GroupId:             g.GroupId,
		GroupManagementType: z.Type.Value(),
	})
	if err != nil {
		return err
	}
	z.Group.Row(update)
	return nil
}

func (z *Type) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Type{}, func(r rc_recipe.Recipe) {
		m := r.(*Type)
		m.Name = "Sales"
		m.Type.SetSelect("user_managed")
	})
}
