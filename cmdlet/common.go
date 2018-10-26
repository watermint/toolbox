package cmdlet

import (
	"errors"
	"flag"
	"fmt"
	"github.com/cihub/seelog"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/infra"
	"github.com/watermint/toolbox/infra/util"
	"strings"
)

type Commandlet interface {
	Name() string
	Desc() string
	Usage() string
	FlagConfig(f *flag.FlagSet)
	Exec(ec *infra.ExecContext, args []string)
	Init(parent Commandlet)
	Parent() Commandlet
}

type CommandletBase struct {
	Commandlet
}

func (*CommandletBase) PrintUsage(clt Commandlet) {
	seelog.Flush()

	var c Commandlet
	cmds := make([]string, 0)
	c = clt
	for c != nil {
		cmds = append(cmds, c.Name())
		c = c.Parent()
	}
	tmpl := clt.Usage()
	if tmpl == "" {
		tmpl = "{{.Command}} [Options]"
	}

	chainSize := len(cmds) - 1
	for i := len(cmds)/2 - 1; i >= 0; i-- {
		cmds[i], cmds[chainSize-i] = cmds[chainSize-i], cmds[i]
	}
	cmd := strings.Join(cmds, " ")

	usage, tmplErr := util.CompileTemplate(tmpl,
		struct {
			Command string
		}{
			Command: cmd,
		})
	if tmplErr != nil {
		panic(tmplErr)
	}

	fmt.Printf("Usage:\n\n%s\n\n", usage)
}

type SimpleCommandlet struct {
	*CommandletBase
	parent Commandlet
}

func (c *SimpleCommandlet) Parent() Commandlet {
	return c.parent
}

func (c *SimpleCommandlet) Init(parent Commandlet) {
	c.parent = parent
}

type CommandletGroup struct {
	*CommandletBase
	flagset     *flag.FlagSet
	parent      Commandlet
	SubCommands []Commandlet

	CommandName string
	CommandDesc string
}

func (c *CommandletGroup) Name() string {
	return c.CommandName
}
func (c *CommandletGroup) Desc() string {
	return c.CommandDesc
}
func (c *CommandletGroup) Parent() Commandlet {
	return c.parent
}

func (c *CommandletGroup) Init(parent Commandlet) {
	c.parent = parent
}

func (c *CommandletGroup) Usage() string {
	u := `{{.Command}} COMMAND

Available commmands:
`
	for _, s := range c.SubCommands {
		u += fmt.Sprintf("  %-10s %s\n", s.Name(), s.Desc())
	}

	u += `

Run '{{.Command}} COMMAND help' for more information on a command.
`
	return u
}

func (c *CommandletGroup) FlagConfig(f *flag.FlagSet) {
	c.flagset = f
}

func (c *CommandletGroup) Exec(ec *infra.ExecContext, args []string) {
	if len(args) < 1 {
		c.PrintUsage(c)
		return
	}

	subCmd := args[0]
	subArgs := args[1:]
	subCmds := make(map[string]Commandlet)
	for _, s := range c.SubCommands {
		subCmds[s.Name()] = s
	}
	if sc, ok := subCmds[subCmd]; ok {
		sc.Init(c)
		sc.FlagConfig(c.flagset)
		if err := c.flagset.Parse(subArgs); err != nil {
			seelog.Errorf("Command Parse error %s", err)
			c.PrintUsage(c)
			return
		}
		remainders := c.flagset.Args()
		if len(remainders) > 0 && remainders[0] == "help" {
			c.PrintUsage(sc)

			fmt.Println("Available options:")
			c.flagset.PrintDefaults()
			return
		}

		sc.Exec(ec, remainders)
		return
	}

	err := errors.New(fmt.Sprintf("invalid command [%s]", subCmd))
	ea := dbx_api.ErrorAnnotation{
		ErrorType: dbx_api.ErrorBadInputParam,
		Error:     err,
	}
	c.PrintUsage(c)
	DefaultErrorHandler(ea)
}

var (
	errorQueue = make([]dbx_api.ErrorAnnotation, 0)
)

func ErrorQueue() []dbx_api.ErrorAnnotation {
	return errorQueue
}

func DefaultErrorHandler(ea dbx_api.ErrorAnnotation) bool {
	if ea.IsSuccess() {
		return true
	}

	seelog.Errorf("Error: ErrorType[%s] UserMessage[%s]",
		ea.ErrorTypeLabel(),
		ea.UserMessage(),
	)
	errorQueue = append(errorQueue, ea)
	return false
}

func DefaultErrorHandlerIgnoreError(ea dbx_api.ErrorAnnotation) bool {
	if ea.IsSuccess() {
		return true
	}

	seelog.Warnf("Error: ErrorType[%s] UserMessage[%s]",
		ea.ErrorTypeLabel(),
		ea.UserMessage(),
	)
	errorQueue = append(errorQueue, ea)
	return true
}
