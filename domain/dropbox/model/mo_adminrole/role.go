package mo_adminrole

type Role struct {
	RoleId      string `path:"role_id" json:"role_id"`
	Name        string `path:"name" json:"name"`
	Description string `path:"description" json:"description"`
}

type MemberRole struct {
	TeamMemberId string `path:"team_member_id" json:"team_member_id"`
	Email        string `path:"email" json:"email"`
	RoleId       string `path:"role_id" json:"role_id"`
	Name         string `path:"name" json:"name"`
	Description  string `path:"description" json:"description"`
}

func NewMemberRole(teamMemberId, email string, role *Role) *MemberRole {
	return &MemberRole{
		TeamMemberId: teamMemberId,
		Email:        email,
		RoleId:       role.RoleId,
		Name:         role.Name,
		Description:  role.Description,
	}
}
