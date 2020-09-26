package mo_sharedfolder_member

import (
	"github.com/watermint/toolbox/domain/dropbox/model/mo_member"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgExternalOpt struct {
	Desc app_msg.Message
}

var (
	MExternalOpt = app_msg.Apply(&MsgExternalOpt{}).(*MsgExternalOpt)
)

func NewExternalOpt() FolderMemberFilter {
	return &externalOpt{
		internal: NewInternalOpt(),
	}
}

type externalOpt struct {
	enabled  bool
	internal FolderMemberFilter
}

func (z *externalOpt) Capture() interface{} {
	return z.enabled
}

func (z *externalOpt) Restore(v es_json.Json) error {
	if w, found := v.Bool(); found {
		z.enabled = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *externalOpt) SetMembers(members []*mo_member.Member) {
	z.internal.SetMembers(members)
}

func (z *externalOpt) Accept(v interface{}) bool {
	l := esl.Default()
	switch m := v.(type) {
	case Member:
		return !z.internal.Accept(m)
	}
	l.Debug("Unexpected value type", esl.Any("value", v))
	return false
}

func (z *externalOpt) Bind() interface{} {
	return &z.enabled
}

func (z *externalOpt) NameSuffix() string {
	return "External"
}

func (z *externalOpt) Desc() app_msg.Message {
	return MExternalOpt.Desc
}

func (z *externalOpt) Enabled() bool {
	return z.enabled
}
