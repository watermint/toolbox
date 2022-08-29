package bulk

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_client"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_error"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_group_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_model"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgOperation struct {
	SkipTheUserAlreadyInTheGroup app_msg.Message
	SkipTheUserNotInTheGroup     app_msg.Message
}

var (
	MOperation = app_msg.Apply(&MsgOperation{}).(*MsgOperation)
)

func memberAdd(r *MemberRecord, svg sv_group.Group, c app_control.Control, d dbx_client.Client, ol rp_model.TransactionReport) error {
	l := c.Log().With(esl.Any("record", r))
	group, err := svg.ResolveByName(r.GroupName)
	if err != nil {
		l.Debug("Unable to resolve the group", esl.Error(err))
		return err
	}

	updated, err := sv_group_member.New(d, group).Add(sv_group_member.ByEmail(r.MemberEmail))
	if err != nil {
		if dbx_error.NewErrors(err).IsDuplicateUser() {
			ol.Skip(MOperation.SkipTheUserAlreadyInTheGroup, r)
			l.Debug("The user is already in the member", esl.Error(err))
			return nil
		}

		l.Debug("Unable to update group member", esl.Error(err))
		ol.Failure(err, r)
		return err
	}
	l.Debug("The member is successfully updated", esl.Any("updated", updated))
	ol.Success(r, nil)
	return nil
}

func memberDelete(r *MemberRecord, svg sv_group.Group, c app_control.Control, d dbx_client.Client, ol rp_model.TransactionReport) error {
	l := c.Log().With(esl.Any("record", r))
	group, err := svg.ResolveByName(r.GroupName)
	if err != nil {
		l.Debug("Unable to resolve the group", esl.Error(err))
		return err
	}

	updated, err := sv_group_member.New(d, group).Remove(sv_group_member.ByEmail(r.MemberEmail))
	if err != nil {
		if dbx_error.NewErrors(err).IsMemberNotInGroup() {
			ol.Skip(MOperation.SkipTheUserNotInTheGroup, r)
			l.Debug("The user is not in the group", esl.Error(err))
			return nil
		}

		l.Debug("Unable to remove the member", esl.Error(err))
		ol.Failure(err, r)
		return err
	}
	l.Debug("The member is successfully removed", esl.Any("updated", updated))
	ol.Success(r, nil)
	return nil
}
