package cmd

import (
	"errors"
	"flag"
	"fmt"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_util"
	"github.com/watermint/toolbox/model/dbx_api"
	"go.uber.org/zap"
	"strings"
)

type Commandlet interface {
	Name() string
	Desc() string
	Usage() string
	FlagConfig(f *flag.FlagSet)
	Exec(args []string)
	Init(parent Commandlet)
	Setup(ec *app.ExecContext)
	Parent() Commandlet
	Log() *zap.Logger
	DefaultErrorHandler(ea dbx_api.ErrorAnnotation) bool
}

type CommandletBase struct {
	Commandlet
}

func (*CommandletBase) PrintUsage(clt Commandlet) {
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

	usage, tmplErr := app_util.CompileTemplate(tmpl,
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
	parent      Commandlet
	logger      *zap.Logger
	ExecContext *app.ExecContext
}

func (c *SimpleCommandlet) Parent() Commandlet {
	return c.parent
}

func (c *SimpleCommandlet) Init(parent Commandlet) {
	c.parent = parent
}

func (c *SimpleCommandlet) Setup(ec *app.ExecContext) {
	c.ExecContext = ec
}

func (c *SimpleCommandlet) Log() *zap.Logger {
	return c.ExecContext.Log()
}

func (c *SimpleCommandlet) DefaultErrorHandler(ea dbx_api.ErrorAnnotation) bool {
	if ea.IsSuccess() {
		return true
	}

	c.Log().Error("Default error handler caught an error",
		zap.String("error_type", ea.ErrorTypeLabel()),
		zap.String("error_message", ea.UserMessage()),
	)
	errorQueue = append(errorQueue, ea)
	addError(ea)
	return false
}

func (c *SimpleCommandlet) DefaultErrorHandlerIgnoreError(ea dbx_api.ErrorAnnotation) bool {
	if ea.IsSuccess() {
		return true
	}

	c.Log().Error("Default error handler caught an error",
		zap.String("error_type", ea.ErrorTypeLabel()),
		zap.String("error_message", ea.UserMessage()),
	)
	addError(ea)
	return true
}

type CommandletGroup struct {
	*CommandletBase
	flagset     *flag.FlagSet
	parent      Commandlet
	logger      *zap.Logger
	SubCommands []Commandlet

	ExecContext *app.ExecContext
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

func (c *CommandletGroup) Setup(ec *app.ExecContext) {
	c.ExecContext = ec
}

func (c *CommandletGroup) Log() *zap.Logger {
	return c.ExecContext.Log()
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

func (c *CommandletGroup) Exec(args []string) {
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
			c.Log().Error("Command ParseModel error", zap.Error(err))
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

		c.ExecContext.ApplyFlags()
		defer c.ExecContext.Shutdown()
		sc.Setup(c.ExecContext)
		sc.Exec(remainders)
		return
	}

	err := errors.New(fmt.Sprintf("invalid command [%s]", subCmd))
	ea := dbx_api.ErrorAnnotation{
		ErrorType: dbx_api.ErrorBadInputParam,
		Error:     err,
	}
	c.PrintUsage(c)
	addError(ea)
}

func (c *CommandletGroup) DefaultErrorHandler(ea dbx_api.ErrorAnnotation) bool {
	if ea.IsSuccess() {
		return true
	}

	c.Log().Error("Default error handler caught an error",
		zap.String("error_type", ea.ErrorTypeLabel()),
		zap.String("error_message", ea.UserMessage()),
	)
	errorQueue = append(errorQueue, ea)
	addError(ea)
	return false
}

var (
	errorQueue = make([]dbx_api.ErrorAnnotation, 0)
)

func ErrorQueue() []dbx_api.ErrorAnnotation {
	return errorQueue
}

func addError(ea dbx_api.ErrorAnnotation) {
	errorQueue = append(errorQueue, ea)
}
