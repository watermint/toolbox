package main

import (
	"fmt"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/cmdlet/cmd_group"
	"github.com/watermint/toolbox/cmdlet/cmd_member"
	"github.com/watermint/toolbox/cmdlet/cmd_team"
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
		fmt.Fprintln(os.Stdout, usage)
	} else {
		panic(fmt.Sprintf("template erorr: err[%s] tmpl[%s]", tmplErr, tmpl))
	}

	fg := cl.FlagSet()
	if fg != nil {
		fg.SetOutput(os.Stdout)
		fmt.Println("Options:")
		fg.PrintDefaults()
		fmt.Println("")
	}

	fmt.Printf("Error: %s\n", err)
}

func main() {
	rootCmd := &cmdlet.RootCommandlet{
		ParentCommandlet: &cmdlet.ParentCommandlet{
			SubCommands: []cmdlet.Commandlet{
				cmd_member.NewCmdMember(),
				cmd_group.NewCmdGroup(),
				cmd_team.NewCmdTeam(),
			},
		},
	}

	err := rootCmd.Exec(cmdlet.CommandletContext{
		Command: rootCmd,
		Cmd:     os.Args[0],
		Args:    os.Args[1:],
	})
	switch ce := err.(type) {
	case nil:
		// nop
		os.Exit(TBL_EXIT_SUCCESS)

	case *cmdlet.CommandError:
		ce.PrintError()
		os.Exit(TBL_EXIT_FAILURE)

	case *cmdlet.CommandShowUsageError:
		printUsage(ce.Context, ce.Context.Command, err)
	}
}
