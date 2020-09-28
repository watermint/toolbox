package mo_filter

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
)

func NewNameSuffixFilter() FilterOpt {
	return &nameFilterSuffixOpt{}
}

type nameFilterSuffixOpt struct {
	suffix string
}

func (z *nameFilterSuffixOpt) Capture() interface{} {
	return z.suffix
}

func (z *nameFilterSuffixOpt) Restore(v es_json.Json) error {
	if w, found := v.String(); found {
		z.suffix = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *nameFilterSuffixOpt) Accept(v interface{}) bool {
	return ExpectString(v, func(s string) bool {
		return strings.HasSuffix(s, z.suffix)
	})
}

func (z *nameFilterSuffixOpt) Bind() interface{} {
	return &z.suffix
}

func (z *nameFilterSuffixOpt) NameSuffix() string {
	return "NameSuffix"
}

func (z *nameFilterSuffixOpt) Desc() app_msg.Message {
	return MFilter.DescFilterNameSuffix
}

func (z *nameFilterSuffixOpt) Enabled() bool {
	return z.suffix != ""
}
