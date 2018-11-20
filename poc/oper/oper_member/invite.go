package oper_member

import (
	"github.com/watermint/toolbox/poc/oper"
	"github.com/watermint/toolbox/poc/oper/oper_auth"
	"github.com/watermint/toolbox/poc/oper/oper_io"
)

type Invite struct {
	oper.OperationBase
	OptApi          *oper_auth.DropboxBusinessManagement
	OptCsvFile      *oper_io.MustInputFile
	OptReportFormat string `default:"json"`
	OptSilent       bool   `default:"false"`
}

func (z *Invite) Exec() {
	z.Tell(z.Message("start_invitation"))

	//z.UI.Tell(z.Res.Msg("Start invitation"))
	//z.Log.Info("CSV: ", zap.String("file", z.OptCsvFile.Path))
	//z.Log.Info("ReportFormat", zap.String("format", z.OptReportFormat))
	//z.Log.Info("Silent: ", zap.Bool("silent", z.OptSilent))
	//z.UI.TellDone(z.Res.Msg("Done"))
}
