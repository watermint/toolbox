package mo_string

import "github.com/watermint/toolbox/essentials/log/esl"

type SelectString interface {
	String
	Options() []string
	SetOptions(selected string, opts ...string)
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

func (z *selectString) SetOptions(selected string, opts ...string) {
	z.opts = opts
	z.SetSelect(selected)
}

func (z *selectString) SetSelect(selected string) {
	z.selected = selected
	if !z.IsValid() {
		l := esl.Default()
		l.Debug("The value `selected` is not in `opts`",
			esl.String("selected", selected),
			esl.Strings("opts", z.opts))
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
