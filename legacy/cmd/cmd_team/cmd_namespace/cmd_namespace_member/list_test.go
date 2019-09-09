package cmd_namespace_member

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
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

	cmd2.CmdTest(t, NewCmdTeamNamespaceMember(), []string{"list", "-report-path", csvFile.Name()})
}

func TestCmdTeamNamespaceMemberList_ExpandGroup(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "group_member_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmd2.CmdTest(t, NewCmdTeamNamespaceMember(), []string{"list", "-report-path", csvFile.Name()})
}
