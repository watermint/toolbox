package lang

import (
	"github.com/watermint/toolbox/essentials/log/es_fallback"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

func Select(bcp47 string, supported []Lang) Lang {
	l := es_fallback.Fallback().With(zap.String("bcp47", bcp47))

	tag, err := language.Parse(bcp47)
	if err != nil {
		l.Debug("unable to parse language, fallback", zap.Error(err))
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
			zap.Error(err),
			zap.String("confidence", confidence.String()))
		return Default

	default:
		l.Debug("selected",
			zap.String("selected", selected.String()),
			zap.String("confidence", confidence.String()))
		return New(selected)
	}
}
