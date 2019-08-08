package cmd_device

import (
	"bufio"
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_context"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_device"
	"github.com/watermint/toolbox/domain/service/sv_device"
	"go.uber.org/zap"
	"io"
	"os"
)

type CmdTeamDeviceUnlink struct {
	*cmd.SimpleCommandlet
	optStdin          bool
	optDeleteOnUnlink bool
	optFile           string
}

func (CmdTeamDeviceUnlink) Name() string {
	return "unlink"
}

func (CmdTeamDeviceUnlink) Desc() string {
	return "cmd.team.device.unlink.desc"
}

func (z *CmdTeamDeviceUnlink) Usage() func(usage cmd.CommandUsage) {
	return func(usage cmd.CommandUsage) {
		z.ExecContext.Msg("cmd.team.device.unlink.usage").WithData(usage).Tell()
	}
}

func (z *CmdTeamDeviceUnlink) FlagConfig(f *flag.FlagSet) {
	descStdin := z.ExecContext.Msg("cmd.team.device.unlink.flag.stdin").T()
	f.BoolVar(&z.optStdin, "stdin", false, descStdin)

	descFile := z.ExecContext.Msg("cmd.team.device.unlink.flag.file").T()
	f.StringVar(&z.optFile, "file", "", descFile)

	descDeleteOnUnlink := z.ExecContext.Msg("cmd.team.device.unlink.flag.delete_on_unlink").T()
	f.BoolVar(&z.optDeleteOnUnlink, "delete-on-unlink", true, descDeleteOnUnlink)
}

func (z *CmdTeamDeviceUnlink) Exec(args []string) {
	if !z.optStdin && z.optFile == "" {
		z.ExecContext.Msg("cmd.team.device.unlink.err.specify_file").TellError()
		return
	}

	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	if z.optStdin {
		z.ExecContext.Log().Debug("Read from STDIN")
		z.unlink(ctx, os.Stdin)

	} else {
		z.ExecContext.Log().Debug("Open file", zap.String("path", z.optFile))
		f, err := os.Open(z.optFile)
		if err != nil {
			z.ExecContext.Msg("cmd.team.device.unlink.err.open_file").WithData(struct {
				File string
			}{
				File: z.optFile,
			}).TellError()

			z.ExecContext.Log().Warn(
				"Unable to open file",
				zap.String("file", z.optFile),
				zap.Error(err),
			)
			return
		}
		defer f.Close()

		z.unlink(ctx, f)
	}
}

func (z *CmdTeamDeviceUnlink) unlink(ctx api_context.Context, records io.Reader) {
	br := bufio.NewReaderSize(records, 8192)
	svd := sv_device.New(ctx)

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			z.ExecContext.Log().Debug("reached to EOF")
			return
		}

		ms := &mo_device.MemberSession{}
		err = json.Unmarshal([]byte(line), &ms)
		if err != nil {
			z.ExecContext.Log().Debug("Unable to unmarshal", zap.Error(err))
			continue
		}

		z.ExecContext.Log().Debug("tyring unlink", zap.Any("record", ms))
		z.ExecContext.Msg("cmd.team.device.unlink.progress.record").WithData(ms).Tell()

		if err := svd.Revoke(ms.Session()); err != nil {
			api_util.UIMsgFromError(err).TellError()
		}
	}
}
