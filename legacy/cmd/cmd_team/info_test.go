package cmd_team

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdTeamInfo_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "team_info")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmd2.CmdTest(t, NewCmdTeam(), []string{"info", "-report-path", csvFile.Name()})
}
