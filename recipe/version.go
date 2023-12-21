package recipe

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	app_definitions2 "github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/control/app_lifecycle"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/resources"
	"runtime"
	"strconv"
	"time"
)

type VersionInfo struct {
	Key       string `json:"key"`
	Component string `json:"component"`
	Version   string `json:"version"`
}

type Version struct {
	rc_recipe.RemarkTransient
	Versions         rp_model.RowReport
	HeaderAppHash    app_msg.Message
	HeaderBuildTime  app_msg.Message
	HeaderBranch     app_msg.Message
	HeaderProduction app_msg.Message
	HeaderGoVersion  app_msg.Message
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
		Key:       "app.name",
		Component: "app.name",
		Version:   app_definitions2.Name,
	})
	z.Versions.Row(&VersionInfo{
		Key:       "app.version",
		Component: app_definitions2.Name,
		Version:   app_definitions2.BuildId,
	})
	z.Versions.Row(&VersionInfo{
		Key:       "app.hash",
		Component: ui.Text(z.HeaderAppHash),
		Version:   app_definitions2.BuildInfo.Hash,
	})
	z.Versions.Row(&VersionInfo{
		Key:       "app.branch",
		Component: ui.Text(z.HeaderBranch),
		Version:   app_definitions2.BuildInfo.Branch,
	})
	z.Versions.Row(&VersionInfo{
		Key:       "app.production",
		Component: ui.Text(z.HeaderProduction),
		Version:   strconv.FormatBool(app_definitions2.BuildInfo.Production),
	})
	z.Versions.Row(&VersionInfo{
		Key:       "core.release",
		Component: app_definitions2.Pkg,
		Version:   resources.CoreRelease(),
	})
	z.Versions.Row(&VersionInfo{
		Key:       "build.time",
		Component: ui.Text(z.HeaderBuildTime),
		Version:   app_definitions2.BuildInfo.Timestamp,
	})
	z.Versions.Row(&VersionInfo{
		Key:       "go.version",
		Component: ui.Text(z.HeaderGoVersion),
		Version:   runtime.Version(),
	})
	z.Versions.Row(&VersionInfo{
		Key:       "lifecycle.bestbefore",
		Component: "lifecycle",
		Version:   app_lifecycle.LifecycleControl().TimeBestBefore().Format(time.RFC3339),
	})
	z.Versions.Row(&VersionInfo{
		Key:       "lifecycle.expiration",
		Component: "lifecycle",
		Version:   app_lifecycle.LifecycleControl().TimeExpiration().Format(time.RFC3339),
	})
	z.Versions.Row(&VersionInfo{
		Key:       "lifecycle.expiration_mode",
		Component: "lifecycle",
		Version:   string(app_definitions2.LifecycleExpirationMode),
	})
	return nil
}

func (z *Version) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Version{}, rc_recipe.NoCustomValues)
}
