package cmd_sharedlink

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdTeamSharedLinkList_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdTeamSharedLink(), []string{"list"})
}
