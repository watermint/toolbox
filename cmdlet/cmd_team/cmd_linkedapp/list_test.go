package cmd_linkedapp

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdMemberLinkedAppList_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdMemberLinkedApp(), []string{"list"})
}
