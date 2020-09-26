package mo_file_filter

import (
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_util"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/model/mo_filter"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type MsgFileFilterOpt struct {
	Desc app_msg.Message
}

var (
	MFileFilterOpt = app_msg.Apply(&MsgFileFilterOpt{}).(*MsgFileFilterOpt)
)

func NewIgnoreFileFilter() mo_filter.FilterOpt {
	return &filterIgnoreFileOpt{}
}

type filterIgnoreFileOpt struct {
	disabled bool
}

func (z *filterIgnoreFileOpt) Capture() interface{} {
	return z.disabled
}

func (z *filterIgnoreFileOpt) Restore(v es_json.Json) error {
	if w, found := v.Bool(); found {
		z.disabled = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *filterIgnoreFileOpt) Accept(v interface{}) bool {
	return mo_filter.ExpectString(v, func(s string) bool {
		return !dbx_util.IsFileNameIgnored(s)
	})
}

func (z *filterIgnoreFileOpt) Bind() interface{} {
	return &z.disabled
}

func (z *filterIgnoreFileOpt) NameSuffix() string {
	return "DisableIgnore"
}

func (z *filterIgnoreFileOpt) Desc() app_msg.Message {
	return MFileFilterOpt.Desc
}

func (z *filterIgnoreFileOpt) Enabled() bool {
	return !z.disabled
}
