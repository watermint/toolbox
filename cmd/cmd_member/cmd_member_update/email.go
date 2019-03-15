package cmd_member_update

import (
	"flag"
	"github.com/watermint/toolbox/app/app_io"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/model/dbx_auth"
	"go.uber.org/zap"
)

type CmdMemberUpdateEmail struct {
	*cmd.SimpleCommandlet
	optCsv string
	report app_report.Factory

	// email address mapping. key is for existing email, value is for new address
	emailMapping map[string]string
}

func (*CmdMemberUpdateEmail) Name() string {
	return "email"
}

func (*CmdMemberUpdateEmail) Desc() string {
	return "cmd.member.update.email.desc"
}

func (z *CmdMemberUpdateEmail) Usage() func(cmd.CommandUsage) {
	return func(usage cmd.CommandUsage) {
		z.ExecContext.Msg("cmd.member.update.email.usage").WithData(usage).Tell()
	}
}

func (z *CmdMemberUpdateEmail) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descCsv := z.ExecContext.Msg("cmd.member.update.email.flag.csv").T()
	f.StringVar(&z.optCsv, "csv", "", descCsv)
}

func (z *CmdMemberUpdateEmail) loadMapping() error {
	z.emailMapping = make(map[string]string)
	loader := app_io.NewCsvLoader(z.ExecContext, z.optCsv).
		OnRow(func(cols []string) error {
			// skip
			if len(cols) < 2 {
				z.Log().Debug("Skip", zap.Strings("cols", cols))
				return nil
			}

			from := cols[0]
			to := cols[1]
			if !api_util.RegexEmail.MatchString(from) {
				z.Log().Debug("`from` email doesn't match to the pattern", zap.Strings("cols", cols))
				return nil
			}
			if !api_util.RegexEmail.MatchString(to) {
				z.Log().Debug("`to` email doesn't match to the pattern", zap.Strings("cols", cols))
				return nil
			}
			z.emailMapping[from] = to
			return nil
		})
	return loader.Load()
}

func (z *CmdMemberUpdateEmail) Exec(args []string) {
	if err := z.loadMapping(); err != nil {
		return
	}

	ctx, err := api_auth_impl.Auth(z.ExecContext, dbx_auth.DropboxTokenBusinessManagement)
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svc := sv_member.New(ctx)
	members, err := svc.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	memberDic := make(map[string]*mo_member.Member)
	for _, m := range members {
		memberDic[m.Email] = m
	}

	type Report struct {
		FromEmail string `json:"from_email"`
		ToEmail   string `json:"to_email"`
		Result    string `json:"result"`
		Reason    string `json:"reason"`
	}

	for f, t := range z.emailMapping {
		r := Report{
			FromEmail: f,
			ToEmail:   t,
		}

		m, e := memberDic[f]
		if !e {
			z.Log().Debug("member not found", zap.String("from_email", f))
			z.ExecContext.Msg("cmd.member.update.email.err.member_not_found").WithData(struct {
				Email string
			}{
				Email: f,
			}).TellFailure()
			r.Result = z.ExecContext.Msg("cmd.member.update.email.report.result.failure").T()
			r.Reason = z.ExecContext.Msg("cmd.member.update.email.report.reason.member_not_found").T()
			z.report.Report(r)
			continue
		}

		m.Email = t

		_, err := svc.Update(m)
		if err != nil {
			z.Log().Debug("can't update email", zap.String("from_email", f), zap.String("to_email", t), zap.Error(err))
			z.ExecContext.Msg("cmd.member.update.email.err.cant_update").WithData(struct {
				From  string
				To    string
				Error string
			}{
				From:  f,
				To:    t,
				Error: ctx.ErrorMsg(err).T(),
			}).TellFailure()
			r.Result = z.ExecContext.Msg("cmd.member.update.email.report.result.failure").T()
			r.Reason = ctx.ErrorMsg(err).T()
			z.report.Report(r)
			continue
		}

		z.ExecContext.Msg("cmd.member.update.email.progress.updated").WithData(struct {
			From string
			To   string
		}{
			From: f,
			To:   t,
		}).TellSuccess()

		r.Result = z.ExecContext.Msg("cmd.member.update.email.report.result.success").T()
		z.report.Report(r)
	}
}
