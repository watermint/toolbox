package es_lang

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"golang.org/x/text/language"
)

func Select(bcp47 string, supported []Lang) Lang {
	l := esl.Default().With(esl.String("bcp47", bcp47))

	tag, err := language.Parse(bcp47)
	if err != nil {
		l.Debug("unable to parse language, fallback", esl.Error(err))
		return Default
	}
	supportedTags := make([]language.Tag, len(supported))
	for i, s := range supported {
		supportedTags[i] = s.Tag()
	}
	m := language.NewMatcher(supportedTags)
	selected, _, confidence := m.Match(tag)
	switch confidence {
	case language.No, language.Low:
		l.Debug("fallback to default, due to lower confidence",
			esl.Error(err),
			esl.String("confidence", confidence.String()))
		return Default

	default:
		l.Debug("selected",
			esl.String("selected", selected.String()),
			esl.String("confidence", confidence.String()))
		return New(selected)
	}
}
