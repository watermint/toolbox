package cmd_namespace

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdTeamNamespaceList_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdTeamNamespace(), []string{"list"})
}
