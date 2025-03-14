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

type Target struct {
	Member *mo_member.Member
	Entry  *mo_sharedlink.SharedLinkMember
}

type DeleteOnSuccess func(t *Target)
type DeleteOnFailure func(t *Target, cause error)

func DeleteMemberLinkWithSel(target *Target, c app_control.Control, client dbx_client.Client, onSuccess DeleteOnSuccess, onFailure DeleteOnFailure, sel Selector, baseNamespace dbx_filesystem.BaseNamespaceType) error {
	defer func() {
		_ = sel.Processed(target.Entry.Url)
	}()
	return DeleteMemberLink(target, c, client, onSuccess, onFailure, baseNamespace)
}

func DeleteMemberLink(target *Target, c app_control.Control, client dbx_client.Client, onSuccess DeleteOnSuccess, onFailure DeleteOnFailure, baseNamespace dbx_filesystem.BaseNamespaceType) error {
	l := c.Log().With(esl.String("member", target.Member.Email))
	mc := client.AsMemberId(target.Member.TeamMemberId, baseNamespace)

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
