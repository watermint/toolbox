package cmd_group

import (
	"github.com/watermint/toolbox/cmdlet"
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

	cmdlet.CmdTest(t, NewCmdGroup(), []string{"list", "-report-path", csvFile.Name()})
}
