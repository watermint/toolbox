package app_lang

import (
	"github.com/GeertJohan/go.rice"
	"github.com/cloudfoundry-attic/jibber_jabber"
	"github.com/watermint/toolbox/app86/app_msg_container"
	"github.com/watermint/toolbox/app86/app_msg_container_impl"
	"github.com/watermint/toolbox/app86/app_root"
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

func NewContainer(box *rice.Box) app_msg_container.Container {
	cm := make(map[language.Tag]app_msg_container.Container)
	langs := make([]language.Tag, 0)

	usrLang := DetectLang()
	if usrLang != language.English {
		langs = append(langs, usrLang)
	}
	langs = append(langs, language.English)

	for _, lang := range langs {
		c, err := app_msg_container_impl.NewResource(lang, box)
		if err != nil {
			app_root.Log().Error("unable to load message resource", zap.Error(err))
			app_root.Root.Fatal()
		}
		cm[lang] = c
	}

	return &app_msg_container_impl.Chain{
		LangPriority: langs,
		Containers:   cm,
	}
}

func DetectLang() language.Tag {
	bcp47, err := jibber_jabber.DetectIETF()
	if err != nil {
		app_root.Log().Debug("unable to detect language", zap.Error(err))
		return language.English
	}

	return chooseLanguage(bcp47)
}

func chooseLanguage(bcp47 string) language.Tag {
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
