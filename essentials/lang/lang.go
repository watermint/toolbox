package lang

import (
	"fmt"
	"golang.org/x/text/language"
	"golang.org/x/text/language/display"
)

var (
	English  = New(language.English)
	Japanese = New(language.Japanese)

	// Default language
	Default = English

	Supported = []Lang{English, Japanese}
)

type Lang interface {
	fmt.Stringer

	// ISO 639-1 two letter code. e.g. `en` for English, `ja` for Japanese.
	// https://en.wikipedia.org/wiki/List_of_ISO_639-1_codes
	Code() Iso639One

	// ISO 639-1 two letter code in string.
	CodeString() string

	// Is default language of this system.
	IsDefault() bool

	// Language tag
	Tag() language.Tag

	// Resource path suffix
	Suffix() string

	// Display language name in this language
	Self() string

	// True when the x matches high confidence or exact match.
	Equals(x Lang) bool
}

// ISO 639-1 two letter code
type Iso639One string

func New(tag language.Tag) Lang {
	c, _, _ := tag.Raw()
	return &langImpl{
		tag:  tag,
		code: c.String(),
	}
}

type langImpl struct {
	tag  language.Tag
	code string
}

func (z langImpl) Self() string {
	return display.Self.Name(z.tag)
}

func (z langImpl) String() string {
	return z.tag.String()
}

func (z langImpl) Equals(x Lang) bool {
	matcher := language.NewMatcher([]language.Tag{z.tag})
	_, _, c := matcher.Match(x.Tag())
	switch c {
	case language.High, language.Exact:
		return true
	default:
		return false
	}
}

func (z langImpl) IsDefault() bool {
	return z.Equals(Default)
}

func (z langImpl) Suffix() string {
	if z.IsDefault() {
		return ""
	} else {
		return "_" + z.CodeString()
	}
}

func (z langImpl) Code() Iso639One {
	return Iso639One(z.code)
}

func (z langImpl) CodeString() string {
	return z.code
}

func (z langImpl) Tag() language.Tag {
	return z.tag
}
