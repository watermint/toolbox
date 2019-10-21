package cmd_root

import (
	"flag"
	app2 "github.com/watermint/toolbox/legacy/app"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_dev2 "github.com/watermint/toolbox/legacy/cmd/cmd_dev"
	cmd_member2 "github.com/watermint/toolbox/legacy/cmd/cmd_member"
	"os"
)

type Commands struct {
	rootCmd cmd2.Commandlet
}

func NewCommands() Commands {
	return Commands{
		rootCmd: &cmd2.CommandletGroup{
			CommandName: os.Args[0],
			SubCommands: []cmd2.Commandlet{
				cmd_member2.NewCmdMember(),
				cmd_dev2.NewCmdDev(),
			},
		},
	}
}

func (z *Commands) RootCommand() cmd2.Commandlet {
	return z.rootCmd
}

func (z *Commands) Exec(ec *app2.ExecContext, args []string) {
	f := flag.NewFlagSet(args[0], flag.ExitOnError)
	ec.PrepareFlags(f)
	z.rootCmd.Init(nil)
	z.rootCmd.FlagConfig(f)
	z.rootCmd.Setup(ec)
	z.rootCmd.Exec(args[1:])
}
