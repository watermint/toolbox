package mo_user

type Feature struct {
	PaperAsFiles bool `json:"paper_as_files"`
	FileLocking  bool `json:"file_locking"`
}

type MemberFeature struct {
	Email        string `path:"member.profile.email" json:"email"`
	PaperAsFiles bool   `json:"paper_as_files"`
	FileLocking  bool   `json:"file_locking"`
}

func NewMemberFeature(email string, feature *Feature) MemberFeature {
	return MemberFeature{
		Email:        email,
		PaperAsFiles: feature.PaperAsFiles,
		FileLocking:  feature.FileLocking,
	}
}
