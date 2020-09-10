package mo_filter

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"strings"
)

func NewNamePrefixFilter() FilterOpt {
	return &nameFilterPrefixOpt{}
}

type nameFilterPrefixOpt struct {
	prefix string
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
