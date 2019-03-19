package cmd_teamfolder

import (
	"github.com/watermint/toolbox/cmd"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdTeamFolderList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "teamfolder_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmd.CmdTest(t, NewCmdTeamFolder(), []string{"list", "-report-path", csvFile.Name()})
}
