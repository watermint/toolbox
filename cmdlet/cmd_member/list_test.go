package cmd_member

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdMemberList_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdMember(), []string{"list"})
}
