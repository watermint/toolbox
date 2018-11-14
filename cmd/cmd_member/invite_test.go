package cmd_member

import (
	"fmt"
	"github.com/watermint/toolbox/cmd"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdMemberInvite_Exec(t *testing.T) {
	reportCsvFile, err := ioutil.TempFile("", "member_invite_report")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(reportCsvFile.Name())

	inviteCsvFile, err := ioutil.TempFile("", "member_invite")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(inviteCsvFile.Name())

	csvContent := ""
	for i := 0; i < 3; i++ {
		csvContent = csvContent + fmt.Sprintf("toolbox%d@example.com,toolbox,%d\n", i, i)
	}
	if _, err := inviteCsvFile.WriteString(csvContent); err != nil {
		t.Error(err)
	}
	if err := inviteCsvFile.Close(); err != nil {
		t.Error(err)
	}

	cmd.CmdTest(
		t,
		NewCmdMember(),
		[]string{
			"invite",
			"-csv",
			inviteCsvFile.Name(),
			"-report-path",
			reportCsvFile.Name(),
		},
	)
}
