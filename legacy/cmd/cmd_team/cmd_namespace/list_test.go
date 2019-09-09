package cmd_namespace

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
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

	cmd2.CmdTest(t, NewCmdTeamNamespace(), []string{"list", "-report-path", csvFile.Name()})
}

func TestCmdTeamNamespaceList_ExecCsv(t *testing.T) {
	cmd2.CmdTest(
		t,
		NewCmdTeamNamespace(),
		[]string{
			"list",
		},
	)
}
