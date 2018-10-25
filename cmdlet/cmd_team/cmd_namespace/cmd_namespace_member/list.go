package cmd_namespace_member

import (
	"flag"
	"github.com/watermint/toolbox/cmdlet"
	"github.com/watermint/toolbox/dbx_api"
	"github.com/watermint/toolbox/dbx_api/dbx_profile"
	"github.com/watermint/toolbox/dbx_api/dbx_sharing"
	"github.com/watermint/toolbox/dbx_api/dbx_team"
	"github.com/watermint/toolbox/infra"
)

type CmdTeamNamespaceMemberList struct {
	*cmdlet.SimpleCommandlet

	apiContext *dbx_api.Context
	report     cmdlet.Report
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
					ng := &NamespaceGroup{
						Namespace: namespace,
						Group:     group,
					}
					c.report.Report(ng)
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
