package cmd_member

import (
	"github.com/watermint/toolbox/cmdlet"
	"io/ioutil"
	"testing"
)

func TestCmdMemberList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempDir("", "member_list")
	if err != nil {
		t.Error(err)
	}

	cmdlet.CmdTest(t, NewCmdMember(), []string{"list", "-report-path", csvFile})
}
