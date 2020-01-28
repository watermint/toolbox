package teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/service/sv_teamfolder"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"go.uber.org/zap"
	"strings"
)

type Permdelete struct {
	Peer rc_conn.ConnBusinessFile
	Name string
}

func (z *Permdelete) Preset() {
}

func (z *Permdelete) Exec(c app_control.Control) error {
	ui := c.UI()

	if z.Name == "" {
		ui.ErrorK("recipe.teamfolder.permdelete.err.missing_option.name")
		return errors.New("missing required option")
	}

	teamfolders, err := sv_teamfolder.New(z.Peer.Context()).List()
	if err != nil {
		ui.ErrorK("recipe.teamfolder.permdelete.err.unable_to_resolve_teamfolder",
			app_msg.P{
				"Error": err.Error(),
			})
		return err
	}
	var teamfolder *mo_teamfolder.TeamFolder
	for _, tf := range teamfolders {
		if strings.ToLower(z.Name) == strings.ToLower(tf.Name) {
			teamfolder = tf
			break
		}
	}
	if teamfolder == nil {
		ui.ErrorK("recipe.teamfolder.permdelete.err.unable_to_resolve_teamfolder",
			app_msg.P{
				"Error": "Unable to find team folder",
			})
		return errors.New("unable to find team folder")
	}

	c.Log().Debug("Archiving team folder", zap.Any("teamfolder", teamfolder))

	err = sv_teamfolder.New(z.Peer.Context()).PermDelete(teamfolder)
	if err != nil {
		ui.ErrorK("recipe.teamfolder.permdelete.err.unable_to_remove_teamfolder", app_msg.P{
			"Error": err.Error(),
		})
		return err
	}
	ui.Success("recipe.teamfolder.permdelete.success.permdeleted", app_msg.P{
		"TeamFolderName": teamfolder.Name,
	})
	return nil
}

func (z *Permdelete) Test(c app_control.Control) error {
	// should fail
	{
		err := rc_exec.Exec(c, &Permdelete{}, func(r rc_recipe.Recipe) {
			rc := r.(*Permdelete)
			rc.Name = ""
		})
		if err == nil {
			return errors.New("empty name should fail")
		}
	}
	{
		err := rc_exec.Exec(c, &Permdelete{}, func(r rc_recipe.Recipe) {
			rc := r.(*Permdelete)
			rc.Name = "No existent"
		})
		if err == nil {
			return errors.New("non exist team folder name should fail")
		}
	}
	return nil
}
