package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/infra"
)

type CmdTeamFeature struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     cmdlet.Report
}

func (c *CmdTeamFeature) Name() string {
	return "feature"
}

func (c *CmdTeamFeature) Desc() string {
	return "List team feature values"
}

func (CmdTeamFeature) Usage() string {
	return ""
}

func (c *CmdTeamFeature) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)
}

func (c *CmdTeamFeature) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()

	apiInfo, err := ec.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.Open()
	defer c.report.Close()

	l := dbx_team.FeatureList{
		OnError: cmdlet.DefaultErrorHandler,
		OnEntry: func(feature *dbx_team.Feature) bool {
			c.report.Report(feature)
			return true
		},
	}
	l.List(apiInfo)
}
