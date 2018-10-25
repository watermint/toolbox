package cmd_namespace_member

import (
	"flag"
	"github.com/cihub/seelog"
	"github.com/tidwall/gjson"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_sharing"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/infra"
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
	Namespace *dbx_team.Namespace         `json:"namespace"`
	User      *dbx_sharing.MembershipUser `json:"user"`
}

type NamespaceGroup struct {
	Namespace *dbx_team.Namespace          `json:"namespace"`
	Group     *dbx_sharing.MembershipGroup `json:"group"`
}

type NamespaceInvitee struct {
	Namespace *dbx_team.Namespace            `json:"namespace"`
	Invitee   *dbx_sharing.MembershipInvitee `json:"invitee"`
}

func (c *CmdTeamNamespaceMemberList) Exec(ec *infra.ExecContext, args []string) {
	if err := ec.Startup(); err != nil {
		return
	}
	defer ec.Shutdown()

	apiFile, err := ec.LoadOrAuthBusinessFile()
	if err != nil {
		return
	}

	admin, ea, _ := dbx_profile.AuthenticatedAdmin(apiFile)
	if ea.IsFailure() {
		cmdlet.DefaultErrorHandler(ea)
		return
	}
	c.report.Open()
	defer c.report.Close()

	if c.optExpandGroup {
		if !c.expandGroup(apiFile) {
			seelog.Warnf("Unable to list group members")
			return
		}
	}

	l := dbx_team.NamespaceList{
		OnError: cmdlet.DefaultErrorHandler,
		OnEntry: func(namespace *dbx_team.Namespace) bool {
			if namespace.NamespaceType != "shared_folder" &&
				namespace.NamespaceType != "team_folder" {
				return true
			}

			sl := dbx_sharing.SharedFolderMembers{
				AsAdminId: admin.TeamMemberId,
				OnError:   cmdlet.DefaultErrorHandler,
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
							seelog.Warnf("Could not expand group[id={%s} name={%s}]", group.Group.GroupId, group.Group.GroupName)
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

	seelog.Debugf("Expand group")
	gl := dbx_team.GroupList{
		OnError: cmdlet.DefaultErrorHandler,
		OnEntry: func(group *dbx_team.Group) bool {
			seelog.Debugf("Group[%s] Name[%s]", group.GroupId, group.GroupName)

			gml := dbx_team.GroupMemberList{
				OnError: cmdlet.DefaultErrorHandler,
				OnEntry: func(gm *dbx_team.GroupMember) bool {

					if g, ok := c.groups[group.GroupId]; ok {
						g = append(g, gm)
						c.groups[group.GroupId] = g
						seelog.Debugf("Group[%s] GroupMembers[%d]", group.GroupId, len(g))
					} else {
						g = make([]*dbx_team.GroupMember, 1)
						g[0] = gm
						c.groups[group.GroupId] = g
						seelog.Debugf("Group[%s] GroupMembers[%d]", group.GroupId, len(g))
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
		seelog.Debugf("GroupId[%s] GroupCount[%d]", k, len(v))
	}

	return true
}
