package cmd_teamfolder_file

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"io/ioutil"
	"testing"
)

func TestCmdTeamFolderFileList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "member_list")
	if err != nil {
		t.Error(err)
	}
	//	defer os.Remove(csvFile.Name())

	cmd2.CmdTest(t, NewCmdTeamFolderFile(), []string{"list", "-report-path", csvFile.Name()})
}