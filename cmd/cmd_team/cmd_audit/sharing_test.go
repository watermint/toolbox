package cmd_audit

import (
	"github.com/watermint/toolbox/cmd"
	"io/ioutil"
	"testing"
)

func TestCmdTeamAuditSharing_Exec(t *testing.T) {
	reportPath, err := ioutil.TempDir("", "audit")
	if err != nil {
		t.Error(err)
	}
	//	defer os.RemoveAll(reportPath)

	cmd.CmdTest(t, NewCmdTeamAudit(), []string{"sharing", "-report-path", reportPath})
}
