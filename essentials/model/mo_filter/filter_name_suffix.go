package mo_filter

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
)

func NewNameSuffixFilter() FilterOpt {
	return &nameFilterSuffixOpt{}
}

type nameFilterSuffixOpt struct {
	suffix string
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
