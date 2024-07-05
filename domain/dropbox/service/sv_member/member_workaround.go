package sv_member

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
)

func NewWorkaround853(client dbx_client.Client) Member {
	return &memberWorkaround853{
		v2: NewV2(client),
		v1: NewListV1(client),
	}
}

// Workaround for the issue #853 when `members/list_v2` API returns empty result.
type memberWorkaround853 struct {
	v2 Member
	v1 Member
}

func (z memberWorkaround853) List(opts ...ListOpt) (members []*mo_member.Member, err error) {
	members, err = z.v2.List(opts...)
	if err != nil {
		return z.v1.List(opts...)
	}
	return
}

func (z memberWorkaround853) ListEach(f func(member *mo_member.Member) bool, opts ...ListOpt) error {
	err := z.v2.ListEach(f, opts...)
	if err != nil {
		return z.v1.ListEach(f, opts...)
	}
	return nil
}

func (z memberWorkaround853) Update(member *mo_member.Member, opts ...UpdateOpt) (updated *mo_member.Member, err error) {
	return z.v2.Update(member, opts...)
}

func (z memberWorkaround853) UpdateVisibility(email string, visible bool) (updated *mo_member.Member, err error) {
	return z.v2.UpdateVisibility(email, visible)
}

func (z memberWorkaround853) Resolve(teamMemberId string) (member *mo_member.Member, err error) {
	return z.v2.Resolve(teamMemberId)
}

func (z memberWorkaround853) ResolveByEmail(email string) (member *mo_member.Member, err error) {
	return z.v2.ResolveByEmail(email)
}

func (z memberWorkaround853) Add(email string, opts ...AddOpt) (member *mo_member.Member, err error) {
	return z.v2.Add(email, opts...)
}

func (z memberWorkaround853) Remove(member *mo_member.Member, opts ...RemoveOpt) (err error) {
	return z.v2.Remove(member, opts...)
}

func (z memberWorkaround853) Suspend(member *mo_member.Member, opts ...SuspendOpt) (err error) {
	return z.v2.Suspend(member, opts...)
}

func (z memberWorkaround853) Unsuspend(member *mo_member.Member) (err error) {
	return z.v2.Unsuspend(member)
}
