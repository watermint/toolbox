package auth

import (
	"github.com/watermint/toolbox/domain/github/api/gh_conn"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Connect struct {
	rc_recipe.RemarkSecret
	Github gh_conn.ConnGithubRepo
}

func (z *Connect) Preset() {
	z.Github.SetPeerName(app.PeerDeploy)
}

func (z *Connect) Exec(c app_control.Control) error {
	return nil
}

func (z *Connect) Test(c app_control.Control) error {
	return qt_errors.ErrorNoTestRequired
}
