package cmd_sharedlink

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdTeamSharedLinkList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "sharedlink_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmd2.CmdTest(t, NewCmdTeamSharedLink(), []string{"list", "-report-path", csvFile.Name()})
}
