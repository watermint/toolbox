package mo_filter

import (
	"flag"
	"github.com/iancoleman/strcase"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/infra/ui/app_ui"
)

type Acceptor interface {
	Accept(v interface{}) bool
}

type FlagSetter interface {
	// Flag name
	Name() string

	// Set flags
	ApplyFlags(f *flag.FlagSet, fieldDesc app_msg.Message, ui app_ui.UI)
}

type Filter interface {
	Acceptor
	FlagSetter

	SetOptions(o ...FilterOpt)
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
		desc := ui.Text(fieldDesc) + " " + ui.Text(f.Desc())
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
