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
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"go.uber.org/zap"
	"strings"
)

type ArchiveVO struct {
	Peer rc_conn.OldConnBusinessFile
	Name string
}

type Archive struct {
}

func (z *Archive) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Archive) Console() {
}

func (z *Archive) Requirement() rc_vo.ValueObject {
	return &ArchiveVO{}
}

func (z *Archive) Exec(k rc_kitchen.Kitchen) error {
	ui := k.UI()
	vo := k.Value().(*ArchiveVO)

	if vo.Name == "" {
		ui.Error("recipe.teamfolder.archive.err.missing_option.name")
		return errors.New("missing required option")
	}

	ctx, err := vo.Peer.Connect(k.Control())
	if err != nil {
		return err
	}

	teamfolders, err := sv_teamfolder.New(ctx).List()
	if err != nil {
		ui.Error("recipe.teamfolder.archive.err.unable_to_resolve_teamfolder",
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
		ui.Error("recipe.teamfolder.archive.err.unable_to_resolve_teamfolder",
			app_msg.P{
				"Error": "Unable to find team folder",
			})
		return errors.New("unable to find team folder")
	}

	k.Log().Debug("Archiving team folder", zap.Any("teamfolder", teamfolder))

	_, err = sv_teamfolder.New(ctx).Archive(teamfolder)
	if err != nil {
		ui.Error("recipe.teamfolder.archive.err.unable_to_remove_teamfolder", app_msg.P{
			"Error": err.Error(),
		})
		return err
	}
	ui.Success("recipe.teamfolder.archive.success.archived", app_msg.P{
		"TeamFolderName": teamfolder.Name,
	})
	return nil
}

func (z *Archive) Test(c app_control.Control) error {
	vo := &ArchiveVO{}
	if !qt_recipe.ApplyTestPeers(c, vo) {
		return qt_endtoend.NotEnoughResource()
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
	return qt_endtoend.HumanInteractionRequired()
}
