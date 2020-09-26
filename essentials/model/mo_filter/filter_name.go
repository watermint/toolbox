package mo_filter

import (
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

func NewNameFilter() FilterOpt {
	return &nameFilterOpt{}
}

func NewTestNameFilter(name string) FilterOpt {
	return &nameFilterOpt{
		name: name,
	}
}

type nameFilterOpt struct {
	name string
}

func (z *nameFilterOpt) Capture() interface{} {
	return z.name
}

func (z *nameFilterOpt) Restore(v es_json.Json) error {
	if w, found := v.String(); found {
		z.name = w
		return nil
	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
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
