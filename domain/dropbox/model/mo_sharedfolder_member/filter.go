package mo_sharedfolder_member

import (
	"github.com/watermint/toolbox/domain/common/model/mo_filter"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
)

type FolderMemberFilter interface {
	mo_filter.FilterOpt

	// Set team members to identify external/internal.
	SetMembers(member []*mo_member.Member)
}
