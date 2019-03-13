package cmd_device

import (
	"bufio"
	"encoding/json"
	"flag"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_device"
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

	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	if z.optStdin {
		z.ExecContext.Log().Debug("Read from STDIN")
		z.unlink(apiFile, os.Stdin)

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

		z.unlink(apiFile, f)
	}
}

func (z *CmdTeamDeviceUnlink) unlink(c *dbx_api.Context, records io.Reader) {
	type UnlinkRecord struct {
		Tag          string `json:"tag"`
		TeamMemberId string `json:"team_member_id"`
		SessionId    string `json:"session_id"`
	}

	br := bufio.NewReaderSize(records, 8192)
	revoke := dbx_device.RevokeSession{
		OnError: z.DefaultErrorHandler,
	}

	for {
		line, _, err := br.ReadLine()
		if err == io.EOF {
			z.ExecContext.Log().Debug("reached to EOF")
			return
		}

		ur := UnlinkRecord{}
		err = json.Unmarshal([]byte(line), &ur)
		if err != nil {
			z.ExecContext.Log().Debug("Unable to unmarshal", zap.Error(err))
			continue
		}

		z.ExecContext.Log().Debug("tyring unlink", zap.Any("record", ur))
		z.ExecContext.Msg("cmd.team.device.unlink.progress.record").WithData(ur).Tell()

		switch ur.Tag {
		case "web_session":
			revoke.WebSession(c, ur.TeamMemberId, ur.SessionId)

		case "desktop_client":
			revoke.DesktopClient(c, ur.TeamMemberId, ur.SessionId, z.optDeleteOnUnlink)

		case "mobile_client":
			revoke.MobileClient(c, ur.TeamMemberId, ur.SessionId)

		default:
			z.ExecContext.Log().Warn("Invalid device type found", zap.String("type", ur.Tag))
		}
	}
}
