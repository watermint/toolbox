package app_run

import (
	rice "github.com/GeertJohan/go.rice"
	"github.com/watermint/toolbox/app86/app_lang"
	"github.com/watermint/toolbox/app86/app_msg_container"
	"github.com/watermint/toolbox/app86/app_msg_container_impl"
	"github.com/watermint/toolbox/app86/app_root"
	"go.uber.org/zap"
	"golang.org/x/text/language"
)

func NewContainer(box *rice.Box) app_msg_container.Container {
	cm := make(map[language.Tag]app_msg_container.Container)
	langs := make([]language.Tag, 0)

	usrLang := app_lang.DetectLang()
	if usrLang != language.English {
		langs = append(langs, usrLang)
	}
	langs = append(langs, language.English)
	var lastErr error = nil

	for _, lang := range langs {
		c, err := app_msg_container_impl.NewResource(lang, box)
		if err != nil {
			lastErr = err
			continue
		}
		cm[lang] = c
	}
	if len(cm) < 1 {
		app_root.Log().Error("No resources loaded", zap.Error(lastErr))
		panic("At least one message resource required")
	}

	return &app_msg_container_impl.Chain{
		LangPriority: langs,
		Containers:   cm,
	}
}