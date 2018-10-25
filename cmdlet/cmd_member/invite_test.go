package cmd_member

import (
	"fmt"
	"github.com/watermint/toolbox/cmdlet"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdMemberInvite_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "team_member_invite")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	csvContent := ""
	for i := 0; i < 3; i++ {
		csvContent = csvContent + fmt.Sprintf("toolbox%d@example.com,toolbox,%d\n", i, i)
	}
	if _, err := csvFile.WriteString(csvContent); err != nil {
		t.Error(err)
	}
	if err := csvFile.Close(); err != nil {
		t.Error(err)
	}

	cmdlet.CmdTest(t, NewCmdMember(), []string{"invite", "-csv", csvFile.Name()})
}
