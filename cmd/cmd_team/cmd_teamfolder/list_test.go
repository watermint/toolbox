package cmd_teamfolder

import (
	"github.com/watermint/toolbox/cmd"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdTeamTeamFolderList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "teamfolder_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmd.CmdTest(t, NewCmdTeamTeamFolder(), []string{"list", "-report-path", csvFile.Name()})
}
