package mo_group

import "encoding/json"

type Group struct {
	Raw                 json.RawMessage
	GroupName           string `path:"group_name"`
	GroupId             string `path:"group_id"`
	GroupManagementType string `path:"group_management_type.\\.tag"`
	GroupExternalId     string `path:"group_external_id"`
	MemberCount         int    `path:"member_count"`
}
