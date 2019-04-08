package cmd_dev_auth

import (
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/app"
	"github.com/watermint/toolbox/app/app_zap"
	"github.com/watermint/toolbox/cmd"
	"go.uber.org/zap"
)

type CmdDevAuthAppKey struct {
	*cmd.SimpleCommandlet
}

func (z *CmdDevAuthAppKey) Name() string {
	return "appkey"
}

func (z *CmdDevAuthAppKey) Desc() string {
	return "cmd.dev.auth.appkey.desc"
}

func (z *CmdDevAuthAppKey) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdDevAuthAppKey) FlagConfig(f *flag.FlagSet) {
}

func (z *CmdDevAuthAppKey) Exec(args []string) {
	kb, err := app_zap.Unzap(z.ExecContext)
	if err != nil {
		z.Log().Error("Unable to load", zap.Error(err))
		z.ExecContext.Fatal(app.FatalNoAppKey)
		return
	}
	var keys map[string]string
	err = json.Unmarshal(kb, &keys)
	if err != nil {
		z.Log().Error("Unable load", zap.Error(err), zap.ByteString("seq", kb))
		z.ExecContext.Fatal(app.FatalNoAppKey)
		return
	}
}
