package cmd_team

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdTeamFeature_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdTeam(), []string{"feature"})
}
