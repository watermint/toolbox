package cmd_update

import (
	"github.com/watermint/toolbox/cmdlet"
	"testing"
)

func TestCmdTeamSharedLinkUpdateExpire_Exec(t *testing.T) {
	cmdlet.CmdTest(t, NewCmdMemberSharedLinkUpdate(), []string{"expire", "-days", "365"})
}
