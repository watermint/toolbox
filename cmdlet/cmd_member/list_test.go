package cmd_member

import (
	"github.com/watermint/toolbox/cmdlet"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdMemberList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "member_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmdlet.CmdTest(t, NewCmdMember(), []string{"list", "-report-path", csvFile.Name()})
}
