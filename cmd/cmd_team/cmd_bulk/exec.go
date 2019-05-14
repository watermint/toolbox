package cmd_bulk

import (
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/sequence"
	"go.uber.org/zap"
)

type CmdTeamBulkExec struct {
	*cmd.SimpleCommandlet
	optSeqFile   string
	optRetryable bool
	optMaxRetry  int
}

func (z *CmdTeamBulkExec) Name() string {
	return "exec"
}

func (z *CmdTeamBulkExec) Desc() string {
	return "cmd.team.bulk.exec.desc"
}

func (z *CmdTeamBulkExec) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamBulkExec) FlagConfig(f *flag.FlagSet) {
	descSeqFile := "Sequence file"
	f.StringVar(&z.optSeqFile, "seq-file", "", descSeqFile)

	descRetryable := "Mark task sequence retryable"
	f.BoolVar(&z.optRetryable, "retryable", false, descRetryable)

	descMaxRetry := "Maximum number of retry"
	f.IntVar(&z.optMaxRetry, "max-retry", 10, descMaxRetry)
}

func (z *CmdTeamBulkExec) Exec(args []string) {
	opts := make([]sequence.RunOpt, 0)
	if z.optRetryable {
		opts = append(opts, sequence.Retryable())
	}
	if z.optMaxRetry > 0 {
		opts = append(opts, sequence.MaxRetry(z.optMaxRetry))
	}
	seq := sequence.New(z.ExecContext)
	if err := seq.Load(z.optSeqFile); err != nil {
		z.Log().Error("One or more invalid line found", zap.Error(err))
	}
	if err := seq.Run(opts...); err != nil {
		z.Log().Error("One or more task failed", zap.Error(err))
	}
}
