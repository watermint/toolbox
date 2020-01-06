package namespace

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_namespace"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
)

type List struct {
	Peer      rc_conn.ConnBusinessFile
	Namespace rp_model.RowReport
}

func (z *List) Preset() {
	z.Namespace.SetModel(&mo_namespace.Namespace{})
}

func (z *List) Exec(c app_control.Control) error {
	if err := z.Namespace.Open(); err != nil {
		return err
	}

	namespaces, err := sv_namespace.New(z.Peer.Context()).List()
	if err != nil {
		return err
	}

	for _, namespace := range namespaces {
		z.Namespace.Row(namespace)
	}
	return nil
}

func (z *List) Test(c app_control.Control) error {
	if err := rc_exec.Exec(c, &List{}, rc_recipe.NoCustomValues); err != nil {
		return err
	}
	return qt_recipe.TestRows(c, "namespace", func(cols map[string]string) error {
		if _, ok := cols["namespace_id"]; !ok {
			return errors.New("`namespace_id` is not found")
		}
		return nil
	})
}
