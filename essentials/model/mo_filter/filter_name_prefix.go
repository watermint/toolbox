package mo_filter

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
)

func NewNamePrefixFilter() FilterOpt {
	return &nameFilterPrefixOpt{}
}

type nameFilterPrefixOpt struct {
	prefix string
}

func (z *nameFilterPrefixOpt) Capture() interface{} {
	return z.prefix
}

func (z *nameFilterPrefixOpt) Restore(v es_json.Json) error {
	if w, found := v.String(); found {
		z.prefix = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *nameFilterPrefixOpt) Accept(v interface{}) bool {
	return ExpectString(v, func(s string) bool {
		return strings.HasPrefix(s, z.prefix)
	})
}

func (z *nameFilterPrefixOpt) Bind() interface{} {
	return &z.prefix
}

func (z *nameFilterPrefixOpt) NameSuffix() string {
	return "NamePrefix"
}

func (z *nameFilterPrefixOpt) Desc() app_msg.Message {
	return MFilter.DescFilterNamePrefix
}

func (z *nameFilterPrefixOpt) Enabled() bool {
	return z.prefix != ""
}
