package lang

import (
	"github.com/watermint/toolbox/essentials/log/es_log"
	"golang.org/x/text/language"
)

func Select(bcp47 string, supported []Lang) Lang {
	l := es_log.Default().With(es_log.String("bcp47", bcp47))

	tag, err := language.Parse(bcp47)
	if err != nil {
		l.Debug("unable to parse language, fallback", es_log.Error(err))
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
			es_log.Error(err),
			es_log.String("confidence", confidence.String()))
		return Default

	default:
		l.Debug("selected",
			es_log.String("selected", selected.String()),
			es_log.String("confidence", confidence.String()))
		return New(selected)
	}
}
