package mo_user

import "github.com/watermint/toolbox/domain/dropbox/model/mo_team"

type Feature struct {
	PaperAsFiles       bool `json:"paper_as_files"`
	FileLocking        bool `json:"file_locking"`
	TeamSharedDropbox  bool `json:"team_shared_dropbox"`
	DistinctMemberHome bool `json:"distinct_member_home"`
}

func (z Feature) FileSystemType() mo_team.TeamFileSystemType {
	return mo_team.IdentifyFileSystemTypeByParam(z.DistinctMemberHome, z.TeamSharedDropbox)
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
