package mo_filter

import (
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func NewNameFilter() FilterOpt {
	return &nameFilterOpt{}
}

type nameFilterOpt struct {
	name string
}

func (z *nameFilterOpt) Enabled() bool {
	return z.name != ""
}

func (z *nameFilterOpt) Bind() interface{} {
	return &z.name
}

func (z nameFilterOpt) Accept(v interface{}) bool {
	return ExpectString(v, func(s string) bool {
		return s == z.name
	})
}

func (z nameFilterOpt) NameSuffix() string {
	return "Name"
}

func (z nameFilterOpt) Desc() app_msg.Message {
	return MFilter.DescFilterName
}
