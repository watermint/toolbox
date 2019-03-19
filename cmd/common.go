package cmd

import (
	"errors"
	"flag"
	"fmt"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_ui"
	"go.uber.org/zap"
	"sort"
	"strings"
)

type CommandUsage struct {
	Command string
}

type Commandlet interface {
	Name() string
	Desc() string
	Usage() func(CommandUsage)
	FlagConfig(f *flag.FlagSet)
	Exec(args []string)
	Init(parent Commandlet)
	Setup(ec *app.ExecContext)
	Parent() Commandlet
	Log() *zap.Logger
	DefaultErrorHandler(err error) bool
	IsGroup() bool
	IsHidden() bool
}

type CommandletBase struct {
	Commandlet
}

func (*CommandletBase) PrintUsage(ec *app.ExecContext, clt Commandlet) {
	var c Commandlet
	cmds := make([]string, 0)
	c = clt
	for c != nil {
		cmds = append(cmds, c.Name())
		c = c.Parent()
	}

	// reverse array
	chainSize := len(cmds) - 1
	for i := len(cmds)/2 - 1; i >= 0; i-- {
		cmds[i], cmds[chainSize-i] = cmds[chainSize-i], cmds[i]
	}
	cmd := strings.Join(cmds, " ")
	p := struct {
		Command string
	}{
		Command: cmd,
	}

	ec.Msg("cmd.common.base.usage.head").WithData(p).Tell()
	if clt.Usage() == nil {
		ec.Msg("cmd.common.base.usage.default").WithData(p).Tell()
	} else {
		clt.Usage()(p)
	}
}

type SimpleCommandlet struct {
	*CommandletBase
	parent      Commandlet
	logger      *zap.Logger
	ExecContext *app.ExecContext
}

func (z *SimpleCommandlet) IsHidden() bool {
	return false
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

func (z *SimpleCommandlet) DefaultErrorHandler(err error) bool {
	if err != nil {
		return true
	}

	z.Log().Error("Default error handler caught an error", zap.Error(err))
	errorQueue = append(errorQueue, err)
	addError(err)
	return false
}

func (z *SimpleCommandlet) DefaultErrorHandlerIgnoreError(err error) bool {
	if err != nil {
		return true
	}

	z.Log().Error("Default error handler caught an error", zap.Error(err))
	addError(err)
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
	Hidden      bool
}

func (z *CommandletGroup) IsHidden() bool {
	return z.Hidden
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

func (z *CommandletGroup) Usage() func(CommandUsage) {
	f := func(c CommandUsage) {
		z.ExecContext.Msg("cmd.common.group.usage.head").WithData(c).Tell()
		cmds := make(map[string]string)
		names := make([]string, 0)
		cmdLen := 0
		for _, s := range z.SubCommands {
			cmds[s.Name()] = z.ExecContext.Msg(s.Desc()).T()
			names = append(names, s.Name())
			if cmdLen < len(s.Name()) {
				cmdLen = len(s.Name())
			}
		}
		sort.Strings(names)
		subCmdFmt := fmt.Sprintf("  %%-%ds  %%s", cmdLen+1)
		for _, n := range names {
			desc, e := cmds[n]
			if !e {
				continue
			}
			t := fmt.Sprintf(subCmdFmt, n, desc)
			tm := app_ui.NewTextMessage(t, z.ExecContext.UI(), z.Log())
			tm.Tell()
		}
		app_ui.NewTextMessage("\n\n", z.ExecContext.UI(), z.Log()).Tell()
		z.ExecContext.Msg("cmd.common.group.usage.tail").WithData(c).Tell()
	}

	return f
}

func (z *CommandletGroup) FlagConfig(f *flag.FlagSet) {
	z.flags = f
}

func (z *CommandletGroup) Exec(args []string) {
	if len(args) < 1 {
		z.PrintUsage(z.ExecContext, z)
		z.ExecContext.Shutdown()
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
		sc.Setup(z.ExecContext)
		sc.FlagConfig(z.flags)
		if err := z.flags.Parse(subArgs); err != nil {
			z.Log().Error("Command ParseModel error", zap.Error(err))
			z.PrintUsage(z.ExecContext, z)
			return
		}
		remainders := z.flags.Args()
		if len(remainders) > 0 && remainders[0] == "help" {
			z.PrintUsage(z.ExecContext, sc)

			z.ExecContext.Msg("cmd.common.group.usage.options").Tell()
			z.flags.PrintDefaults()
			return
		}

		if !sc.IsGroup() {
			if err := z.ExecContext.ApplyFlags(); err != nil {
				return
			}
		}
		sc.Exec(remainders)
		if !sc.IsGroup() {
			z.ExecContext.Shutdown()
		}
		return
	}

	err := errors.New(fmt.Sprintf("invalid command [%s]", subCmd))
	z.PrintUsage(z.ExecContext, z)
	addError(err)
}

func (z *CommandletGroup) DefaultErrorHandler(err error) bool {
	if err != nil {
		return true
	}

	z.Log().Error("Default error handler caught an error", zap.Error(err))
	errorQueue = append(errorQueue, err)
	addError(err)
	return false
}

var (
	errorQueue = make([]error, 0)
)

func ErrorQueue() []error {
	return errorQueue
}

func addError(lastErr error) {
	errorQueue = append(errorQueue, lastErr)
}
