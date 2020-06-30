package uc_team_content

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgScan struct {
	ProgressScanNamespaceMetadata app_msg.Message
	ProgressScanNamespaceMember   app_msg.Message
	ProgressScanMember            app_msg.Message
}

var (
	ErrorTeamOwnedNamespaceIsNotInitialized = errors.New("team owned namespace is not initialized")
	MScanMetadata                           = app_msg.Apply(&MsgScan{}).(*MsgScan)
)

type ScanNamespace interface {
	Scan(ctl app_control.Control, ctx dbx_context.Context, namespaceName string, namespaceId string)
}

type ScanNamespaceMetadataAndMembership struct {
	Metadata   ScanNamespace
	Membership ScanNamespace
}

func (z *ScanNamespaceMetadataAndMembership) Scan(ctl app_control.Control, ctx dbx_context.Context, namespaceName string, namespaceId string) {
	z.Membership.Scan(ctl, ctx, namespaceName, namespaceId)
	z.Metadata.Scan(ctl, ctx, namespaceName, namespaceId)
}

type Membership struct {
	NamespaceId   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	Path          string `json:"path"`
	FolderType    string `json:"folder_type"`
	OwnerTeamId   string `json:"owner_team_id"`
	OwnerTeamName string `json:"owner_team_name"`
	AccessType    string `json:"access_type"`
	MemberType    string `json:"member_type"`
	MemberId      string `json:"member_id"`
	MemberName    string `json:"member_name"`
	MemberEmail   string `json:"member_email"`
	SameTeam      string `json:"same_team"`
}

type NoMember struct {
	OwnerTeamId   string `json:"owner_team_id"`
	OwnerTeamName string `json:"owner_team_name"`
	NamespaceId   string `json:"namespace_id"`
	NamespaceName string `json:"namespace_name"`
	Path          string `json:"path"`
	FolderType    string `json:"folder_type"`
}

type FolderPolicy struct {
	NamespaceId        string `json:"namespace_id"`
	NamespaceName      string `json:"namespace_name"`
	Path               string `json:"path"`
	IsTeamFolder       bool   `json:"is_team_folder"`
	OwnerTeamId        string `json:"owner_team_id"`
	OwnerTeamName      string `json:"owner_team_name"`
	PolicyManageAccess string `json:"policy_manage_access"`
	PolicySharedLink   string `json:"policy_shared_link"`
	PolicyMember       string `json:"policy_member"`
	PolicyViewerInfo   string `json:"policy_viewer_info"`
}
