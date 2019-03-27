package cmd_dev_testsuite

import (
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/usecase/uc_team_migration"
	"go.uber.org/zap"
	"io/ioutil"
	"os"
)

type CmdDevTestSuiteMigration struct {
	*cmd.SimpleCommandlet
	report       app_report.Factory
	optActorFile string
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

	descActorFile := z.ExecContext.Msg("cmd.dev.testsuite.migration.flags.actor_file").T()
	f.StringVar(&z.optActorFile, "actor-file", "", descActorFile)
}

func (z *CmdDevTestSuiteMigration) Exec(args []string) {
	f, err := os.Open(z.optActorFile)
	if err != nil {
		z.Log().Warn("Unable to open file", zap.String("file", z.optActorFile), zap.Error(err))
		return
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		z.Log().Warn("Unable to read file", zap.String("file", z.optActorFile), zap.Error(err))
		return
	}
	err = f.Close()
	if err != nil {
		z.Log().Warn("Unable to close file", zap.String("file", z.optActorFile), zap.Error(err))
		return
	}
	actors := &uc_team_migration.Actors{}
	err = json.Unmarshal(b, actors)
	if err != nil {
		z.Log().Warn("Unable to unmarshal", zap.String("file", z.optActorFile), zap.Error(err))
		return
	}

	scenario := uc_team_migration.NewScenario(z.ExecContext, actors)
	if err := scenario.Auth(); err != nil {
		z.Log().Warn("Authentication failed", zap.Error(err))
		return
	}
	if err := scenario.Cleanup(); err != nil {
		z.Log().Warn("Clean up failed", zap.Error(err))
	}
}
