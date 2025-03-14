package mo_user

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
)

type Feature struct {
	PaperAsFiles       bool `json:"paper_as_files"`
	FileLocking        bool `json:"file_locking"`
	TeamSharedDropbox  bool `json:"team_shared_dropbox"`
	DistinctMemberHome bool `json:"distinct_member_home"`
}

func (z Feature) FileSystemType() dbx_filesystem.TeamFileSystemType {
	return dbx_filesystem.IdentifyFileSystemTypeByParam(z.DistinctMemberHome, z.TeamSharedDropbox)
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
