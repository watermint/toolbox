package mo_user

type UserSelector struct {
	Tag          string `json:".tag"`
	TeamMemberId string `json:"team_member_id,omitempty"`
	ExternalId   string `json:"external_id,omitempty"`
	Email        string `json:"email,omitempty"`
}

func NewUserSelectorByTeamMemberId(teamMemberId string) UserSelector {
	return UserSelector{
		Tag:          "team_member_id",
		TeamMemberId: teamMemberId,
	}
}

func NewUserSelectorByExternalid(externalid string) UserSelector {
	return UserSelector{
		Tag:        "external_id",
		ExternalId: externalid,
	}
}

func NewUserSelectorByEmail(email string) UserSelector {
	return UserSelector{
		Tag:   "email",
		Email: email,
	}
}
