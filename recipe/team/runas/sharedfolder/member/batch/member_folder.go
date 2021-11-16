package batch

type AddMember struct {
	MemberEmail  string `json:"member_email"`
	Path         string `json:"path"`
	AccessLevel  string `json:"access_level"`
	GroupOrEmail string `json:"group_or_email"`
}

type DeleteMember struct {
	MemberEmail  string `json:"member_email"`
	Path         string `json:"path"`
	GroupOrEmail string `json:"group_or_email"`
}
