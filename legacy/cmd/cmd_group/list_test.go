package cmd_group

import (
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
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

	cmd2.CmdTest(t, NewCmdGroup(), []string{"list", "-report-path", csvFile.Name()})
}
