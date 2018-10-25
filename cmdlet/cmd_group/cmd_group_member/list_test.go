package cmd_group_member

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdGroupMemberList_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdGroupMember(), []string{"list"})
}
