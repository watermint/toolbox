package mo_namespace

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_file"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/essentials/api/api_parser"
	"github.com/watermint/toolbox/essentials/log/esl"
)

type Namespace struct {
	Raw           json.RawMessage
	Name          string `path:"name" json:"name"`
	NamespaceId   string `path:"namespace_id" json:"namespace_id" gorm:"primaryKey"`
	NamespaceType string `path:"namespace_type.\\.tag" json:"namespace_type"`
	TeamMemberId  string `path:"team_member_id" json:"team_member_id"`
}

func NewNamespaceMember(namespace *Namespace, member mo_sharedfolder_member.Member) (nm *NamespaceMember) {
	raws := make(map[string]json.RawMessage)
	raws["namespace"] = namespace.Raw
	raws["member"] = member.EntryRaw()
	raw := api_parser.CombineRaw(raws)

	nm = &NamespaceMember{}
	if err := api_parser.ParseModelRaw(nm, raw); err != nil {
		esl.Default().Error("unable to parse", esl.Error(err))
	}
	return nm
}

type NamespaceMember struct {
	Raw                   json.RawMessage
	NamespaceName         string `path:"namespace.name" json:"namespace_name"`
	NamespaceId           string `path:"namespace.namespace_id" json:"namespace_id"`
	NamespaceType         string `path:"namespace.namespace_type.\\.tag" json:"namespace_type"`
	NamespaceTeamMemberId string `path:"namespace.team_member_id" json:"namespace_team_member_id"`
	EntryAccessType       string `path:"member.access_type.\\.tag" json:"entry_access_type"`
	EntryIsInherited      bool   `path:"member.is_inherited" json:"entry_is_inherited"`
	AccountId             string `path:"member.user.account_id" json:"account_id"`
	TeamMemberId          string `path:"member.user.team_member_id" json:"team_member_id"`
	Email                 string `path:"member.user.email" json:"email"`
	DisplayName           string `path:"member.user.display_name" json:"display_name"`
	GroupName             string `path:"member.group.group_name" json:"group_name"`
	GroupId               string `path:"member.group.group_id" json:"group_id"`
	InviteeEmail          string `path:"member.invitee.email" json:"invitee_email"`
}

func (z *NamespaceMember) Namespace() (namespace *Namespace) {
	namespace = &Namespace{}
	if err := api_parser.ParseModelPathRaw(namespace, z.Raw, "namespace"); err != nil {
		esl.Default().Warn("unexpected data format", esl.String("entry", string(z.Raw)), esl.Error(err))
		// return empty
		return namespace
	}
	return namespace
}

func (z *NamespaceMember) Member() (member mo_sharedfolder_member.Member) {
	member = &mo_sharedfolder_member.Metadata{}
	if err := api_parser.ParseModelPathRaw(member, z.Raw, "member"); err != nil {
		esl.Default().Warn("unexpected data format", esl.String("entry", string(z.Raw)), esl.Error(err))
		// return empty
		return member
	}
	return member
}

type NamespaceEntry struct {
	Raw                  json.RawMessage
	NamespaceType        string `path:"namespace.namespace_type.\\.tag" json:"namespace_type"`
	NamespaceId          string `path:"namespace.namespace_id" json:"namespace_id"`
	NamespaceName        string `path:"namespace.name" json:"namespace_name"`
	NamespaceMemberEmail string `json:"namespace_member_email"`
	Id                   string `path:"entry.id" json:"file_id"`
	Tag                  string `path:"entry.\\.tag" json:"tag"`
	Name                 string `path:"entry.name" json:"name"`
	PathDisplay          string `path:"entry.path_display" json:"path_display"`
	ClientModified       string `path:"entry.client_modified" json:"client_modified"`
	ServerModified       string `path:"entry.server_modified" json:"server_modified"`
	Revision             string `path:"entry.rev" json:"revision"`
	Size                 int64  `path:"entry.size" json:"size"`
	ContentHash          string `path:"entry.content_hash" json:"content_hash"`
	SharedFolderId       string `path:"entry.sharing_info.shared_folder_id" json:"shared_folder_id"`
	ParentSharedFolderId string `path:"entry.sharing_info.parent_shared_folder_id" json:"parent_shared_folder_id"`
}

func NewNamespaceEntry(namespace *Namespace, entry *mo_file.ConcreteEntry) (ne *NamespaceEntry) {
	raws := make(map[string]json.RawMessage)
	raws["namespace"] = namespace.Raw
	raws["entry"] = entry.Raw
	raw := api_parser.CombineRaw(raws)

	ne = &NamespaceEntry{}
	if err := api_parser.ParseModelRaw(ne, raw); err != nil {
		esl.Default().Error("unable to parse", esl.Error(err))
	}
	return ne
}
