package sv_member

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_list"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

func NewListV1(ctx dbx_client.Client, limit int) Member {
	return &memberListV1{
		ctx:      ctx,
		delegate: New(ctx),
	}
}

// memberListV1 uses `members/list` instead of `members/list_v2` for special use cases.
type memberListV1 struct {
	ctx      dbx_client.Client
	limit    int
	delegate Member
}

func (z memberListV1) List(opts ...ListOpt) (members []*mo_member.Member, err error) {
	members = make([]*mo_member.Member, 0)
	err = z.ListEach(func(member *mo_member.Member) bool {
		members = append(members, member)
		return true
	}, opts...)
	return
}

func (z memberListV1) ListEach(f func(member *mo_member.Member) bool, opts ...ListOpt) error {
	lo := listOpts{}.Apply(opts...)
	p := struct {
		IncludeRemoved bool `json:"include_removed,omitempty"`
		Limit          int  `json:"limit,omitempty"`
	}{
		IncludeRemoved: lo.includeRemoved,
		Limit:          z.limit,
	}

	ErrorBreak := errors.New("break")
	res := z.ctx.List("team/members/list_v2", api_request.Param(p)).Call(
		dbx_list.Continue("team/members/list/continue"),
		dbx_list.UseHasMore(),
		dbx_list.ResultTag("members"),
		dbx_list.OnEntry(func(entry es_json.Json) error {
			m := &mo_member.Member{}
			if err := entry.Model(m); err != nil {
				return err
			}
			if !f(m) {
				return ErrorBreak
			}
			return nil
		}),
	)
	if err, fail := res.Failure(); fail {
		if errors.Is(err, ErrorBreak) {
			return nil
		}
		return err
	}
	return nil
}

func (z memberListV1) Update(member *mo_member.Member, opts ...UpdateOpt) (updated *mo_member.Member, err error) {
	return z.delegate.Update(member, opts...)
}

func (z memberListV1) UpdateVisibility(email string, visible bool) (updated *mo_member.Member, err error) {
	return z.delegate.UpdateVisibility(email, visible)
}

func (z memberListV1) Resolve(teamMemberId string) (member *mo_member.Member, err error) {
	return z.delegate.Resolve(teamMemberId)
}

func (z memberListV1) ResolveByEmail(email string) (member *mo_member.Member, err error) {
	return z.delegate.ResolveByEmail(email)
}

func (z memberListV1) Add(email string, opts ...AddOpt) (member *mo_member.Member, err error) {
	return z.delegate.Add(email, opts...)
}

func (z memberListV1) Remove(member *mo_member.Member, opts ...RemoveOpt) (err error) {
	return z.delegate.Remove(member, opts...)
}

func (z memberListV1) Suspend(member *mo_member.Member, opts ...SuspendOpt) (err error) {
	return z.delegate.Suspend(member, opts...)
}

func (z memberListV1) Unsuspend(member *mo_member.Member) (err error) {
	return z.delegate.Unsuspend(member)
}
