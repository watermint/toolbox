package cmd_teamfolder

import (
	"github.com/watermint/toolbox/cmd"
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

	cmd.CmdTest(t, NewCmdTeamFolder(), []string{"size", "-report-path", csvFile.Name()})
}
