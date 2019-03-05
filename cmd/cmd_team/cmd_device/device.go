package cmd_device

import "github.com/watermint/toolbox/cmd"

func NewCmdTeamDevice() cmd.Commandlet {
	return &cmd.CommandletGroup{
		CommandName: "device",
		CommandDesc: "cmd.team.device.desc",
		SubCommands: []cmd.Commandlet{
			&CmdTeamDeviceList{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
			&CmdTeamDeviceUnlink{
				SimpleCommandlet: &cmd.SimpleCommandlet{},
			},
		},
	}
}
