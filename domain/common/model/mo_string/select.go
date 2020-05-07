package mo_string

import "github.com/watermint/toolbox/essentials/log/es_log"

type SelectString interface {
	String
	Options() []string
	SetOptions(opts []string, selected string)
	SetSelect(selected string)
	IsValid() bool
}

func NewSelect() SelectString {
	return &selectString{
		opts:     []string{},
		selected: "",
	}
}

type selectString struct {
	opts     []string
	selected string
}

func (z *selectString) Value() string {
	return z.selected
}

func (z *selectString) Options() []string {
	return z.opts
}

func (z *selectString) SetOptions(opts []string, selected string) {
	z.opts = opts
	z.SetSelect(selected)
}

func (z *selectString) SetSelect(selected string) {
	z.selected = selected
	if !z.IsValid() {
		l := es_log.Default()
		l.Debug("The value `selected` is not in `opts`",
			es_log.String("selected", selected),
			es_log.Strings("opts", z.opts))
	}
}

func (z *selectString) IsValid() bool {
	for _, s := range z.opts {
		if s == z.selected {
			return true
		}
	}
	return false
}
