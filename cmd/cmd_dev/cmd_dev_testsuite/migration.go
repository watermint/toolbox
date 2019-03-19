package cmd_dev_testsuite

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
)

type CmdDevTestSuiteMigration struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
}

func (z *CmdDevTestSuiteMigration) Name() string {
	return "migration"
}

func (z *CmdDevTestSuiteMigration) Desc() string {
	return "cmd.dev.testsuite.migration.desc"
}

func (z *CmdDevTestSuiteMigration) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdDevTestSuiteMigration) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdDevTestSuiteMigration) Exec(args []string) {
}
