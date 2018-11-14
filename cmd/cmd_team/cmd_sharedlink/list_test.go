package cmd_sharedlink

import (
	"github.com/watermint/toolbox/cmd"
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

	cmd.CmdTest(t, NewCmdTeamSharedLink(), []string{"list", "-report-path", csvFile.Name()})
}
