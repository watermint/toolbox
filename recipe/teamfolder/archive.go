package teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
)

type Archive struct {
	Peer                           dbx_conn.ConnBusinessFile
	Name                           string
	ErrorUnableToResolveTeamfolder app_msg.Message
	ErrorUnableToArchiveTeamfolder app_msg.Message
	SuccessArchived                app_msg.Message
}

func (z *Archive) Preset() {
}

func (z *Archive) Exec(c app_control.Control) error {
	ui := c.UI()

	teamfolders, err := sv_teamfolder.New(z.Peer.Context()).List()
	if err != nil {
		ui.Error(z.ErrorUnableToResolveTeamfolder.With("Error", err))
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
		ui.Error(z.ErrorUnableToResolveTeamfolder.With("Error", err))
		return errors.New("unable to find team folder")
	}

	c.Log().Debug("Archiving team folder", es_log.Any("teamfolder", teamfolder))

	_, err = sv_teamfolder.New(z.Peer.Context()).Archive(teamfolder)
	if err != nil {
		ui.Error(z.ErrorUnableToArchiveTeamfolder.With("Error", err))
		return err
	}
	ui.Success(z.SuccessArchived.With("TeamFolderName", teamfolder.Name))
	return nil
}

func (z *Archive) Test(c app_control.Control) error {
	// should fail
	{
		err := rc_exec.Exec(c, &Archive{}, func(r rc_recipe.Recipe) {
			rc := r.(*Archive)
			rc.Name = ""
		})
		if err == nil {
			return errors.New("empty name should fail")
		}
	}
	{
		err := rc_exec.Exec(c, &Archive{}, func(r rc_recipe.Recipe) {
			rc := r.(*Archive)
			rc.Name = "No existent"
		})
		if err == nil {
			return errors.New("non exist team folder name should fail")
		}
	}
	return nil
}
