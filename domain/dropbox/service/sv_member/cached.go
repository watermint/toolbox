package sv_member

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"strings"
)

type cachedMember struct {
	impl    Member
	members []*mo_member.Member
}

func (z *cachedMember) Suspend(member *mo_member.Member, opts ...SuspendOpt) (err error) {
	return z.impl.Suspend(member, opts...)
}

func (z *cachedMember) Unsuspend(member *mo_member.Member) (err error) {
	return z.impl.Unsuspend(member)
}

func (z *cachedMember) UpdateVisibility(email string, visible bool) (updated *mo_member.Member, err error) {
	return z.impl.UpdateVisibility(email, visible)
}

func (z *cachedMember) ListEach(f func(member *mo_member.Member) bool, opts ...ListOpt) error {
	return z.ListEach(f, opts...)
}

func (z *cachedMember) Update(member *mo_member.Member, opts ...UpdateOpt) (updated *mo_member.Member, err error) {
	//	z.members = nil // invalidate
	return z.impl.Update(member, opts...)
}

func (z *cachedMember) List(opt ...ListOpt) (members []*mo_member.Member, err error) {
	if z.members == nil {
		z.members, err = z.impl.List(opt...)
		if err != nil {
			return nil, err
		}
	}
	return z.members, nil
}

func (z *cachedMember) Resolve(teamMemberId string) (member *mo_member.Member, err error) {
	if z.members == nil {
		z.members, err = z.impl.List()
		if err != nil {
			return nil, err
		}
	}
	for _, m := range z.members {
		if m.TeamMemberId == teamMemberId {
			return m, nil
		}
	}
	return nil, ErrorMemberNotFoundForTeamMemberId
}

func (z *cachedMember) ResolveByEmail(email string) (member *mo_member.Member, err error) {
	if z.members == nil {
		z.members, err = z.impl.List()
		if err != nil {
			return nil, err
		}
	}
	em := strings.ToLower(email)
	for _, m := range z.members {
		if strings.ToLower(m.Email) == em {
			return m, nil
		}
	}
	return nil, ErrorMemberNotFoundForEmail
}

func (z *cachedMember) Add(email string, opts ...AddOpt) (member *mo_member.Member, err error) {
	z.members = nil // invalidate cache
	return z.impl.Add(email, opts...)
}

func (z *cachedMember) Remove(member *mo_member.Member, opts ...RemoveOpt) (err error) {
	z.members = nil // invalidate cache
	return z.impl.Remove(member, opts...)
}
