package dev

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/recipe/rc_vo"
	"github.com/watermint/toolbox/infra/report/rp_spec"
	"github.com/watermint/toolbox/infra/ui/app_lang"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/ui/app_msg_container_impl"
	"github.com/watermint/toolbox/infra/util/ut_doc"
	"go.uber.org/zap"
	"golang.org/x/text/language"
	"sort"
	"strings"
)

type DocVO struct {
	Test           bool
	Badge          bool
	MarkdownReadme bool
	Lang           string
	Filename       string
	CommandPath    string
}

type Doc struct {
}

func (z *Doc) Reports() []rp_spec.ReportSpec {
	return []rp_spec.ReportSpec{}
}

func (z *Doc) Console() {
}

func (z *Doc) Hidden() {
}

func (z *Doc) Requirement() rc_vo.ValueObject {
	return &DocVO{
		Test:        false,
		Badge:       true,
		Filename:    "README.md",
		CommandPath: "doc/generated/",
	}
}

func (z *Doc) Exec(k rc_kitchen.Kitchen) error {
	vo := k.Value().(*DocVO)
	l := k.Log()
	ctl := k.Control()

	if vo.Lang != "" {
		wc := ctl.(app_control_launcher.WithMessageContainer)
		langPriority := make([]language.Tag, 0)
		ul := app_lang.Select(vo.Lang)
		if ul != language.English {
			langPriority = append(langPriority, ul)
		}
		langPriority = append(langPriority, language.English)
		langContainers := make(map[language.Tag]app_msg_container.Container)

		for _, lang := range langPriority {
			mc, err := app_msg_container_impl.New(lang, ctl)
			if err != nil {
				return err
			}
			langContainers[lang] = mc
		}

		ctl = wc.With(app_msg_container_impl.NewMultilingual(langPriority, langContainers))
	}

	rme := ut_doc.NewReadme(ctl, vo.Filename, vo.Badge, vo.Test, vo.MarkdownReadme, vo.CommandPath)
	cmd := ut_doc.NewCommand(ctl, vo.CommandPath, vo.Test)
	if err := rme.Generate(); err != nil {
		l.Error("Failed to generate README", zap.Error(err))
		return err
	}
	if err := cmd.GenerateAll(); err != nil {
		l.Error("Failed to generate command manuals", zap.Error(err))
		return err
	}

	qm := ctl.Messages().(app_msg_container.Quality)
	missing := qm.MissingKeys()
	if len(missing) > 0 {
		suggested := make([]string, 0)
		for _, k := range missing {
			l.Error("Key missing", zap.String("key", k))
			suggested = append(suggested, "\""+k+"\":\"\",")
		}
		sort.Strings(suggested)
		fmt.Println(strings.Join(suggested, "\n"))
	}
	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return z.Exec(rc_kitchen.NewKitchen(c, &DocVO{
		Test:     true,
		Badge:    false,
		Filename: "",
	}))
}
