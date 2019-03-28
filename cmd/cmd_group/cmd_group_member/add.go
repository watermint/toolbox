package cmd_group_member

import (
	"flag"
	"github.com/watermint/toolbox/app/app_io"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/model/mo_group"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"strings"
)

type CmdGroupMemberAdd struct {
	*cmd.SimpleCommandlet
	report app_report.Factory
	optCsv string
}

func (z *CmdGroupMemberAdd) Name() string {
	return "add"
}

func (z *CmdGroupMemberAdd) Desc() string {
	return "cmd.group.member.add.desc"
}

func (z *CmdGroupMemberAdd) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdGroupMemberAdd) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descCsv := z.ExecContext.Msg("cmd.group.member.add.flag.csv").T()
	f.StringVar(&z.optCsv, "csv", "", descCsv)
}

func (z *CmdGroupMemberAdd) Exec(args []string) {
	if z.optCsv == "" {
		z.ExecContext.Msg("cmd.group.member.add.err.not_enough_flag").TellError()
		return
	}
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessManagement())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	groupsByName := make(map[string]*mo_group.Group)
	groups, err := sv_group.New(ctx).List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}
	for _, group := range groups {
		groupsByName[strings.ToLower(group.GroupName)] = group
	}

	type Report struct {
		Email        string `json:"email"`
		TeamMemberId string `json:"team_member_id"`
		DisplayName  string `json:"display_name"`
		GroupId      string `json:"group_id"`
		GroupName    string `json:"group_name"`
		Result       string `json:"result"`
		Reason       string `json:"reason"`
	}

	err = app_io.NewCsvLoader(z.ExecContext, z.optCsv).OnRow(func(cols []string) error {
		if len(cols) < 2 {
			z.Log().Warn("Skip: No data in row.")
			return nil
		}

		email := strings.TrimSpace(cols[0])
		groupNames := make([]string, 0)
		for _, name := range strings.Split(cols[1], ";") {
			groupNames = append(groupNames, strings.TrimSpace(strings.ToLower(name)))
		}

		member, err := sv_member.New(ctx).ResolveByEmail(email)
		if err != nil {
			z.report.Report(&Report{
				Email:  email,
				Result: "failure",
				Reason: ctx.ErrorMsg(err).T(),
			})
			ctx.ErrorMsg(err).TellError()
			return nil
		}

		for _, name := range groupNames {
			if g, e := groupsByName[name]; !e {
				msg := z.ExecContext.Msg("cmd.group.member.add.reason.group_not_found").WithData(struct {
					GroupName string
				}{
					GroupName: name,
				})
				z.report.Report(&Report{
					Email:        email,
					TeamMemberId: member.TeamMemberId,
					DisplayName:  member.DisplayName,
					Result:       "failure",
					Reason:       msg.T(),
				})
				msg.TellError()
				continue
			} else {
				_, err := sv_group_member.New(ctx, g).Add([]string{member.TeamMemberId})
				if err != nil {
					switch {
					case strings.HasPrefix(api_util.ErrorSummary(err), "duplicate_user"):
						z.report.Report(&Report{
							Email:        email,
							TeamMemberId: member.TeamMemberId,
							DisplayName:  member.DisplayName,
							GroupId:      g.GroupId,
							GroupName:    g.GroupName,
							Result:       "skip",
							Reason:       ctx.ErrorMsg(err).T(),
						})

					default:
						z.report.Report(&Report{
							Email:        email,
							TeamMemberId: member.TeamMemberId,
							DisplayName:  member.DisplayName,
							GroupId:      g.GroupId,
							GroupName:    g.GroupName,
							Result:       "failure",
							Reason:       ctx.ErrorMsg(err).T(),
						})
						ctx.ErrorMsg(err).TellError()

					}
				} else {
					z.report.Report(&Report{
						Email:        email,
						TeamMemberId: member.TeamMemberId,
						DisplayName:  member.DisplayName,
						GroupId:      g.GroupId,
						GroupName:    g.GroupName,
						Result:       "success",
					})
				}
			}
		}
		return nil
	}).Load()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
	}
}
