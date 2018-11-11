package cmd_namespace_member

import (
	"flag"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/cmd"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_group"
	"github.com/watermint/toolbox/dbx_api/dbx_namespace"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_sharing"
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
	return "List all namespace members of the team"
}

func (CmdTeamNamespaceMemberList) Usage() string {
	return ""
}

func (c *CmdTeamNamespaceMemberList) FlagConfig(f *flag.FlagSet) {
	c.report.FlagConfig(f)

	descExpandGroup := "Expand group into members"
	f.BoolVar(&c.optExpandGroup, "expand-group", false, descExpandGroup)
}

func (c *CmdTeamNamespaceMemberList) Exec(args []string) {
	apiFile, err := c.ExecContext.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	admin, ea, _ := dbx_profile.AuthenticatedAdmin(apiFile)
	if ea.IsFailure() {
		c.DefaultErrorHandler(ea)
		return
	}
	c.report.Init(c.Log())
	defer c.report.Close()

	if c.optExpandGroup {
		c.groupMembers = dbx_group.GroupMembers(apiFile, c.Log(), c.DefaultErrorHandler)
		if c.groupMembers == nil {
			c.Log().Warn("Unable to list group members")
			return
		}
	}

	l := dbx_namespace.NamespaceList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(namespace *dbx_namespace.Namespace) bool {
			if namespace.NamespaceType != "shared_folder" &&
				namespace.NamespaceType != "team_folder" {
				return true
			}

			sl := dbx_sharing.SharedFolderMembers{
				AsAdminId: admin.TeamMemberId,
				OnError:   c.DefaultErrorHandler,
				OnUser: func(user *dbx_sharing.MembershipUser) bool {
					nu := &dbx_namespace.NamespaceUser{
						Namespace: namespace,
						User:      user,
					}
					c.report.Report(nu)
					return true
				},
				OnGroup: func(group *dbx_sharing.MembershipGroup) bool {
					if c.optExpandGroup {
						if gmm, ok := c.groupMembers[group.Group.GroupId]; ok {
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
								c.report.Report(nu)
							}
						} else {
							c.Log().Warn(
								"Could not expand group",
								zap.String("group_id", group.Group.GroupId),
								zap.String("group_name", group.Group.GroupName),
							)
							ng := &dbx_namespace.NamespaceGroup{
								Namespace: namespace,
								Group:     group,
							}
							c.report.Report(ng)
						}
					} else {
						ng := &dbx_namespace.NamespaceGroup{
							Namespace: namespace,
							Group:     group,
						}
						c.report.Report(ng)
					}
					return true
				},
				OnInvitee: func(invitee *dbx_sharing.MembershipInvitee) bool {
					ni := &dbx_namespace.NamespaceInvitee{
						Namespace: namespace,
						Invitee:   invitee,
					}
					c.report.Report(ni)
					return true
				},
			}
			sl.List(apiFile, namespace.NamespaceId)
			return true
		},
	}
	l.List(apiFile)
}
