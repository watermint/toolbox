package cmd_event

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_common"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_log"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"time"
)

type CmdEventList struct {
	optAll       bool
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

	descAll := "List all events that match to criteria"
	f.BoolVar(&c.optAll, "all", false, descAll)

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
	for _, e := range res.Events {
		seelog.Infof("Event: %s", e)
	}

	if c.optAll {
		for res.HasMore {
			res, err = c.apiContext.TeamLogImpl().RawGetEventsContinue(team_log.NewGetTeamEventsContinueArg(res.Cursor))
			if err != nil {
				seelog.Warnf("Unable to get logs : error[%s]", err)
				return err
			}
			for _, e := range res.Events {
				seelog.Infof("Event: %s", e)
			}
		}
	}

	return
}
