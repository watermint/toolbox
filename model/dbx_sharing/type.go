package dbx_sharing

import "encoding/json"

type Membership struct {
	AccessType  string          `json:"access_type"`
	Permissions json.RawMessage `json:"permissions,omitempty"`
	IsInherited bool            `json:"is_inherited"`
}

type User struct {
	UserAccountId string `json:"user_account_id,omitempty"`
	Email         string `json:"email"`
	DisplayName   string `json:"display_name,omitempty"`
	SameTeam      bool   `json:"same_team"`
	TeamMemberId  string `json:"team_member_id,omitempty"`
}

type MembershipUser struct {
	Membership *Membership `json:"membership"`
	User       *User       `json:"user"`
}

type Group struct {
	GroupName           string `json:"group_name,omitempty"`
	GroupId             string `json:"group_id"`
	GroupManagementType string `json:"group_management_type,omitempty"`
	GroupType           string `json:"group_type,omitempty"`
	IsMember            bool   `json:"is_member"`
	IsOwner             bool   `json:"is_owner"`
	SameTeam            bool   `json:"same_team"`
	MemberCount         int64  `json:"member_count,omitempty"`
}

type MembershipGroup struct {
	Membership *Membership `json:"membership"`
	Group      *Group      `json:"group"`
}

type Invitee struct {
	Email string `json:"email"`
}

type MembershipInvitee struct {
	Membership *Membership `json:"membership"`
	Invitee    *Invitee    `json:"invitee"`
	User       *User       `json:"user,omitempty"`
}
