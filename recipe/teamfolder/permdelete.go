package teamfolder

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_teamfolder"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_teamfolder"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/ingredient/teamfolder"
	"strings"
)

type Permdelete struct {
	rc_recipe.RemarkIrreversible
	Peer                           dbx_conn.ConnScopedTeam
	Name                           string
	ErrorUnableToResolveTeamfolder app_msg.Message
	ErrorUnableToRemoveTeamfolder  app_msg.Message
	ErrorTeamSpaceNotSupported     app_msg.Message
	SuccessPermdeleted             app_msg.Message
}

func (z *Permdelete) Preset() {
	z.Peer.SetScopes(
	//dbx_auth.ScopeTeamDataTeamSpace,
	//dbx_auth.ScopeTeamInfoRead,
	)
}

func (z *Permdelete) Exec(c app_control.Control) error {
	if ok, _ := teamfolder.IsTeamSpaceSupported(z.Peer.Context()); ok {
		c.UI().Error(z.ErrorTeamSpaceNotSupported)
		return errors.New("team space is not supported by this command")
	}

	ui := c.UI()

	teamfolders, err := sv_teamfolder.New(z.Peer.Context()).List()
	if err != nil {
		ui.Error(z.ErrorUnableToResolveTeamfolder.With("Error", err))
		return err
	}
	var tf *mo_teamfolder.TeamFolder
	for _, f := range teamfolders {
		if strings.ToLower(z.Name) == strings.ToLower(f.Name) {
			tf = f
			break
		}
	}
	if tf == nil {
		ui.Error(z.ErrorUnableToResolveTeamfolder.With("Error", err))
		return errors.New("unable to find team folder")
	}

	c.Log().Debug("Archiving team folder", esl.Any("teamfolder", tf))

	err = sv_teamfolder.New(z.Peer.Context()).PermDelete(tf)
	if err != nil {
		ui.Error(z.ErrorUnableToRemoveTeamfolder.With("Error", err))
		return err
	}
	ui.Success(z.SuccessPermdeleted.With("TeamFolderName", tf.Name))
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
