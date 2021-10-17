package mo_member

type MemberSelector struct {
	Tag       string `json:".tag"`
	DropboxId string `json:"dropbox_id,omitempty"`
	Email     string `json:"email,omitempty"`
}

func NewMemberSelectorDropboxId(id string) *MemberSelector {
	return &MemberSelector{
		Tag:       ".dropbox_id",
		DropboxId: id,
		Email:     "",
	}
}

func NewMemberSelectorEmail(email string) *MemberSelector {
	return &MemberSelector{
		Tag:       ".email",
		DropboxId: "",
		Email:     email,
	}
}
