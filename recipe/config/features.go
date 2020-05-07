package config

import (
	"github.com/watermint/toolbox/infra/control/app_catalogue"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_feature"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
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
	ui := c.UI()
	cat := app_catalogue.Current()

	ui.Header(z.SectionDescription)
	ui.WithTable("Description", func(uit app_ui.Table) {
		uit.Header(z.HeaderKey, z.HeaderDesc)
		for _, f := range cat.Features() {
			uit.Row(
				app_msg.Raw(app_feature.OptInName(f)),
				app_feature.OptInDescription(f),
			)
		}
	})

	ui.Header(z.SectionSettings)
	ui.WithTable("Settings", func(uit app_ui.Table) {
		uit.Header(z.HeaderKey, z.HeaderStatus, z.HeaderOptInUser, z.HeaderOptInTimestamp)
		for _, f := range cat.Features() {
			if g, found := c.Feature().OptInGet(f); found {
				uit.RowRaw(
					app_feature.OptInName(g),
					strconv.FormatBool(g.OptInIsEnabled()),
					g.OptInUser(),
					g.OptInTimestamp(),
				)
			}
		}
	})

	return nil
}

func (z *Features) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Features{}, rc_recipe.NoCustomValues)
}
