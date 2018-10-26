package cmd_teamfolder

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdTeamTeamFolderList_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdTeamTeamFolder(), []string{"list"})
}
