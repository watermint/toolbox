package group

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
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Add struct {
	rc_recipe.RemarkIrreversible
	Peer                  dbx_conn.ConnScopedTeam
	Name                  string
	ManagementType        mo_string.SelectString
	ErrorUnableToAddGroup app_msg.Message
	AddedGroup            rp_model.RowReport
}

func (z *Add) Exec(c app_control.Control) error {
	if err := z.AddedGroup.Open(); err != nil {
		return err
	}
	opts := make([]sv_group.CreateOpt, 0)
	opts = append(opts, sv_group.ManagementType(z.ManagementType.Value()))

	group, err := sv_group.New(z.Peer.Context()).Create(z.Name, opts...)
	if err != nil {
		c.UI().Error(z.ErrorUnableToAddGroup.With("Name", z.Name).With("Error", err))
		return err
	}
	z.AddedGroup.Row(group)
	return nil
}

func (z *Add) Test(c app_control.Control) error {
	err := rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Name = "Marketing"
		m.ManagementType.SetSelect("company_managed")
	})
	if err, _ = qt_errors.ErrorsForTest(c.Log(), err); err != nil {
		return err
	}
	err = rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Name = "Marketing"
		m.ManagementType.SetSelect("user_managed")
	})
	if err, _ = qt_errors.ErrorsForTest(c.Log(), err); err != nil {
		return err
	}

	return qt_errors.ErrorScenarioTest
}

func (z *Add) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeGroupsWrite,
	)
	z.ManagementType.SetOptions(
		"company_managed",
		"company_managed", "user_managed",
	)
	z.AddedGroup.SetModel(
		&mo_group.Group{},
		rp_model.HiddenColumns(
			"group_id",
			"group_external_id",
		),
	)
}
