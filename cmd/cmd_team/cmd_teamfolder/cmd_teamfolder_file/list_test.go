package cmd_teamfolder_file

import (
	"github.com/watermint/toolbox/cmd"
	"io/ioutil"
	"testing"
)

func TestCmdTeamTeamFolderFileList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "member_list")
	if err != nil {
		t.Error(err)
	}
	//	defer os.Remove(csvFile.Name())

	cmd.CmdTest(t, NewCmdTeamTeamFolderFile(), []string{"list", "-report-path", csvFile.Name()})
}
