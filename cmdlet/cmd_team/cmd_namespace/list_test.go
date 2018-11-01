package cmd_namespace

import (
	"github.com/watermint/toolbox/cmdlet"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdTeamNamespaceList_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "namespace_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmdlet.CmdTest(t, NewCmdTeamNamespace(), []string{"list", "-report-path", csvFile.Name()})
}

func TestCmdTeamNamespaceList_ExecCsv(t *testing.T) {
	cmdlet.CmdTest(
		t,
		NewCmdTeamNamespace(),
		[]string{
			"list",
			"-report-format",
			"csv",
			"-report-omit-header=false",
		},
	)
}
