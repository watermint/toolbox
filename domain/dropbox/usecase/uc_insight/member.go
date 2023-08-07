package uc_insight

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedfolder_member"
)

func ConvertSameTeam(sameTeam string) string {
	switch sameTeam {
	case "true":
		return "yes"
	case "false":
		return "no"
	default:
		return ""
	}
}

func IsExternalGroup(member mo_sharedfolder_member.Member) bool {
	if _, ok := member.Group(); ok {
		if mo_sharedfolder_member.IsSameTeam(member.SameTeam()) {
			return true
		}
	}
	return false
}
