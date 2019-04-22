package sq_group

import (
	"github.com/watermint/toolbox/domain/infra/api_util"
	"github.com/watermint/toolbox/domain/service"
	"github.com/watermint/toolbox/domain/service/sv_group_member"
	"go.uber.org/zap"
	"strings"
)

type AddMember struct {
	GroupName   string `json:"group_name"`
	MemberEmail string `json:"member_email"`
}

func (z *AddMember) Do(biz service.Business) error {
	l := biz.Log().With(zap.Any("task", z))

	l.Debug("Resolve group")
	group, err := biz.Group().ResolveByName(z.GroupName)
	if err != nil {
		l.Debug("Unable to find group", zap.Error(err))
		return err
	}
	l = l.With(zap.Any("group", group))

	l.Debug("Add member")
	group, err = biz.GroupMember(group.GroupId).Add(sv_group_member.ByEmail(z.MemberEmail))
	if err != nil {
		es := api_util.ErrorSummary(err)
		switch {
		case strings.HasPrefix(es, "duplicate_user"):
			l.Debug("Skip duplicated user")
			return nil

		default:
			l.Debug("Unable to add member", zap.Error(err))
			return err
		}
	}

	l.Debug("Success")
	return nil
}
