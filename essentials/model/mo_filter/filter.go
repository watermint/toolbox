package mo_filter

import (
	"flag"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
	"github.com/watermint/toolbox/essentials/islet/estring/ecase"
	"github.com/watermint/toolbox/essentials/model/mo_multi"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type Acceptor interface {
	Accept(v interface{}) bool
}

type Filter interface {
	Acceptor
	mo_multi.MultiValue

	SetOptions(o ...FilterOpt)

	// Debug information
	Debug() interface{}

	// Weather the filter enabled or not.
	// Note: Accept() will return true when the filter is not enabled.
	IsEnabled() bool
}

type FilterOpt interface {
	Acceptor

	// Bind variable for configure flag
	Bind() interface{}

	// Filter option name suffix.
	NameSuffix() string

	// Description of this filter option.
	Desc() app_msg.Message

	// True if the option enabled thru the flag.
	Enabled() bool

	// Serialize settings
	Capture() interface{}

	// Restore settings
	Restore(v es_json.Json) error
}

func New(name string) Filter {
	return &filterImpl{
		name:    name,
		filters: make([]FilterOpt, 0),
	}
}

type filterImpl struct {
	name    string
	filters []FilterOpt
}

func (z *filterImpl) Capture() interface{} {
	s := make(map[string]interface{})
	for _, f := range z.filters {
		s[f.NameSuffix()] = f.Capture()
	}
	return s
}

func (z *filterImpl) Restore(v es_json.Json) error {
	if obj, found := v.Object(); found {
		for _, f := range z.filters {
			if fv, ok := obj[f.NameSuffix()]; ok {
				if err := f.Restore(fv); err != nil {
					return err
				}
			}
		}
		return nil

	} else {
		return rc_recipe.ErrorValueRestoreFailed
	}
}

func (z *filterImpl) IsEnabled() bool {
	for _, f := range z.filters {
		if f.Enabled() {
			return true
		}
	}
	return false
}

func (z *filterImpl) Debug() interface{} {
	filterNames := make([]string, 0)
	for _, f := range z.filters {
		filterNames = append(filterNames, f.NameSuffix())
	}
	return map[string]interface{}{
		"name":    z.name,
		"options": filterNames,
	}
}

func (z *filterImpl) Name() string {
	return z.name
}

func (z *filterImpl) SetName(name string) {
	z.name = name
}

func (z *filterImpl) Accept(v interface{}) bool {
	noEnabled := true
	for _, f := range z.filters {
		if f.Enabled() {
			noEnabled = false
			if f.Accept(v) {
				return true
			}
		}
	}

	// always accept if there is no enabled filters.
	return noEnabled
}

func (z *filterImpl) ApplyFlags(fl *flag.FlagSet, fieldDesc app_msg.Message, ui app_ui.UI) {
	for _, f := range z.filters {
		name := ecase.ToLowerKebabCase(z.Name() + f.NameSuffix())
		desc := ui.Text(app_ui.Join(ui, fieldDesc, f.Desc()))
		bind := f.Bind()
		switch bv := bind.(type) {
		case *bool:
			fl.BoolVar(bv, name, *bv, desc)
		case *int64:
			fl.Int64Var(bv, name, *bv, desc)
		case *string:
			fl.StringVar(bv, name, *bv, desc)
		}
	}
}

func (z *filterImpl) SetOptions(o ...FilterOpt) {
	z.filters = o
}

func (z *filterImpl) Fields() []string {
	fields := make([]string, 0)
	for _, f := range z.filters {
		fields = append(fields, z.Name()+f.NameSuffix())
	}
	return fields
}

func (z *filterImpl) FieldDesc(base app_msg.Message, name string) app_msg.Message {
	fields := make(map[string]FilterOpt)
	for _, f := range z.filters {
		key := z.Name() + f.NameSuffix()
		fields[key] = f
	}
	if m, ok := fields[name]; ok {
		return app_msg.Join(base, m.Desc())
	} else {
		return base
	}
}
