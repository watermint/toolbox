package cmd_teamfolder

import (
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/cmd/cmd_team/cmd_teamfolder/cmd_teamfolder_file"
)

func NewCmdTeamTeamFolder() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "teamfolder",
		CommandDesc: "cmd.team.teamfolder.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamTeamFolderList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamTeamFolderSize{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamTeamFolderMirror{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			cmd_teamfolder_file.NewCmdTeamTeamFolderFile(),
		},
	}
}
