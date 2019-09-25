package cmd_sharedlink

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"github.com/watermint/toolbox/legacy/app/app_report_legacy"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
)

type CmdSharedLinkCreate struct {
	*cmd2.SimpleCommandlet
	report      app_report_legacy.Factory
	optTeamOnly bool
	optPassword string
	optExpires  string
}

func (z *CmdSharedLinkCreate) Name() string {
	return "create"
}

func (z *CmdSharedLinkCreate) Desc() string {
	return "cmd.sharedlink.create.desc"
}

func (z *CmdSharedLinkCreate) Usage() func(cmd2.CommandUsage) {
	return nil
}

func (z *CmdSharedLinkCreate) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descTeamOnly := z.ExecContext.Msg("cmd.sharedlink.create.flag.team_only").T()
	f.BoolVar(&z.optTeamOnly, "team-only", false, descTeamOnly)

	descPassword := z.ExecContext.Msg("cmd.sharedlink.create.flag.password").T()
	f.StringVar(&z.optPassword, "password", "", descPassword)

	descExpires := z.ExecContext.Msg("cmd.sharedlink.create.flag.expires").T()
	f.StringVar(&z.optExpires, "expires", "", descExpires)
}

func (z *CmdSharedLinkCreate) Exec(args []string) {
	opts := make([]sv_sharedlink.LinkOpt, 0)
	if z.optExpires != "" {
		if expires, e := ut_time.ParseTimestamp(z.optExpires); e {
			opts = append(opts, sv_sharedlink.Expires(expires))
		} else {
			z.ExecContext.Msg("cmd.sharedlink.create.err.unsupported_time_format").WithData(struct {
				Time string
			}{
				Time: z.optExpires,
			}).TellError()
			return
		}
	}
	if z.optTeamOnly {
		opts = append(opts, sv_sharedlink.TeamOnly())
	}
	if z.optPassword != "" {
		opts = append(opts, sv_sharedlink.Password(z.optPassword))
	}
	if len(args) < 1 {
		z.ExecContext.Msg("cmd.sharedlink.create.err.not_enough_argument").TellError()
		return
	}

	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.Full())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_sharedlink.New(ctx)

	for _, p := range args {
		link, err := svc.Create(mo_path.NewPath(p), opts...)
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			continue
		}
		z.report.Report(link)
	}
}
