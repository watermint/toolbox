package cmd_team

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/report"
)

type CmdTeamFeature struct {
	*cmd.SimpleCommandlet

	apiContext *dbx_api.Context
	report     report.Factory
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

func (c *CmdTeamFeature) Exec(args []string) {
	apiInfo, err := c.ExecContext.LoadOrAuthBusinessInfo()
	if err != nil {
		return
	}

	c.report.Init(c.Log())
	defer c.report.Close()

	l := dbx_team.FeatureList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(feature *dbx_team.Feature) bool {
			c.report.Report(feature)
			return true
		},
	}
	l.List(apiInfo)
}
