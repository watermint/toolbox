package cmd_device

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

func NewCmdTeamDevice() cmd2.Commandlet {
	return &cmd2.CommandletGroup{
		CommandName: "device",
		CommandDesc: "cmd.team.device.desc",
		SubCommands: []cmd2.Commandlet{
			&CmdTeamDeviceList{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
			&CmdTeamDeviceUnlink{
				SimpleCommandlet: &cmd2.SimpleCommandlet{},
			},
		},
	}
}
