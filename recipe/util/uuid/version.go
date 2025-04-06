package uuid

import (
	"strings"

	"github.com/watermint/toolbox/essentials/strings/es_uuid"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Version struct {
	Uuid     string
	Metadata rp_model.RowReport
}

func (z *Version) Preset() {
	z.Metadata.SetModel(&es_uuid.UUIDMetadata{})
}

func (z *Version) Exec(c app_control.Control) error {
	if err := z.Metadata.Open(); err != nil {
		return err
	}
	u, err := es_uuid.Parse(strings.TrimSpace(z.Uuid))
	if err != nil {
		return err
	}
	z.Metadata.Row(u.Metadata())
	return nil
}

func (z *Version) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Version{}, func(r rc_recipe.Recipe) {
		m := r.(*Version)
		m.Uuid = es_uuid.NewV7().String()
	})
}
