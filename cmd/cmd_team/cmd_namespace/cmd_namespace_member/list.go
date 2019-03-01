package cmd_namespace_member

import (
	"flag"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/model/dbx_api"
	"github.com/watermint/toolbox/model/dbx_auth"
	"github.com/watermint/toolbox/model/dbx_group"
	"github.com/watermint/toolbox/model/dbx_namespace"
	"github.com/watermint/toolbox/model/dbx_profile"
	"github.com/watermint/toolbox/model/dbx_sharing"
	"github.com/watermint/toolbox/report"
	"go.uber.org/zap"
)

type CmdTeamNamespaceMemberList struct {
	*cmd.SimpleCommandlet

	apiContext     *dbx_api.Context
	report         report.Factory
	groupMembers   map[string][]*dbx_group.GroupMember
	optExpandGroup bool
}

func (CmdTeamNamespaceMemberList) Name() string {
	return "list"
}

func (CmdTeamNamespaceMemberList) Desc() string {
	return "cmd.team.namespace.member.list.desc"
}

func (CmdTeamNamespaceMemberList) Usage() string {
	return ""
}

func (z *CmdTeamNamespaceMemberList) FlagConfig(f *flag.FlagSet) {
	z.report.ExecContext = z.ExecContext
	z.report.FlagConfig(f)

	descExpandGroup := z.ExecContext.Msg("cmd.team.namespace.member.list.flag.expand_group").Text()
	f.BoolVar(&z.optExpandGroup, "expand-group", false, descExpandGroup)
}

func (z *CmdTeamNamespaceMemberList) Exec(args []string) {
	au := dbx_auth.NewDefaultAuth(z.ExecContext)
	apiFile, err := au.Auth(dbx_auth.DropboxTokenBusinessFile)
	if err != nil {
		return
	}

	admin, ea, _ := dbx_profile.AuthenticatedAdmin(apiFile)
	if ea.IsFailure() {
		z.DefaultErrorHandler(ea)
		return
	}
	z.report.Init(z.ExecContext)
	defer z.report.Close()

	if z.optExpandGroup {
		z.groupMembers = dbx_group.GroupMembers(apiFile, z.Log(), z.DefaultErrorHandler)
		if z.groupMembers == nil {
			z.ExecContext.Msg("cmd.team.namespace.member.list.err.fail_expand_group").TellError()
			z.Log().Warn("Unable to list group members")
			return
		}
	}

	l := dbx_namespace.NamespaceList{
		OnError: z.DefaultErrorHandler,
		OnEntry: func(namespace *dbx_namespace.Namespace) bool {
			if namespace.NamespaceType != "shared_folder" &&
				namespace.NamespaceType != "team_folder" {
				return true
			}

			sl := dbx_sharing.SharedFolderMembers{
				AsAdminId: admin.TeamMemberId,
				OnError:   z.DefaultErrorHandler,
				OnUser: func(user *dbx_sharing.MembershipUser) bool {
					nu := &dbx_namespace.NamespaceUser{
						Namespace: namespace,
						User:      user,
					}
					z.report.Report(nu)
					return true
				},
				OnGroup: func(group *dbx_sharing.MembershipGroup) bool {
					if z.optExpandGroup {
						if gmm, ok := z.groupMembers[group.Group.GroupId]; ok {
							for _, gm := range gmm {
								nu := &dbx_namespace.NamespaceUser{
									Namespace: namespace,
									User: &dbx_sharing.MembershipUser{
										Membership: group.Membership,
										User: &dbx_sharing.User{
											UserAccountId: gm.Profile.AccountId,
											Email:         gm.Profile.Email,
											DisplayName:   gjson.Get(string(gm.Profile.Profile), "name.display_name").String(),
											SameTeam:      true,
											TeamMemberId:  gm.TeamMemberId,
										},
									},
								}
								z.report.Report(nu)
							}
						} else {
							z.ExecContext.Msg("cmd.team.namespace.member.list.err.cant_expand").WithData(struct {
								Id   string
								Name string
							}{
								Id:   group.Group.GroupId,
								Name: group.Group.GroupName,
							}).TellError()
							z.Log().Warn(
								"Could not expand group",
								zap.String("group_id", group.Group.GroupId),
								zap.String("group_name", group.Group.GroupName),
							)
							ng := &dbx_namespace.NamespaceGroup{
								Namespace: namespace,
								Group:     group,
							}
							z.report.Report(ng)
						}
					} else {
						ng := &dbx_namespace.NamespaceGroup{
							Namespace: namespace,
							Group:     group,
						}
						z.report.Report(ng)
					}
					return true
				},
				OnInvitee: func(invitee *dbx_sharing.MembershipInvitee) bool {
					ni := &dbx_namespace.NamespaceInvitee{
						Namespace: namespace,
						Invitee:   invitee,
					}
					z.report.Report(ni)
					return true
				},
			}
			sl.List(apiFile, namespace.NamespaceId)
			return true
		},
	}
	l.List(apiFile)
}
