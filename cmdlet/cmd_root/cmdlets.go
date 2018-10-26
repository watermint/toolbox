package cmd_root

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_group"
	"github.com/watermint/toolbox/cmdlet/cmd_member"
	"github.com/watermint/toolbox/cmdlet/cmd_team"
	"github.com/watermint/toolbox/infra"
	"os"
)

type Commands struct {
	rootCmd cmdlet.Commandlet
}

func NewCommands() Commands {
	return Commands{
		rootCmd: &cmdlet.CommandletGroup{
			CommandName: os.Args[0],
			SubCommands: []cmdlet.Commandlet{
				cmd_team.NewCmdTeam(),
				cmd_member.NewCmdMember(),
				cmd_group.NewCmdGroup(),
			},
		},
	}
}

func (c *Commands) RootCommand() cmdlet.Commandlet {
	return c.rootCmd
}

func (c *Commands) Exec(ec *infra.ExecContext, args []string) {
	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	ec.PrepareFlags(f)
	c.rootCmd.Init(nil)
	c.rootCmd.FlagConfig(f)
	c.rootCmd.Exec(ec, args[1:])
}
