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
	IsGroup() bool
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

func (z *SimpleCommandlet) IsGroup() bool {
	return false
}

func (z *SimpleCommandlet) Parent() Commandlet {
	return z.parent
}

func (z *SimpleCommandlet) Init(parent Commandlet) {
	z.parent = parent
}

func (z *SimpleCommandlet) Setup(ec *app.ExecContext) {
	z.ExecContext = ec
}

func (z *SimpleCommandlet) Log() *zap.Logger {
	return z.ExecContext.Log()
}

func (z *SimpleCommandlet) DefaultErrorHandler(ea dbx_api.ErrorAnnotation) bool {
	if ea.IsSuccess() {
		return true
	}

	z.Log().Error("Default error handler caught an error",
		zap.String("error_type", ea.ErrorTypeLabel()),
		zap.String("error_message", ea.UserMessage()),
	)
	errorQueue = append(errorQueue, ea)
	addError(ea)
	return false
}

func (z *SimpleCommandlet) DefaultErrorHandlerIgnoreError(ea dbx_api.ErrorAnnotation) bool {
	if ea.IsSuccess() {
		return true
	}

	z.Log().Error("Default error handler caught an error",
		zap.String("error_type", ea.ErrorTypeLabel()),
		zap.String("error_message", ea.UserMessage()),
	)
	addError(ea)
	return true
}

type CommandletGroup struct {
	*CommandletBase
	flags       *flag.FlagSet
	parent      Commandlet
	logger      *zap.Logger
	SubCommands []Commandlet

	ExecContext *app.ExecContext
	CommandName string
	CommandDesc string
}

func (z *CommandletGroup) IsGroup() bool {
	return true
}

func (z *CommandletGroup) Name() string {
	return z.CommandName
}
func (z *CommandletGroup) Desc() string {
	return z.CommandDesc
}
func (z *CommandletGroup) Parent() Commandlet {
	return z.parent
}

func (z *CommandletGroup) Init(parent Commandlet) {
	z.parent = parent
}

func (z *CommandletGroup) Setup(ec *app.ExecContext) {
	z.ExecContext = ec
}

func (z *CommandletGroup) Log() *zap.Logger {
	return z.ExecContext.Log()
}

func (z *CommandletGroup) Usage() string {
	u := `{{.Command}} COMMAND

Available commmands:
`
	for _, s := range z.SubCommands {
		u += fmt.Sprintf("  %-10s %s\n", s.Name(), s.Desc())
	}

	u += `

Run '{{.Command}} COMMAND help' for more information on a command.
`
	return u
}

func (z *CommandletGroup) FlagConfig(f *flag.FlagSet) {
	z.flags = f
}

func (z *CommandletGroup) Exec(args []string) {
	if len(args) < 1 {
		z.PrintUsage(z)
		return
	}

	subCmd := args[0]
	subArgs := args[1:]
	subCmds := make(map[string]Commandlet)
	for _, s := range z.SubCommands {
		subCmds[s.Name()] = s
	}
	if sc, ok := subCmds[subCmd]; ok {
		sc.Init(z)
		sc.FlagConfig(z.flags)
		if err := z.flags.Parse(subArgs); err != nil {
			z.Log().Error("Command ParseModel error", zap.Error(err))
			z.PrintUsage(z)
			return
		}
		remainders := z.flags.Args()
		if len(remainders) > 0 && remainders[0] == "help" {
			z.PrintUsage(sc)

			fmt.Println("Available options:")
			z.flags.PrintDefaults()
			return
		}

		if !sc.IsGroup() {
			z.ExecContext.ApplyFlags()
		}
		sc.Setup(z.ExecContext)
		sc.Exec(remainders)
		if !sc.IsGroup() {
			z.ExecContext.Shutdown()
		}
		return
	}

	err := errors.New(fmt.Sprintf("invalid command [%s]", subCmd))
	ea := dbx_api.ErrorAnnotation{
		ErrorType: dbx_api.ErrorBadInputParam,
		Error:     err,
	}
	z.PrintUsage(z)
	addError(ea)
}

func (z *CommandletGroup) DefaultErrorHandler(ea dbx_api.ErrorAnnotation) bool {
	if ea.IsSuccess() {
		return true
	}

	z.Log().Error("Default error handler caught an error",
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
