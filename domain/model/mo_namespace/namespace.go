package mo_namespace

import "encoding/json"

type Namespace struct {
	Raw           json.RawMessage
	Name          string `path:"name"`
	NamespaceId   string `path:"namespace_id"`
	NamespaceType string `path:"namespace_type.\\.tag"`
	TeamMemberId  string `path:"team_member_id"`
}
