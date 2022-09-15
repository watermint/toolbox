package auth

import (
	"errors"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Delete struct {
	KeyName  string
	PeerName string
	Deleted  rp_model.RowReport
}

func (z *Delete) Preset() {
	z.Deleted.SetModel(&api_auth.EntityNoCredential{})
}

func (z *Delete) Exec(c app_control.Control) error {
	if err := z.Deleted.Open(); err != nil {
		return err
	}

	rt, ok := c.AuthRepository().(api_auth.RepositoryTraversable)
	if !ok {
		return errors.New("no traversable repository found")
	}

	for _, e := range rt.All() {
		if e.KeyName == z.KeyName && e.PeerName == z.PeerName {
			c.AuthRepository().Delete(e.KeyName, e.Scope, e.PeerName)
			z.Deleted.Row(e.NoCredential())
		}
	}

	return nil
}

func (z *Delete) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Delete{}, func(r rc_recipe.Recipe) {
		m := r.(*Delete)
		m.KeyName = "no_existent"
		m.PeerName = "no_existent"
	})
}
