package mo_sharedfolder_member

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgInternalOpt struct {
	Desc app_msg.Message
}

var (
	MInternalOpt = app_msg.Apply(&MsgInternalOpt{}).(*MsgInternalOpt)
)

func NewInternalOpt() FolderMemberFilter {
	return &internalOpt{}
}

type internalOpt struct {
	enabled bool
	members []*mo_member.Member
}

func (z *internalOpt) SetMembers(members []*mo_member.Member) {
	z.members = members
}

func (z *internalOpt) Accept(v interface{}) bool {
	l := esl.Default()
	switch m := v.(type) {
	case *User:
		return m.IsSameTeam

	case *Group:
		return m.IsSameTeam

	case *Invitee:
		// Returns true when the invitee is active or suspended.
		for _, tm := range z.members {
			if m.InviteeEmail == tm.Email {
				return tm.Profile().IsAccepted()
			}
		}
		return false

	case Member:
		if u, ok := m.User(); ok {
			return z.Accept(u)
		}
		if g, ok := m.Group(); ok {
			return z.Accept(g)
		}
		if i, ok := m.Invitee(); ok {
			return z.Accept(i)
		}
		l.Debug("Unexpected type of shared folder member", esl.String("member", m.MemberType()))
		// unknown type
		return false

	}
	l.Debug("Unexpected value type", esl.Any("value", v))
	return false
}

func (z *internalOpt) Bind() interface{} {
	return &z.enabled
}

func (z *internalOpt) NameSuffix() string {
	return "Internal"
}

func (z *internalOpt) Desc() app_msg.Message {
	return MInternalOpt.Desc
}

func (z *internalOpt) Enabled() bool {
	return z.enabled
}
