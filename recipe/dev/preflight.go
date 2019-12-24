package dev

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_kitchen"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"go.uber.org/zap"
	"sort"
	"strings"
)

type Preflight struct {
	TestMode bool
}

func (z *Preflight) Init() {
	z.TestMode = false
}

func (z *Preflight) Hidden() {
}

func (z *Preflight) Console() {
}

func (z *Preflight) Test(c app_control.Control) error {
	z.TestMode = true
	return z.Exec(rc_kitchen.NewKitchen(c, z))
}

func (z *Preflight) Exec(k rc_kitchen.Kitchen) error {
	l := k.Log()
	{
		l.Info("Generating English documents")
		r := Doc{}
		rv := &DocVO{
			Test:           z.TestMode,
			Badge:          true,
			MarkdownReadme: true,
			Lang:           "",
			Filename:       "README.md",
			CommandPath:    "doc/generated/",
		}
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), rv))
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}
	{
		l.Info("Generating Japanese documents")
		r := Doc{}
		rv := &DocVO{
			Test:           z.TestMode,
			Badge:          true,
			MarkdownReadme: true,
			Lang:           "ja",
			Filename:       "README_ja.md",
			CommandPath:    "doc/generated_ja/",
		}
		err := r.Exec(rc_kitchen.NewKitchen(k.Control(), rv))
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}

	l.Info("Verify message resources")
	qm := k.Control().Messages().(app_msg_container.Quality)
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

	return qt_messages.VerifyMessages(k.Control())
}
