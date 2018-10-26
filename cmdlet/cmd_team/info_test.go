package cmd_team

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdTeamInfo_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdTeam(), []string{"info"})
}
