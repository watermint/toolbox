package cmd_team

import (
	"github.com/watermint/toolbox/cmdlet"
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

	cmdlet.CmdTest(t, NewCmdTeam(), []string{"info", "-report-path", csvFile.Name()})
}
