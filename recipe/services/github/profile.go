package github

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/domain/github/model/mo_user"
	"github.com/watermint/toolbox/domain/github/service/sv_profile"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Profile struct {
	rc_recipe.RemarkExperimental
	Peer gh_conn.ConnGithubRepo
	User rp_model.RowReport
}

func (z *Profile) Preset() {
	z.User.SetModel(
		&mo_user.User{},
		rp_model.HiddenColumns(
			"id",
		),
	)
}

func (z *Profile) Exec(c app_control.Control) error {
	if err := z.User.Open(); err != nil {
		return err
	}
	user, err := sv_profile.New(z.Peer.Client()).User()
	if err != nil {
		return err
	}
	z.User.Row(user)
	return nil
}

func (z *Profile) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Profile{}, rc_recipe.NoCustomValues)
}
