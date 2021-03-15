package uc_team_sharedlink

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_context"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type OnSharedLinkMember func(member *mo_member.Member, entry *mo_sharedlink.SharedLinkMember)

func RetrieveMemberLinks(member *mo_member.Member, c app_control.Control, ctx dbx_context.Context, handler OnSharedLinkMember) error {
	l := c.Log().With(esl.String("member", member.Email))
	mc := ctx.AsMemberId(member.TeamMemberId)
	links, err := sv_sharedlink.New(mc).List()
	if err != nil {
		l.Debug("Unable to retrieve shared links for the member", esl.Error(err))
		return err
	}
	l.Debug("Link found", esl.Int("numLinks", len(links)))
	for _, link := range links {
		lm := mo_sharedlink.NewSharedLinkMember(link, member)
		handler(member, lm)
	}
	return nil
}

type DeleteTarget struct {
	Member *mo_member.Member
	Entry  *mo_sharedlink.SharedLinkMember
}

type DeleteOnSuccess func(t *DeleteTarget)
type DeleteOnFailure func(t *DeleteTarget, cause error)

func DeleteMemberLink(target *DeleteTarget, c app_control.Control, ctx dbx_context.Context, onSuccess DeleteOnSuccess, onFailure DeleteOnFailure) error {
	l := c.Log().With(esl.String("member", target.Member.Email))
	mc := ctx.AsMemberId(target.Member.TeamMemberId)
	l.Debug("Delete link", esl.Any("target", target))
	rmErr := sv_sharedlink.New(mc).Remove(target.Entry.SharedLink())
	if rmErr != nil {
		l.Debug("Unable to remove the link", esl.Error(rmErr))
		onFailure(target, rmErr)
		return rmErr
	}

	onSuccess(target)
	return nil
}
