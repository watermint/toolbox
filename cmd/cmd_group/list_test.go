package cmd_group

import (
	"github.com/watermint/toolbox/cmd"
	"io/ioutil"
	"os"
	"testing"
)

func TestCmdGrouplist_Exec(t *testing.T) {
	csvFile, err := ioutil.TempFile("", "group_list")
	if err != nil {
		t.Error(err)
	}
	defer os.Remove(csvFile.Name())

	cmd.CmdTest(t, NewCmdGroup(), []string{"list", "-report-path", csvFile.Name()})
}
