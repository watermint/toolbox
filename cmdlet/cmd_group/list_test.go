package cmd_group

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdGrouplist_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdGroup(), []string{"list"})
}
