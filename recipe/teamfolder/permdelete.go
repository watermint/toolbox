package teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"strings"
)

type PermDeleteVO struct {
	Peer rc_conn.ConnBusinessFile
	Name string
}

type PermDelete struct {
}

func (z *PermDelete) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *PermDelete) Console() {
}

func (z *PermDelete) Requirement() rc_vo.ValueObject {
	return &PermDeleteVO{}
}

func (z *PermDelete) Exec(k rc_kitchen.Kitchen) error {
	ui := k.UI()
	vo := k.Value().(*PermDeleteVO)

	if vo.Name == "" {
		ui.Error("recipe.teamfolder.permdelete.err.missing_option.name")
		return errors.New("missing required option")
	}

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	teamfolders, err := sv_teamfolder.New(ctx).List()
	if err != nil {
		ui.Error("recipe.teamfolder.permdelete.err.unable_to_resolve_teamfolder",
			app_msg.P{
				"Error": err.Error(),
			})
		return err
	}
	var teamfolder *mo_teamfolder.TeamFolder
	for _, tf := range teamfolders {
		if strings.ToLower(vo.Name) == strings.ToLower(tf.Name) {
			teamfolder = tf
			break
		}
	}
	if teamfolder == nil {
		ui.Error("recipe.teamfolder.permdelete.err.unable_to_resolve_teamfolder",
			app_msg.P{
				"Error": "Unable to find team folder",
			})
		return errors.New("unable to find team folder")
	}

	k.Log().Debug("Archiving team folder", zap.Any("teamfolder", teamfolder))

	err = sv_teamfolder.New(ctx).PermDelete(teamfolder)
	if err != nil {
		ui.Error("recipe.teamfolder.permdelete.err.unable_to_remove_teamfolder", app_msg.P{
			"Error": err.Error(),
		})
		return err
	}
	ui.Success("recipe.teamfolder.permdelete.success.permdeleted", app_msg.P{
		"TeamFolderName": teamfolder.Name,
	})
	return nil
}

func (z *PermDelete) Test(c app_control.Control) error {
	vo := &PermDeleteVO{}
	if !qt_recipe.ApplyTestPeers(c, vo) {
		return qt_recipe.HumanInteractionRequired()
	}

	// should fail
	{
		vo.Name = ""
		if err := z.Exec(rc_kitchen.NewKitchen(c, vo)); err == nil {
			return errors.New("empty name should fail")
		}
	}
	{
		vo.Name = "No existent"
		if err := z.Exec(rc_kitchen.NewKitchen(c, vo)); err == nil {
			return errors.New("non exist team folder name should fail")
		}
	}
	return nil
}
