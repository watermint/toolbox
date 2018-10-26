package cmd_namespace_member

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdTeamNamespaceMemberList_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdTeamNamespaceMember(), []string{"list"})
}

func TestCmdTeamNamespaceMemberList_ExpandGroup(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdTeamNamespaceMember(), []string{"list", "-expand-group"})
}
