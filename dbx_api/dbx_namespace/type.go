package dbx_namespace

import "github.com/watermint/toolbox/dbx_api/dbx_sharing"

type NamespaceUser struct {
	Namespace *Namespace                  `json:"namespace"`
	User      *dbx_sharing.MembershipUser `json:"user"`
}

type NamespaceGroup struct {
	Namespace *Namespace                   `json:"namespace"`
	Group     *dbx_sharing.MembershipGroup `json:"group"`
}

type NamespaceInvitee struct {
	Namespace *Namespace                     `json:"namespace"`
	Invitee   *dbx_sharing.MembershipInvitee `json:"invitee"`
}
