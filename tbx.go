package main

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_group"
	"github.com/watermint/toolbox/cmdlet/cmd_member"
	"github.com/watermint/toolbox/cmdlet/cmd_team"
	"github.com/watermint/toolbox/infra"
	"os"
)

func main() {
	rootCmd := cmdlet.CommandletGroup{
		CommandName: os.Args[0],
		SubCommands: []cmdlet.Commandlet{
			cmd_team.NewCmdTeam(),
			cmd_member.NewCmdMember(),
			cmd_group.NewCmdGroup(),
		},
	}

	ec := &infra.ExecContext{}
	f := flag.NewFlagSet(os.Args[0], flag.ExitOnError)
	ec.PrepareFlags(f)
	rootCmd.Init(nil)
	rootCmd.FlagConfig(f)
	rootCmd.Exec(ec, os.Args[1:])
}
