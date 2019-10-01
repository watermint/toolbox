package group

import (
	"errors"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recpie/app_conn"
	"github.com/watermint/toolbox/infra/recpie/app_kitchen"
	"github.com/watermint/toolbox/infra/recpie/app_test"
	"github.com/watermint/toolbox/infra/recpie/app_vo"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
)

type RemoveVO struct {
	Peer app_conn.ConnBusinessMgmt
	Name string
}

type Remove struct {
}

func (z *Remove) Requirement() app_vo.ValueObject {
	return &RemoveVO{}
}

func (z *Remove) Exec(k app_kitchen.Kitchen) error {
	ui := k.UI()
	vo := k.Value().(*RemoveVO)

	if vo.Name == "" {
		ui.Error("recipe.group.remove.err.missing_option.name")
		return errors.New("missing required option")
	}

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	group, err := sv_group.New(ctx).ResolveByName(vo.Name)
	if err != nil {
		ui.Error("recipe.group.remove.err.unable_to_resolve_group",
			app_msg.P{
				"Error": err.Error(),
			})
		return err
	}
	k.Log().Debug("Removing group", zap.Any("group", group))

	err = sv_group.New(ctx).Remove(group.GroupId)
	if err != nil {
		ui.Error("recipe.group.remove.err.unable_to_remove_group", app_msg.P{
			"Error": err.Error(),
		})
		return err
	}
	ui.Success("recipe.group.remove.success.removed", app_msg.P{
		"GroupName":      group.GroupName,
		"ManagementType": group.GroupManagementType,
	})
	return nil
}

func (z *Remove) Test(c app_control.Control) error {
	vo := &RemoveVO{}
	if !app_test.ApplyTestPeers(c, vo) {
		return nil
	}

	// should fail
	{
		vo.Name = ""
		if err := z.Exec(app_kitchen.NewKitchen(c, vo)); err == nil {
			return errors.New("empty name should fail")
		}
	}
	{
		vo.Name = "No existent"
		if err := z.Exec(app_kitchen.NewKitchen(c, vo)); err == nil {
			return errors.New("non exist group name should fail")
		}
	}
	return nil
}
