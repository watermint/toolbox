package app_lang

import (
	"github.com/cloudfoundry-attic/jibber_jabber"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

var (
	SupportedLanguages = []language.Tag{
		// optional languages
		language.Japanese,

		// default & fallback language
		language.English,
	}
)

func Detect() language.Tag {
	bcp47, err := jibber_jabber.DetectIETF()
	if err != nil {
		app_root.Log().Debug("unable to detect language", zap.Error(err))
		return language.English
	}

	return Select(bcp47)
}

func Select(bcp47 string) language.Tag {
	if bcp47 == "" {
		return language.English
	}
	tag, err := language.Parse(bcp47)
	if err != nil {
		app_root.Log().Debug("unable to parse language into tag", zap.String("bcp47", bcp47), zap.Error(err))
		return language.English
	}
	m := language.NewMatcher(SupportedLanguages)
	l, _, c := m.Match(tag)
	app_root.Log().Debug("detect language", zap.Any("lang", l), zap.String("confidence", c.String()))

	return l
}

func Base(l language.Tag) string {
	base, _, _ := l.Raw()
	return base.String()
}

func PathSuffix(l language.Tag) string {
	if l == language.English {
		return ""
	}
	return "_" + Base(l)
}
