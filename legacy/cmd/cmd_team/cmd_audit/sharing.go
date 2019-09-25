package cmd_audit

import (
	"flag"
	"github.com/watermint/toolbox/domain/model/mo_namespace"
	"github.com/watermint/toolbox/domain/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_group"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"github.com/watermint/toolbox/domain/service/sv_member"
	"github.com/watermint/toolbox/domain/service/sv_namespace"
	"github.com/watermint/toolbox/domain/service/sv_profile"
	"github.com/watermint/toolbox/domain/service/sv_sharedfolder_member"
	"github.com/watermint/toolbox/domain/service/sv_sharedlink"
	"github.com/watermint/toolbox/domain/service/sv_team"
	"github.com/watermint/toolbox/infra/api/api_auth_impl"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/legacy/app/app_report_legacy"
	cmd2 "github.com/watermint/toolbox/legacy/cmd"
	"go.uber.org/zap"
)

type CmdTeamAuditSharing struct {
	*cmd2.SimpleCommandlet
	report app_report_legacy.Factory
}

func (CmdTeamAuditSharing) Name() string {
	return "sharing"
}

func (CmdTeamAuditSharing) Desc() string {
	return "cmd.team.audit.sharing.desc"
}

func (z *CmdTeamAuditSharing) Usage() func(usage cmd2.CommandUsage) {
	return nil
}

func (z *CmdTeamAuditSharing) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)
}

func (z *CmdTeamAuditSharing) Exec(args []string) {
	ctx, err := api_auth_impl.Auth(z.ExecContext, api_auth_impl.BusinessFile())
	if err != nil {
		return
	}

	z.report.Init(z.ExecContext)
	defer z.report.Close()

	svp := sv_profile.NewTeam(ctx)
	svm := sv_member.New(ctx)
	svg := sv_group.New(ctx)
	svt := sv_team.New(ctx)
	svn := sv_namespace.New(ctx)

	// Identify admin
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.identify_admin").Tell()
	admin, err := svp.Admin()
	if err != nil {
		z.DefaultErrorHandler(err)
		return
	}
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.do_as_admin").WithData(struct {
		Email string
	}{
		Email: admin.Email,
	}).Tell()
	z.Log().Info("Execute scan as admin", zap.String("email", admin.Email))

	// Scan team info
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.team_info").Tell()
	z.Log().Info("Scanning Team Info")
	{
		info, err := svt.Info()
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
		z.report.Report(info)
	}

	// Scan team feature
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.team_feature").Tell()
	z.Log().Info("Scanning Team Feature")
	{
		features, err := svt.Feature()
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
		z.report.Report(features)
	}

	// Scan Team members
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.members").Tell()
	z.Log().Info("Scanning Team Members")
	members, err := svm.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	for _, member := range members {
		z.report.Report(member)
	}

	// Scan shared links
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.shared_link").Tell()
	z.Log().Info("Scanning Shared links")
	for _, member := range members {
		svl := sv_sharedlink.New(ctx.AsMemberId(member.TeamMemberId))
		links, err := svl.List()
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
		for _, link := range links {
			lm := mo_sharedlink.NewSharedLinkMember(link, member)
			z.report.Report(lm)
		}
	}

	// Scan Team group
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.groups").Tell()
	z.Log().Info("Scanning Team Group")
	groups, err := svg.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	for _, group := range groups {
		z.report.Report(group)
	}

	// Scan Team group members
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.group_members").Tell()
	z.Log().Info("Scanning Team Group Member")
	for _, group := range groups {
		svgm := sv_group_member.New(ctx, group)
		gms, err := svgm.List()
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
		for _, gm := range gms {
			z.report.Report(gm)
		}
	}

	// Scan namespaces
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.namespace").Tell()
	z.Log().Info("Scanning Namespace")
	namespaces, err := svn.List()
	if err != nil {
		api_util.UIMsgFromError(err).TellError()
		return
	}
	for _, namespace := range namespaces {
		z.report.Report(namespace)
	}

	// Scan namespace members
	z.ExecContext.Msg("cmd.team.audit.sharing.progress.namespace_members").Tell()
	z.Log().Info("Scanning Namespace members")
	for _, namespace := range namespaces {
		if namespace.NamespaceType != "team_folder" &&
			namespace.NamespaceType != "shared_folder" {
			continue
		}
		cta := ctx.AsAdminId(admin.TeamMemberId)
		svnm := sv_sharedfolder_member.NewBySharedFolderId(cta, namespace.NamespaceId)
		mems, err := svnm.List()
		if err != nil {
			api_util.UIMsgFromError(err).TellError()
			return
		}
		for _, mem := range mems {
			nm := mo_namespace.NewNamespaceMember(namespace, mem)
			z.report.Report(nm)
		}
	}
}
