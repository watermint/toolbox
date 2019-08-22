package cmd_bulk

import (
	"flag"
	"github.com/watermint/toolbox/domain/sequence"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"go.uber.org/zap"
)

type CmdTeamBulkVerify struct {
	*cmd2.SimpleCommandlet
	optSeqFile string
}

func (z *CmdTeamBulkVerify) Name() string {
	return "verify"
}

func (z *CmdTeamBulkVerify) Desc() string {
	return "cmd.team.bulk.verify.desc"
}

func (z *CmdTeamBulkVerify) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdTeamBulkVerify) FlagConfig(f *flag.FlagSet) {
	descSeqFile := "Sequence file"
	f.StringVar(&z.optSeqFile, "seq-file", "", descSeqFile)
}

func (z *CmdTeamBulkVerify) Exec(args []string) {
	seq := sequence.New(z.ExecContext)
	if err := seq.Load(z.optSeqFile); err != nil {
		z.Log().Error("One or more invalid line found", zap.Error(err))
	}
}
