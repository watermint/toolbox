package rc_spec

import (
	"flag"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_vo_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
	"sort"
)

func newSideCarValue(vc *rc_vo_impl.ValueContainer) rc_recipe.SpecValue {
	keys := make([]string, 0)
	for k := range vc.Values {
		keys = append(keys, k)
	}
	sort.Strings(keys)
	return &SpecSideCarValue{
		vc:        vc,
		valueKeys: keys,
	}
}

type SpecSideCarValue struct {
	scr       rc_recipe.SideCarRecipe
	vc        *rc_vo_impl.ValueContainer
	valueKeys []string
}

func (z *SpecSideCarValue) SetFlags(f *flag.FlagSet, ui app_ui.UI) {
	z.vc.MakeFlagSet(f, ui)
}

func (z *SpecSideCarValue) ValueNames() []string {
	return z.valueKeys
}

func (z *SpecSideCarValue) ValueDesc(name string) app_msg.Message {
	return app_msg.M(z.vc.MessageKey(name))
}

func (z *SpecSideCarValue) ValueDefault(name string) interface{} {
	switch d := z.vc.Values[name].(type) {
	case fd_file.ModelFile:
		return ""
	default:
		return d
	}
}

func (z *SpecSideCarValue) ValueCustomDefault(name string) app_msg.MessageOptional {
	return app_msg.M(z.vc.MessageKey(name) + ".default").AsOptional()
}
