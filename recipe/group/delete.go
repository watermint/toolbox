package group

import (
	"errors"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
)

type Delete struct {
	Peer rc_conn.ConnBusinessMgmt
	Name string
}

func (z *Delete) Preset() {
}

func (z *Delete) Exec(c app_control.Control) error {
	ui := c.UI()

	if z.Name == "" {
		ui.ErrorK("recipe.group.delete.err.missing_option.name")
		return errors.New("missing required option")
	}

	group, err := sv_group.New(z.Peer.Context()).ResolveByName(z.Name)
	if err != nil {
		ui.ErrorK("recipe.group.delete.err.unable_to_resolve_group",
			app_msg.P{
				"Error": err.Error(),
			})
		return err
	}
	c.Log().Debug("Removing group", zap.Any("group", group))

	err = sv_group.New(z.Peer.Context()).Remove(group.GroupId)
	if err != nil {
		ui.ErrorK("recipe.group.delete.err.unable_to_remove_group", app_msg.P{
			"Error": err.Error(),
		})
		return err
	}
	ui.SuccessK("recipe.group.delete.success.removed", app_msg.P{
		"GroupName":      group.GroupName,
		"ManagementType": group.GroupManagementType,
	})
	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	// should fail
	{
		err := rc_exec.Exec(c, &Delete{}, func(r rc_recipe.Recipe) {
			rc := r.(*Delete)
			rc.Name = ""
		})
		if err == nil {
			return errors.New("empty name should fail")
		}
	}
	{
		err := rc_exec.Exec(c, &Delete{}, func(r rc_recipe.Recipe) {
			rc := r.(*Delete)
			rc.Name = "No existent"
		})
		if err == nil {
			return errors.New("non exist group name should fail")
		}
	}
	return nil
}
