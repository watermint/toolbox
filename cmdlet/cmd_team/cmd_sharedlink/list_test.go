package cmd_sharedlink

import (
	"github.com/watermint/toolbox/cmdlet"
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

	cmdlet.CmdTest(t, NewCmdTeamSharedLink(), []string{"list", "-report-path", csvFile.Name()})
}
