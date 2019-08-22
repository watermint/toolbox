package cmd_linkedapp

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdMemberLinkedAppList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "member_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmd2.CmdTest(t, NewCmdMemberLinkedApp(), []string{"list", "-report-path", csvFile.Name()})
}
