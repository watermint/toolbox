package cmd_namespace_member

import (
	"flag"
	"github.com/watermint/toolbox/app/app_report"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/domain/infra/api_auth_impl"
	"github.com/watermint/toolbox/domain/model/mo_namespace"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"go.uber.org/zap"
)

type CmdTeamNamespaceMemberList struct {
	*cmd.SimpleCommandlet

	report         app_report.Factory
	optExpandGroup bool
}

func (CmdTeamNamespaceMemberList) Name() string {
	return "list"
}

func (CmdTeamNamespaceMemberList) Desc() string {
	return "cmd.team.namespace.member.list.desc"
}

func (CmdTeamNamespaceMemberList) Usage() func(cmd.CommandUsage) {
	return nil
}

func (z *CmdTeamNamespaceMemberList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	//descExpandGroup := z.ExecContext.Msg("cmd.team.namespace.member.list.flag.expand_group").T()
	//f.BoolVar(&z.optExpandGroup, "expand-group", false, descExpandGroup)
}

func (z *CmdTeamNamespaceMemberList) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svNamespace := sv_namespace.New(ctx)
	namespaces, err := svNamespace.List()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}

	svTeam := sv_profile.NewTeam(ctx)
	admin, err := svTeam.Admin()
	if err != nil {
		ctx.ErrorMsg(err).TellError()
		return
	}
	adminCtx := ctx.AsAdminId(admin.TeamMemberId)

	for _, namespace := range namespaces {
		if namespace.NamespaceType != "team_folder" &&
			namespace.NamespaceType != "shared_folder" {
			z.Log().Debug("Skip", zap.String("NamespaceId", namespace.NamespaceId), zap.String("type", namespace.NamespaceType), zap.String("name", namespace.Name), zap.String("teamMemberId", namespace.TeamMemberId))
			continue
		}
		svMember := sv_sharedfolder_member.NewBySharedFolderId(adminCtx, namespace.NamespaceId)
		members, err := svMember.List()
		if err != nil {
			adminCtx.ErrorMsg(err).TellError()
			return
		}

		for _, member := range members {
			nm := mo_namespace.NewNamespaceMember(namespace, member)
			z.report.Report(nm)
		}
	}
}
