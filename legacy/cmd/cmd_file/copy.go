package cmd_file

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/usecase/uc_file_relocation"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_assert"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdFileCopy struct {
	*cmd2.SimpleCommandlet
}

func (z *CmdFileCopy) Name() string {
	return "copy"
}

func (z *CmdFileCopy) Desc() string {
	return "cmd.file.copy.desc"
}

func (z *CmdFileCopy) Usage() func(cmd2.CommandUsage) {
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
		err := uc.Copy(mo_path.NewPath(s), mo_path.NewPath(dst))
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			continue
		}
	}
}
