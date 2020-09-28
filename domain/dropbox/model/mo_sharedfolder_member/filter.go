package mo_sharedfolder_member

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
)

type FolderMemberFilter interface {
	mo_filter.FilterOpt

	// Set team members to identify external/internal.
	SetMembers(member []*mo_member.Member)
}

type FolderMemberFilterData struct {
	Enabled bool                `json:"enabled"`
	Members []*mo_member.Member `json:"members"`
}
