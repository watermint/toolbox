package group

import (
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"strings"
)

type Add struct {
	Peer                       rc_conn.ConnBusinessMgmt
	Name                       string
	ManagementType             string
	ErrorInvalidManagementType app_msg.Message
	ErrorUnableToAddGroup      app_msg.Message
	AddedGroup                 rp_model.RowReport
}

func (z *Add) Exec(c app_control.Control) error {
	opts := make([]sv_group.CreateOpt, 0)
	switch strings.ToLower(z.ManagementType) {
	case "company_managed":
		opts = append(opts, sv_group.CompanyManaged())
	case "user_managed":
		opts = append(opts, sv_group.UserManaged())
	default:
		c.UI().Error(z.ErrorInvalidManagementType.With("Type", z.ManagementType))
	}
	if err := z.AddedGroup.Open(); err != nil {
		return err
	}

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
		m.ManagementType = "company_managed"
	})
	if err, _ = qt_recipe.RecipeError(c.Log(), err); err != nil {
		return err
	}
	err = rc_exec.ExecMock(c, &Add{}, func(r rc_recipe.Recipe) {
		m := r.(*Add)
		m.Name = "Marketing"
		m.ManagementType = "user_managed"
	})
	if err, _ = qt_recipe.RecipeError(c.Log(), err); err != nil {
		return err
	}

	return qt_errors.ErrorScenarioTest
}

func (z *Add) Preset() {
	z.ManagementType = "company_managed"
	z.AddedGroup.SetModel(&mo_group.Group{})
}
