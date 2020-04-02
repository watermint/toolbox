package recipe

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"runtime"
)

type VersionInfo struct {
	Component string `json:"component"`
	Version   string `json:"version"`
}

type Version struct {
	Versions        rp_model.RowReport
	HeaderAppHash   app_msg.Message
	HeaderBuildTime app_msg.Message
	HeaderGoVersion app_msg.Message
}

func (z *Version) Preset() {
	z.Versions.SetModel(&VersionInfo{})
}

func (z *Version) Exec(c app_control.Control) error {
	ui := c.UI()
	if err := z.Versions.Open(); err != nil {
		return err
	}
	z.Versions.Row(&VersionInfo{
		Component: app.Name,
		Version:   app.Version,
	})
	z.Versions.Row(&VersionInfo{
		Component: ui.Text(z.HeaderAppHash),
		Version:   app.Hash,
	})
	z.Versions.Row(&VersionInfo{
		Component: ui.Text(z.HeaderBuildTime),
		Version:   app.BuildTimestamp,
	})
	z.Versions.Row(&VersionInfo{
		Component: ui.Text(z.HeaderGoVersion),
		Version:   runtime.Version(),
	})
	return nil
}

func (z *Version) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Version{}, rc_recipe.NoCustomValues)
}
