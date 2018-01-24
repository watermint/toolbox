package cmd_event

import (
	"encoding/json"
	"flag"
	"github.com/cihub/seelog"
	"github.com/dropbox/dropbox-sdk-go-unofficial/dropbox/team_log"
	"github.com/watermint/toolbox/api"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"os"
)

type CmdEventList struct {
	optAll       bool
	optOutFile   string
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

	descOutFile := "Output file name"
	f.StringVar(&c.optOutFile, "out", "", descOutFile)

	return f
}

func (c *CmdEventList) UsageTmpl() string {
	return `
Usage: {{.Command}}
`
}

type outputPrepare func(list *CmdEventList) error
type outputFunc func(json.RawMessage)
type outputClose func()

func (c *CmdEventList) Exec(cc cmdlet.CommandletContext) (err error) {
	if _, err = cmdlet.ParseFlags(cc, c); err != nil {
		return
	}

	c.infraContext.Startup()
	defer c.infraContext.Shutdown()
	seelog.Debugf("copy:%s", util.MarshalObjectToString(c))
	c.apiContext, err = c.infraContext.LoadOrAuthBusinessAudit()

	arg := team_log.NewGetTeamEventsArg()
	res, err := c.apiContext.TeamLogImpl().RawGetEvents(arg)
	if err != nil {
		seelog.Warnf("Unable to get logs : error[%s]", err)
		return
	}

	var f *os.File
	var op outputPrepare
	var of outputFunc
	var oc outputClose

	if c.optOutFile != "" {
		op = func(c *CmdEventList) error {
			f, err = os.Create(c.optOutFile)
			return err
		}
		of = func(e json.RawMessage) {
			f.Write(e)
			f.WriteString("\n")
		}
		oc = func() {
			if f != nil {
				f.Close()
			}
		}
	} else {
		of = func(e json.RawMessage) {
			seelog.Infof("Event: %s", e)
		}
		op = func(list *CmdEventList) error {
			return nil
		}
		oc = func() {}
	}

	if err = op(c); err != nil {
		seelog.Warnf("Failed to prepare output. error[%s]", err)
		return
	}
	defer oc()

	for _, e := range res.Events {
		of(e)
	}

	if c.optAll {
		for res.HasMore {
			res, err = c.apiContext.TeamLogImpl().RawGetEventsContinue(team_log.NewGetTeamEventsContinueArg(res.Cursor))
			if err != nil {
				seelog.Warnf("Unable to get logs : error[%s]", err)
				return
			}
			for _, e := range res.Events {
				of(e)
			}
		}
	}

	return
}
