package mo_string

import (
	"github.com/watermint/toolbox/essentials/log/esl"
)

type SelectString interface {
	String
	Options() []string
	SetOptions(selected string, opts ...string)
	SetSelect(selected string)
	IsValid() bool
}

func NewSelect() SelectString {
	return &selectStringInternal{
		opts:     []string{},
		selected: "",
	}
}

type selectStringInternal struct {
	opts     []string
	selected string
}

func (z *selectStringInternal) Value() string {
	return z.selected
}

func (z *selectStringInternal) Options() []string {
	return z.opts
}

func (z *selectStringInternal) SetOptions(selected string, opts ...string) {
	z.opts = opts
	z.SetSelect(selected)
}

func (z *selectStringInternal) SetSelect(selected string) {
	z.selected = selected
	if !z.IsValid() {
		l := esl.Default()
		l.Debug("The value `selected` is not in `opts`",
			esl.String("selected", selected),
			esl.Strings("opts", z.opts))
	}
}

func (z *selectStringInternal) IsValid() bool {
	for _, s := range z.opts {
		if s == z.selected {
			return true
		}
	}
	return false
}
