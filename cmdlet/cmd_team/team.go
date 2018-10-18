package cmd_team

import "github.com/watermint/toolbox/cmdlet"

type CmdTeam struct {
	*cmdlet.ParentCommandlet
}

func (c *CmdTeam) Name() string {
	return "team"
}

func (c *CmdTeam) Desc() string {
	return "Dropbox Business team management"
}

func NewCmdTeam() cmdlet.Commandlet {
	return &CmdTeam{
		ParentCommandlet: &cmdlet.ParentCommandlet{
			SubCommands: []cmdlet.Commandlet{
				NewCmdTeamInfo(),
			},
		},
	}
}
