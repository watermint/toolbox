package main

import (
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx/task/member"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"github.com/watermint/toolbox/workflow"
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
	PipelinePoc()

	//cmdFile := &cmd_file.CmdFile{
	//	ParentCommandlet: &cmdlet.ParentCommandlet{
	//		SubCommands: []cmdlet.Commandlet{
	//			cmd_file.NewCmdFileCopy(),
	//			cmd_file.NewCmdFileMove(),
	//			cmd_file.NewCmdFileUpload(),
	//		},
	//	},
	//}
	//cmdMember := &cmd_member.CmdMember{
	//	ParentCommandlet: &cmdlet.ParentCommandlet{
	//		SubCommands: []cmdlet.Commandlet{
	//			cmd_member.NewCmdMemberInvite(),
	//		},
	//	},
	//}
	//cmdEvent := &cmd_event.CmdEvent{
	//	ParentCommandlet: &cmdlet.ParentCommandlet{
	//		SubCommands: []cmdlet.Commandlet{
	//			cmd_event.NewCmdEventList(),
	//		},
	//	},
	//}
	//cmdTeam := &cmd_team.CmdTeam{
	//	ParentCommandlet: &cmdlet.ParentCommandlet{
	//		SubCommands: []cmdlet.Commandlet{
	//			cmd_team.NewCmdTeamScan(),
	//		},
	//	},
	//}
	//rootCmd := &cmdlet.RootCommandlet{
	//	ParentCommandlet: &cmdlet.ParentCommandlet{
	//		SubCommands: []cmdlet.Commandlet{
	//			cmdFile,
	//			cmdEvent,
	//			cmdMember,
	//			cmdTeam,
	//		},
	//	},
	//}
	//
	//err := rootCmd.Exec(cmdlet.CommandletContext{
	//	OutDest:            os.Stdout,
	//	OutQuiet:           false,
	//	OutMachineFriendly: true,
	//	Command:            rootCmd,
	//	Cmd:                os.Args[0],
	//	Args:               os.Args[1:],
	//})
	//switch ce := err.(type) {
	//case nil:
	//	// nop
	//
	//case *cmdlet.CommandError:
	//	ce.PrintError()
	//	os.Exit(TBL_EXIT_FAILURE)
	//
	//case *cmdlet.CommandShowUsageError:
	//	printUsage(ce.Context, ce.Context.Command, err)
	//}
}

func PipelinePoc() error {
	c := infra.InfraContext{}
	c.Startup()
	defer c.Shutdown()

	apiMgmt, err := c.LoadOrAuthBusinessManagement()
	if err != nil {
		seelog.Warnf("Unable to acquire token : error[%s]", err)
		return err
	}

	p := workflow.Pipeline{
		Infra: &c,
		Stages: []workflow.Worker{
			&member.WorkerTeamMemberInviteLoaderCsv{},
			&member.WorkerTeamMemberInvite{ApiManagement: apiMgmt, Silent: true},
			&member.WorkerTeamMemberInviteResultAsync{ApiManagement: apiMgmt},
			&member.WorkerTeamMemberInviteResultReduce{},
		},
	}
	p.Init()
	defer p.Close()

	p.Enqueue(member.NewTaskTeamMemberInviteLoaderCsv("invite.csv"))
	p.Loop()

	return nil
}
