package cmdlet

import (
	"flag"
	"fmt"
	"io"
)

type CommandletCommonContext struct {
	OutQuiet           bool
	OutMachineFriendly bool
	Proxy              string
	WorkPath           string
	LogRolls           int
	ContractName       string
}

func (c *CommandletCommonContext) PrepareFlags(f *flag.FlagSet) {
	//TODO
}

type CommandletContext struct {
	Cmd                string
	Args               []string
	Command            Commandlet
	OutDest            io.Writer
	OutQuiet           bool
	OutMachineFriendly bool
}

func NewCommandletContext(cmd string, args []string, cmdlet Commandlet, cc CommandletContext) CommandletContext {
	return CommandletContext{
		Cmd:                cmd,
		Args:               args,
		Command:            cmdlet,
		OutDest:            cc.OutDest,
		OutQuiet:           cc.OutQuiet,
		OutMachineFriendly: cc.OutMachineFriendly,
	}
}

type CommandError struct {
	Context     CommandletContext
	ReasonTag   string
	Description string
}

func (ce *CommandError) Error() string {
	return fmt.Sprintf("command failed with reason[%s] : %s", ce.ReasonTag, ce.Description)
}

func (ce *CommandError) PrintError() {
	if ce.Context.OutQuiet {
		return
	} else if ce.Context.OutMachineFriendly {
		fmt.Fprintf(
			ce.Context.OutDest,
			`{"error":"%s","error_description":"%s"}`, // TODO use appropriate deserializer
			ce.ReasonTag,
			ce.Description,
		)
	} else {
		fmt.Fprintf(
			ce.Context.OutDest,
			`Execution failed.

     Reason: %s
Description: %s
`,
			ce.ReasonTag,
			ce.Description,
		)
	}
}

type Commandlet interface {
	Name() string
	Desc() string
	UsageTmpl() string
	FlagSet() *flag.FlagSet
	Exec(cc CommandletContext) error
}

func ParseFlags(cc CommandletContext, cl Commandlet) (remainder []string, err error) {
	f := cl.FlagSet()
	f.SetOutput(cc.OutDest)
	if err := f.Parse(cc.Args); err != nil {
		return []string{}, &CommandShowUsageError{
			cc,
			err.Error(),
		}
	}
	remainder = f.Args()
	return
}

type CommandShowUsageError struct {
	Context     CommandletContext
	Instruction string
}

func (e *CommandShowUsageError) Error() string {
	if e.Instruction == "" {
		return "invalid command argument(s)"
	} else {
		return e.Instruction
	}
}

type ParentCommandlet struct {
	SubCommands []Commandlet
}

func (c *ParentCommandlet) Exec(cc CommandletContext) error {
	if len(cc.Args) < 1 {
		return &CommandShowUsageError{
			cc,
			"please specify sub command",
		}
	}
	subCmd := cc.Args[0]
	subArgs := cc.Args[1:]
	subCmds := make(map[string]Commandlet)
	for _, s := range c.SubCommands {
		subCmds[s.Name()] = s
	}

	if sc, ok := subCmds[subCmd]; ok {
		scc := NewCommandletContext(
			cc.Cmd+" "+subCmd,
			subArgs,
			sc,
			cc,
		)
		return sc.Exec(scc)

	} else {
		return &CommandShowUsageError{
			cc,
			fmt.Sprintf("Invalid command [%s]", subCmd),
		}
	}
}

func (c *ParentCommandlet) FlagSet() (f *flag.FlagSet) {
	return nil
}

func (c *ParentCommandlet) UsageTmpl() string {
	u := `
Usage: {{.Command}} COMMAND

Available commands:
`

	for _, s := range c.SubCommands {
		u += fmt.Sprintf("  %-10s %s\n", s.Name(), s.Desc())
	}

	u += `

Run '{{.Command}} COMMAND' for more information on a command.
`

	return u
}

type RootCommandlet struct {
	*ParentCommandlet
}

func (c *RootCommandlet) Name() string {
	return "tbx"
}
func (c *RootCommandlet) Desc() string {
	return "Dropbox tools"
}
