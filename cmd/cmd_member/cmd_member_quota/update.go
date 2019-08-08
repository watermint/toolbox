package cmd_member_quota

import (
	"flag"
	"github.com/watermint/toolbox/app/app_io"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_member_quota"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_member_quota"
	"go.uber.org/zap"
)

type CmdMemberQuotaUpdate struct {
	*cmd.SimpleCommandlet
	report  app_report.Factory
	optSize int
	optCsv  string
}

func (z *CmdMemberQuotaUpdate) Name() string {
	return "update"
}

func (z *CmdMemberQuotaUpdate) Desc() string {
	return "cmd.member.quota.update.desc"
}

func (z *CmdMemberQuotaUpdate) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdMemberQuotaUpdate) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descSize := z.ExecContext.Msg("cmd.member.quota.update.flag.size").T()
	f.IntVar(&z.optSize, "size", 0, descSize)

	descCsv := z.ExecContext.Msg("cmd.member.quota.update.flag.csv").T()
	f.StringVar(&z.optCsv, "csv", "", descCsv)
}

func (z *CmdMemberQuotaUpdate) Exec(args []string) {
	if z.optSize < 15 {
		z.ExecContext.Msg("cmd.member.quota.update.err.lower_limit").TellError()
		return
	}
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessManagement())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	type Report struct {
		Email       string `json:"email"`
		DisplayName string `json:"display_name"`
		Result      string `json:"result"`
		Reason      string `json:"reason"`
		Quota       int    `json:"quota"`
	}

	setByEmail := func(email string) error {
		l := z.Log().With(zap.String("email", email))
		l.Debug("Resolve")
		svm := sv_member.New(ctx)
		m, err := svm.ResolveByEmail(email)
		if err != nil {
			l.Debug("Failed resolve", zap.Error(err))
			z.report.Report(&Report{
				Email:  email,
				Result: "failure",
				Reason: api_util.UIMsgFromError(err).T(),
			})
			return err
		}
		l = l.With(zap.String("teamMemberId", m.TeamMemberId))
		l.Debug("Update quota")

		svq := sv_member_quota.NewQuota(ctx)
		q, err := svq.Update(&mo_member_quota.Quota{
			TeamMemberId: m.TeamMemberId,
			Quota:        z.optSize,
		})
		if err != nil {
			l.Debug("Unable to update", zap.Error(err))
			z.report.Report(&Report{
				Email:       email,
				DisplayName: m.DisplayName,
				Result:      "failure",
				Reason:      api_util.UIMsgFromError(err).T(),
			})
			return err
		}

		z.report.Report(&Report{
			Email:       email,
			Result:      "success",
			DisplayName: m.DisplayName,
			Quota:       q.Quota,
		})
		return nil
	}

	for _, m := range args {
		setByEmail(m)
	}

	if z.optCsv != "" {
		app_io.NewCsvLoader(z.ExecContext, z.optCsv).OnRow(func(cols []string) error {
			if len(cols) < 1 {
				return nil
			}
			setByEmail(cols[0])
			return nil
		}).Load()
	}
}
