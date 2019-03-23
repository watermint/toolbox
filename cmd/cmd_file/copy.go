package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/app/app_assert"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_relocation"
)

type CmdFileCopy struct {
	*cmd.SimpleCommandlet
}

func (z *CmdFileCopy) Name() string {
	return "copy"
}

func (z *CmdFileCopy) Desc() string {
	return "cmd.file.copy.desc"
}

func (z *CmdFileCopy) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdFileCopy) FlagConfig(f *flag.FlagSet) {
}

func (z *CmdFileCopy) Exec(args []string) {
	if msg, err := app_assert.AssertArgs(z.ExecContext, args, app_assert.MinNArgs(2)); err != nil {
		msg.TellError()
		return
	}
	n := len(args)
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.Full())
	if err != nil {
		return
	}

	uc := uc_file_relocation.New(ctx)
	src := args[:n-1]
	dst := args[n-1]

	for _, s := range src {
		msg, err := uc.Copy(mo_path.NewPath(s), mo_path.NewPath(dst))
		if err != nil {
			msg.TellError()
			continue
		}
	}
}
