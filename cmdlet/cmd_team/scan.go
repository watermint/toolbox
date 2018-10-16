package cmd_team

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
)

type CmdTeamScan struct {
	apiContext   *api.ApiContext
	infraContext *infra.InfraContext
}

func NewCmdTeamScan() *CmdTeamScan {
	c := CmdTeamScan{
		infraContext: &infra.InfraContext{},
	}
	return &c
}

func (c *CmdTeamScan) Name() string {
	return "scan"
}

func (c *CmdTeamScan) Desc() string {
	return "Scan team state"
}

func (c *CmdTeamScan) UsageTmpl() string {
	return `
Usage: {{.Command}}
`
}

func (c *CmdTeamScan) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	c.infraContext.PrepareFlags(f)
	return f
}

func (c *CmdTeamScan) Exec(cc cmdlet.CommandletContext) error {
	_, err := cmdlet.ParseFlags(cc, c)
	if err != nil {
		return err
	}
	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("invite:%s", util.MarshalObjectToString(c))
	c.apiContext, err = c.infraContext.LoadOrAuthBusinessFile()

	return &cmdlet.CommandError{
		Context:     cc,
		ReasonTag:   "team/scan:no_impl",
		Description: "No impl.",
	}
}
