package mo_filter

import (
	"flag"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/domain/common/model/mo_multi"
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
		name := strcase.ToKebab(z.Name() + f.NameSuffix())
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
