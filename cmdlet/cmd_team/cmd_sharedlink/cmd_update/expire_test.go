package cmd_update

import (
	"github.com/watermint/toolbox/cmdlet"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdTeamSharedLinkUpdateExpire_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "sharedlink_update_expire")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmdlet.CmdTest(t, NewCmdMemberSharedLinkUpdate(), []string{"expire", "-days", "365", "-report-path", csvFile.Name()})
}
