package config

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strconv"
)

type Features struct {
	SectionDescription   app_msg.Message
	SectionSettings      app_msg.Message
	HeaderKey            app_msg.Message
	HeaderDesc           app_msg.Message
	HeaderStatus         app_msg.Message
	HeaderOptInUser      app_msg.Message
	HeaderOptInTimestamp app_msg.Message
}

func (z *Features) Preset() {
}

func (z *Features) Exec(c app_control.Control) error {
	l := c.Log()
	ui := c.UI()

	cl, ok := c.(app_control_launcher.ControlLauncher)
	if !ok {
		l.Info("The catalogue is not available; skip listing features.")
		return nil
	}

	ui.Header(z.SectionDescription)
	{
		uit := ui.InfoTable("description")
		uit.Header(z.HeaderKey, z.HeaderDesc)
		for _, f := range cl.Catalogue().Features() {
			uit.Row(
				app_msg.Raw(f.OptInName(f)),
				app_feature.OptInDescription(f),
			)
		}
		uit.Flush()
	}

	ui.Header(z.SectionSettings)
	{
		uit := ui.InfoTable("settings")
		uit.Header(z.HeaderKey, z.HeaderStatus, z.HeaderOptInUser, z.HeaderOptInTimestamp)
		for _, f := range cl.Catalogue().Features() {
			if g, found := c.Feature().OptInGet(f); found {
				uit.RowRaw(
					g.OptInName(g),
					strconv.FormatBool(g.OptInIsEnabled()),
					g.OptInUser(),
					g.OptInTimestamp(),
				)
			}
		}
		uit.Flush()
	}
	return nil
}

func (z *Features) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Features{}, rc_recipe.NoCustomValues)
}
