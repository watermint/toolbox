package cmd_namespace_member

import (
	"github.com/watermint/toolbox/cmdlet"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdTeamNamespaceMemberList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "namespace_member_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmdlet.CmdTest(t, NewCmdTeamNamespaceMember(), []string{"list", "-report-path", csvFile.Name()})
}

func TestCmdTeamNamespaceMemberList_ExpandGroup(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "group_member_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmdlet.CmdTest(t, NewCmdTeamNamespaceMember(), []string{"list", "-expand-group", "-report-path", csvFile.Name()})
}
