package cmd_root

import (
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/infra"
	"testing"
)

func TestCommands(t *testing.T) {
	ec := &infra.ExecContext{}
	if ec.StartupForTest() != nil {
		return
	}

	// Finish tests if tokens file not available
	if !ec.IsTokensAvailable() {
		seelog.Info("Skip tests")
		return
	}

	cmds := NewCommands()

	cmdTest := func(name string, args []string) {
		seelog.Infof("Testing [%s]", name)
		aa := make([]string, 1)
		aa[0] = "cmdlet_test"
		aa = append(aa, args...)
		seelog.Info("Testing Args[%v]", aa)
		cmds.Exec(ec, aa)
	}

	//cmdTest("root help", []string{})
	//cmdTest("root help", []string{"help"})
	//cmdTest("group help", []string{"group"})
	//cmdTest("group help", []string{"group", "help"})
	cmdTest("group list", []string{"group", "list"})
	//cmdTest("group list help", []string{"group", "list", "help"})
	//cmdTest("group member", []string{"group", "member"})
	//cmdTest("group member help", []string{"group", "member", "help"})
	cmdTest("group member list", []string{"group", "member", "list"})
	//cmdTest("member help", []string{"member"})
	//cmdTest("member help", []string{"member", "help"})
	cmdTest("member list", []string{"member", "list"})
	//cmdTest("member list help", []string{"member", "list", "help"})
	//cmdTest("member invite (missing file)", []string{"member", "invite"})
	cmdTest("team linkedapp list", []string{"team", "app", "list"})
	cmdTest("team namespace list", []string{"team", "namespace", "list"})
	cmdTest("team namespace member list", []string{"team", "namespace", "member", "list"})
	cmdTest("team sharedlink list", []string{"team", "sharedlink", "list"})
	cmdTest("team teamfolder list", []string{"team", "teamfolder", "list"})
	cmdTest("team feature", []string{"team", "feature"})
	cmdTest("team info", []string{"team", "info"})
}
