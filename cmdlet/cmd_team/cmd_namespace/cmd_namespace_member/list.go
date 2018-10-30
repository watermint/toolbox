package cmd_namespace_member

import (
	"flag"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_namespace"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_sharing"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"go.uber.org/zap"
)

type CmdTeamNamespaceMemberList struct {
	*cmdlet.SimpleCommandlet

	apiContext     *dbx_api.Context
	report         cmdlet.Report
	groups         map[string][]*dbx_team.GroupMember
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

type NamespaceUser struct {
	Namespace *dbx_namespace.Namespace    `json:"namespace"`
	User      *dbx_sharing.MembershipUser `json:"user"`
}

type NamespaceGroup struct {
	Namespace *dbx_namespace.Namespace     `json:"namespace"`
	Group     *dbx_sharing.MembershipGroup `json:"group"`
}

type NamespaceInvitee struct {
	Namespace *dbx_namespace.Namespace       `json:"namespace"`
	Invitee   *dbx_sharing.MembershipInvitee `json:"invitee"`
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
	c.report.Open(c)
	defer c.report.Close()

	if c.optExpandGroup {
		if !c.expandGroup(apiFile) {
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
					nu := &NamespaceUser{
						Namespace: namespace,
						User:      user,
					}
					c.report.Report(nu)
					return true
				},
				OnGroup: func(group *dbx_sharing.MembershipGroup) bool {
					if c.optExpandGroup {
						if gmm, ok := c.groups[group.Group.GroupId]; ok {
							for _, gm := range gmm {
								nu := &NamespaceUser{
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
							ng := &NamespaceGroup{
								Namespace: namespace,
								Group:     group,
							}
							c.report.Report(ng)
						}
					} else {
						ng := &NamespaceGroup{
							Namespace: namespace,
							Group:     group,
						}
						c.report.Report(ng)
					}
					return true
				},
				OnInvitee: func(invitee *dbx_sharing.MembershipInvitee) bool {
					ni := &NamespaceInvitee{
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

func (c *CmdTeamNamespaceMemberList) expandGroup(ctx *dbx_api.Context) bool {
	c.groups = make(map[string][]*dbx_team.GroupMember)

	c.Log().Debug("Expand group")
	gl := dbx_team.GroupList{
		OnError: c.DefaultErrorHandler,
		OnEntry: func(group *dbx_team.Group) bool {
			c.Log().Debug("onEntry",
				zap.String("group_id", group.GroupId),
				zap.String("group_name", group.GroupName),
			)

			gml := dbx_team.GroupMemberList{
				OnError: c.DefaultErrorHandler,
				OnEntry: func(gm *dbx_team.GroupMember) bool {

					if g, ok := c.groups[group.GroupId]; ok {
						g = append(g, gm)
						c.groups[group.GroupId] = g

						c.Log().Debug("onEntry",
							zap.String("group_id", group.GroupId),
							zap.Int("group_members", len(g)),
						)
					} else {
						g = make([]*dbx_team.GroupMember, 1)
						g[0] = gm
						c.groups[group.GroupId] = g

						c.Log().Debug("onEntry",
							zap.String("group_id", group.GroupId),
							zap.Int("group_members", len(g)),
						)
					}

					return true
				},
			}
			gml.List(ctx, group)
			return true
		},
	}
	gl.List(ctx)

	for k, v := range c.groups {
		c.Log().Debug("Group summary",
			zap.String("group_id", k),
			zap.Int("member_count", len(v)),
		)
	}

	return true
}
