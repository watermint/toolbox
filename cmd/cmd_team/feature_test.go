package cmd_team

import (
	"github.com/watermint/toolbox/cmd"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdTeamFeature_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "feature")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmd.CmdTest(t, NewCmdTeam(), []string{"feature", "-report-path", csvFile.Name()})
}
