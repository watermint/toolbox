package auth

import (
	"errors"
	"github.com/watermint/toolbox/essentials/api/api_auth"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type List struct {
	Entity rp_model.RowReport
}

func (z *List) Preset() {
	z.Entity.SetModel(&api_auth.EntityNoCredential{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Entity.Open(); err != nil {
		return err
	}

	rt, ok := c.AuthRepository().(api_auth.RepositoryTraversable)
	if !ok {
		return errors.New("no traversable repository found")
	}

	for _, e := range rt.All() {
		z.Entity.Row(e.NoCredential())
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues)
}
