package mo_namespace

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/model/mo_sharedfolder_member"
	"github.com/watermint/toolbox/infra/api/api_parser"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
)

type Namespace struct {
	Raw           json.RawMessage
	Name          string `path:"name"`
	NamespaceId   string `path:"namespace_id"`
	NamespaceType string `path:"namespace_type.\\.tag"`
	TeamMemberId  string `path:"team_member_id"`
}

func NewNamespaceMember(namespace *Namespace, member mo_sharedfolder_member.Member) (nm *NamespaceMember) {
	raws := make(map[string]json.RawMessage)
	raws["namespace"] = namespace.Raw
	raws["member"] = member.EntryRaw()
	raw := api_parser.CombineRaw(raws)

	nm = &NamespaceMember{}
	if err := api_parser.ParseModelRaw(nm, raw); err != nil {
		app_root.Log().Error("unable to parse", zap.Error(err))
	}
	return nm
}

type NamespaceMember struct {
	Raw              json.RawMessage
	NamespaceName    string `path:"namespace.name"`
	NamespaceId      string `path:"namespace.namespace_id"`
	NamespaceType    string `path:"namespace.namespace_type.\\.tag"`
	TeamMemberId     string `path:"namespace.team_member_id"`
	EntryAccessType  string `path:"member.access_type.\\.tag"`
	EntryIsInherited bool   `path:"member.is_inherited"`
	AccountId        string `path:"member.user.account_id"`
	Email            string `path:"member.user.email"`
	DisplayName      string `path:"member.user.display_name"`
	GroupName        string `path:"member.group.group_name"`
	GroupId          string `path:"member.group.group_id"`
	InviteeEmail     string `path:"member.invitee.email"`
}

func (z *NamespaceMember) Namespace() (namespace *Namespace) {
	namespace = &Namespace{}
	if err := api_parser.ParseModelPathRaw(namespace, z.Raw, "namespace"); err != nil {
		app_root.Log().Warn("unexpected data format", zap.String("entry", string(z.Raw)), zap.Error(err))
		// return empty
		return namespace
	}
	return namespace
}

func (z *NamespaceMember) Member() (member mo_sharedfolder_member.Member) {
	member = &mo_sharedfolder_member.Metadata{}
	if err := api_parser.ParseModelPathRaw(member, z.Raw, "member"); err != nil {
		app_root.Log().Warn("unexpected data format", zap.String("entry", string(z.Raw)), zap.Error(err))
		// return empty
		return member
	}
	return member
}
