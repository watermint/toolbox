package uc_team_sharedlink

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_filesystem"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_sharedlink"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_sharedlink"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
)

type OnSharedLinkMember func(member *mo_member.Member, entry *mo_sharedlink.SharedLinkMember)

func RetrieveMemberLinks(member *mo_member.Member, c app_control.Control, ctx dbx_client.Client, handler OnSharedLinkMember, baseNamespace dbx_filesystem.BaseNamespaceType) error {
	l := c.Log().With(esl.String("member", member.Email))
	mc := ctx.AsMemberId(member.TeamMemberId, baseNamespace)
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
