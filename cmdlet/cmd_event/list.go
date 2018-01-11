package cmd_event

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_log"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
)

type CmdEventList struct {
	apiContext   *api.ApiContext
	infraContext *infra.InfraContext
}

func NewCmdEventList() *CmdEventList {
	c := CmdEventList{
		infraContext: &infra.InfraContext{},
	}
	return &c
}

func (c *CmdEventList) Name() string {
	return "list"
}

func (c *CmdEventList) Desc() string {
	return "Event list"
}

func (c *CmdEventList) FlagSet() (f *flag.FlagSet) {
	f = flag.NewFlagSet(c.Name(), flag.ExitOnError)

	return f
}

func (c *CmdEventList) UsageTmpl() string {
	return `
Usage: {{.Command}}
`
}

func (c *CmdEventList) Exec(cc cmdlet.CommandletContext) (err error) {
	if _, err = cmdlet.ParseFlags(cc, c); err != nil {
		return err
	}

	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("copy:%s", util.MarshalObjectToString(c))
	c.apiContext, err = c.infraContext.LoadOrAuthBusinessAudit()

	arg := team_log.NewGetTeamEventsArg()
	res, err := c.apiContext.TeamLogImpl().RawGetEvents(arg)
	if err != nil {
		seelog.Warnf("Unable to get logs : error[%s]", err)
		return err
	}
	seelog.Infof("HasMore: %t", res.HasMore)
	for _, e := range res.Events {
		seelog.Infof("Event: %s", e)
	}

	return
}
