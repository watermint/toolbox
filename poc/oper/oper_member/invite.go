package oper_member

import (
	"github.com/watermint/toolbox/poc/oper/oper_auth"
	"github.com/watermint/toolbox/poc/oper/oper_io"
	"github.com/watermint/toolbox/poc/oper/oper_ui"
	"go.uber.org/zap"
)

type Invite struct {
	UI              oper_ui.UI
	Logger          *zap.Logger
	OptApi          *oper_auth.DropboxBusinessManagement
	OptCsvFile      *oper_io.MustInputFile
	OptReportFormat string
	OptSilent       bool
}

func (z *Invite) Exec() {
	z.Logger.Info("Invite")

	//z.UI.Tell(z.Res.Msg("Start invitation"))
	//z.Log.Info("CSV: ", zap.String("file", z.OptCsvFile.Path))
	//z.Log.Info("ReportFormat", zap.String("format", z.OptReportFormat))
	//z.Log.Info("Silent: ", zap.Bool("silent", z.OptSilent))
	//z.UI.TellDone(z.Res.Msg("Done"))
}
