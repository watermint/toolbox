package main

import (
	"fmt"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_group"
	"github.com/watermint/toolbox/cmdlet/cmd_member"
	"github.com/watermint/toolbox/infra/util"
	"os"
)

const (
	TBL_EXIT_SUCCESS = 0
	TBL_EXIT_FAILURE = 1
)

func printUsage(cc cmdlet.CommandletContext, cl cmdlet.Commandlet, err error) {
	tmpl := cl.UsageTmpl()
	usage, tmplErr := util.CompileTemplate(tmpl,
		struct {
			Command string
		}{
			Command: cc.Cmd,
		})
	if tmplErr == nil {
		fmt.Fprintln(cc.OutDest, usage)
	} else {
		fmt.Fprintf(cc.OutDest, "template erorr: %s", tmplErr)
		fmt.Fprintln(cc.OutDest, tmpl)
	}

	fg := cl.FlagSet()
	if fg != nil {
		fg.SetOutput(cc.OutDest)
		fmt.Fprintln(cc.OutDest, "Options:")
		fg.PrintDefaults()
		fmt.Fprintln(cc.OutDest, "")
	}

	switch csue := err.(type) {
	case *cmdlet.CommandShowUsageError:
		if csue.Instruction != "" {
			fmt.Fprintln(cc.OutDest, csue.Instruction)
		}

	case *cmdlet.CommandError:
		// nop

	case nil:
		// nop

	default:
		fmt.Fprintf(cc.OutDest, "Error: %s\n", err)
	}
}

func main() {
	cmdMember := &cmd_member.CmdMember{
		ParentCommandlet: &cmdlet.ParentCommandlet{
			SubCommands: []cmdlet.Commandlet{
				cmd_member.NewCmdMemberInvite(),
				cmd_member.NewCmdMemberList(),
			},
		},
	}
	cmdGroup := &cmd_group.CmdGroup{
		ParentCommandlet: &cmdlet.ParentCommandlet{
			SubCommands: []cmdlet.Commandlet{
				cmd_group.NewCmdGroupList(),
			},
		},
	}
	rootCmd := &cmdlet.RootCommandlet{
		ParentCommandlet: &cmdlet.ParentCommandlet{
			SubCommands: []cmdlet.Commandlet{
				cmdMember,
				cmdGroup,
			},
		},
	}

	err := rootCmd.Exec(cmdlet.CommandletContext{
		OutDest:            os.Stdout,
		OutQuiet:           false,
		OutMachineFriendly: true,
		Command:            rootCmd,
		Cmd:                os.Args[0],
		Args:               os.Args[1:],
	})
	switch ce := err.(type) {
	case nil:
		// nop

	case *cmdlet.CommandError:
		ce.PrintError()
		os.Exit(TBL_EXIT_FAILURE)

	case *cmdlet.CommandShowUsageError:
		printUsage(ce.Context, ce.Context.Command, err)
	}
}
