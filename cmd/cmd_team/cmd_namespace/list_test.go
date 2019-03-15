package cmd_namespace

import (
	"github.com/watermint/toolbox/cmd"
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

	cmd.CmdTest(t, NewCmdTeamNamespace(), []string{"list", "-report-path", csvFile.Name()})
}

func TestCmdTeamNamespaceList_ExecCsv(t *testing.T) {
	cmd.CmdTest(
		t,
		NewCmdTeamNamespace(),
		[]string{
			"list",
		},
	)
}
