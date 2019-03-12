package cmd_member_update

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
)

type CmdMemberUpdateEmail struct {
	*cmd.SimpleCommandlet
}

func (*CmdMemberUpdateEmail) Name() string {
	return "email"
}

func (*CmdMemberUpdateEmail) Desc() string {
	return "cmd.member.update.email.desc"
}

func (*CmdMemberUpdateEmail) Usage() func(cmd.CommandUsage) {
	return nil
}

func (*CmdMemberUpdateEmail) FlagConfig(f *flag.FlagSet) {

}

func (*CmdMemberUpdateEmail) Exec(args []string) {

}
