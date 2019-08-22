package cmd_audit

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"io/ioutil"
	"testing"
)

func TestCmdTeamAuditSharing_Exec(t *testing.T) {
	reportPath, err := ioutil.TempDir("", "audit")
	if err != nil {
		t.Error(err)
	}
	//	defer os.RemoveAll(reportPath)

	cmd2.CmdTest(t, NewCmdTeamAudit(), []string{"sharing", "-report-path", reportPath})
}
