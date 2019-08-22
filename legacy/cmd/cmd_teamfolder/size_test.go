package cmd_teamfolder

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdTeamFolderSize_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "teamfolder_size")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmd2.CmdTest(t, NewCmdTeamFolder(), []string{"size", "-report-path", csvFile.Name()})
}
