package user

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_team"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_user"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Filesystem struct {
	Peer       dbx_conn.ConnScopedIndividual
	FileSystem rp_model.RowReport
}

func (z *Filesystem) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeAccountInfoRead,
	)
	z.FileSystem.SetModel(&mo_team.FileSystemVersion{})
}

func (z *Filesystem) Exec(c app_control.Control) error {
	if err := z.FileSystem.Open(); err != nil {
		return err
	}
	profile, err := sv_profile.NewProfile(z.Peer.Client()).Current()
	if err != nil {
		return err
	}
	if profile.TeamMemberId == "" {
		z.FileSystem.Row(mo_team.TeamFileSystemIndividual.Version())
		return nil
	}

	feature, err := sv_user.New(z.Peer.Client()).Features()
	if err != nil {
		return err
	}
	z.FileSystem.Row(feature.FileSystemType().Version())
	return nil
}

func (z *Filesystem) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Filesystem{}, rc_recipe.NoCustomValues)
}
