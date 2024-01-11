package recipe

import (
	"github.com/watermint/toolbox/domain/core/dc_version"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
)

type Version struct {
	rc_recipe.RemarkTransient
	Versions rp_model.RowReport
}

func (z *Version) Preset() {
	z.Versions.SetModel(&dc_version.VersionInfo{})
}

func (z *Version) Exec(c app_control.Control) error {
	if err := z.Versions.Open(); err != nil {
		return err
	}

	components := dc_version.VersionComponents(c.UI())
	for _, component := range components {
		z.Versions.Row(component)
	}
	return nil
}

func (z *Version) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Version{}, rc_recipe.NoCustomValues)
}
