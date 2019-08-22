package cmd_root

import (
	"flag"
	app2 "github.com/watermint/toolbox/legacy/app"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	cmd_auth2 "github.com/watermint/toolbox/legacy/cmd/cmd_auth"
	cmd_dev2 "github.com/watermint/toolbox/legacy/cmd/cmd_dev"
	cmd_file2 "github.com/watermint/toolbox/legacy/cmd/cmd_file"
	cmd_group2 "github.com/watermint/toolbox/legacy/cmd/cmd_group"
	cmd_license2 "github.com/watermint/toolbox/legacy/cmd/cmd_license"
	cmd_member2 "github.com/watermint/toolbox/legacy/cmd/cmd_member"
	cmd_sharedfolder2 "github.com/watermint/toolbox/legacy/cmd/cmd_sharedfolder"
	cmd_sharedlink2 "github.com/watermint/toolbox/legacy/cmd/cmd_sharedlink"
	cmd_team2 "github.com/watermint/toolbox/legacy/cmd/cmd_team"
	cmd_teamfolder2 "github.com/watermint/toolbox/legacy/cmd/cmd_teamfolder"
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
				cmd_file2.NewCmdFile(),
				cmd_team2.NewCmdTeam(),
				cmd_member2.NewCmdMember(),
				cmd_group2.NewCmdGroup(),
				cmd_sharedfolder2.NewSharedFolder(),
				cmd_sharedlink2.NewCmdSharedLink(),
				cmd_teamfolder2.NewCmdTeamFolder(),
				cmd_dev2.NewCmdDev(),
				cmd_auth2.NewCmdAuth(),
				&cmd_license2.CmdLicense{
					SimpleCommandlet: &cmd2.SimpleCommandlet{},
				},
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
