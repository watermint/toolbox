package cmd_root

import (
	"flag"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_file"
	"github.com/watermint/toolbox/cmd/cmd_group"
	"github.com/watermint/toolbox/cmd/cmd_member"
	"github.com/watermint/toolbox/cmd/cmd_team"
	"os"
)

type Commands struct {
	rootCmd cmd.Commandlet
}

func NewCommands() Commands {
	return Commands{
		rootCmd: &cmd.CommandletGroup{
			CommandName: os.Args[0],
			SubCommands: []cmd.Commandlet{
				cmd_file.NewCmdFile(),
				cmd_team.NewCmdTeam(),
				cmd_member.NewCmdMember(),
				cmd_group.NewCmdGroup(),
			},
		},
	}
}

func (z *Commands) RootCommand() cmd.Commandlet {
	return z.rootCmd
}

func (z *Commands) Exec(ec *app.ExecContext, args []string) {
	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	ec.PrepareFlags(f)
	z.rootCmd.Init(nil)
	z.rootCmd.FlagConfig(f)
	z.rootCmd.Setup(ec)
	z.rootCmd.Exec(args[1:])
}
