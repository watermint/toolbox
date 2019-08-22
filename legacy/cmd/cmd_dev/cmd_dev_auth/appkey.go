package cmd_dev_auth

import (
	"encoding/json"
	"flag"
	app2 "github.com/watermint/toolbox/legacy/app"
	"github.com/watermint/toolbox/legacy/app/app_zap"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"go.uber.org/zap"
)

type CmdDevAuthAppKey struct {
	*cmd2.SimpleCommandlet
}

func (z *CmdDevAuthAppKey) Name() string {
	return "appkey"
}

func (z *CmdDevAuthAppKey) Desc() string {
	return "cmd.dev.auth.appkey.desc"
}

func (z *CmdDevAuthAppKey) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdDevAuthAppKey) FlagConfig(f *flag.FlagSet) {
}

func (z *CmdDevAuthAppKey) Exec(args []string) {
	kb, err := app_zap.Unzap(z.ExecContext)
	if err != nil {
		z.Log().Error("Unable to load", zap.Error(err))
		z.ExecContext.Fatal(app2.FatalNoAppKey)
		return
	}
	var keys map[string]string
	err = json.Unmarshal(kb, &keys)
	if err != nil {
		z.Log().Error("Unable load", zap.Error(err), zap.ByteString("seq", kb))
		z.ExecContext.Fatal(app2.FatalNoAppKey)
		return
	}
}
