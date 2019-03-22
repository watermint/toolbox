package cmd_member

import (
	"flag"
	"github.com/watermint/toolbox/app/app_io"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/model/dbx_auth"
	"go.uber.org/zap"
	"strings"
)

type CmdMemberInvite struct {
	*cmd.SimpleCommandlet
	optSilent bool
	optCsv    string
	report    app_report.Factory
}

func (z *CmdMemberInvite) Name() string {
	return "invite"
}

func (z *CmdMemberInvite) Desc() string {
	return "cmd.member.invite.desc"
}

func (z *CmdMemberInvite) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdMemberInvite) FlagConfig(f *flag.FlagSet) {
	descSilent := z.ExecContext.Msg("cmd.member.invite.flag.silent").T()
	f.BoolVar(&z.optSilent, "silent", false, descSilent)

	descCsv := z.ExecContext.Msg("cmd.member.invite.flag.csv").T()
	f.StringVar(&z.optCsv, "csv", "", descCsv)

	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdMemberInvite) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessManagement)
	if err != nil {
		return
	}
	svm := sv_member.New(ctx)

	err = app_io.NewCsvLoader(z.ExecContext, z.optCsv).
		OnRow(func(cols []string) error {
			if len(cols) < 1 {
				return nil
			}
			email := strings.TrimSpace(cols[0])
			if !api_util.RegexEmail.MatchString(email) {
				z.Log().Debug("skip: the data is not looking alike an email address", zap.String("email", email))
				return nil
			}
			opts := make([]sv_member.AddOpt, 0)
			if len(cols) >= 2 {
				givenName := cols[1]
				opts = append(opts, sv_member.AddWithGivenName(givenName))
			}
			if len(cols) >= 3 {
				surname := cols[2]
				opts = append(opts, sv_member.AddWithSurname(surname))
			}
			if z.optSilent {
				opts = append(opts, sv_member.AddWithoutSendWelcomeEmail())
			}

			member, err := svm.Add(email, opts...)
			if err != nil {
				ctx.ErrorMsg(err).TellError()
				z.Log().Warn("Unable to invite", zap.String("email", email), zap.Error(err))
				return nil
			}
			z.report.Report(member)

			return nil
		}).
		Load()

	if err != nil {
		z.Log().Debug("Unable to load", zap.Error(err))
	}
}
