package dev

import (
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/recipe/rc_spec"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"go.uber.org/zap"
	"sort"
	"strings"
)

type Preflight struct {
	TestMode bool
}

func (z *Preflight) Preset() {
	z.TestMode = false
}

func (z *Preflight) Hidden() {
}

func (z *Preflight) Console() {
}

func (z *Preflight) Test(c app_control.Control) error {
	z.TestMode = true
	return z.Exec(c)
}

func (z *Preflight) Exec(c app_control.Control) error {
	l := c.Log()
	{
		l.Info("Generating English documents")
		err := rc_exec.Exec(c, &Doc{}, func(r rc_recipe.Recipe) {
			rr := r.(*Doc)
			rr.TestMode = z.TestMode
			rr.Badge = true
			rr.MarkdownReadme = true
			rr.Lang = ""
			rr.Filename = "README.md"
			rr.CommandPath = "doc/generated/"
		})
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}
	{
		l.Info("Generating Japanese documents")
		err := rc_exec.Exec(c, &Doc{}, func(r rc_recipe.Recipe) {
			rr := r.(*Doc)
			rr.TestMode = z.TestMode
			rr.Badge = true
			rr.MarkdownReadme = true
			rr.Lang = "ja"
			rr.Filename = "README_ja.md"
			rr.CommandPath = "doc/generated_ja/"
		})
		if err != nil {
			l.Error("Failed to generate documents", zap.Error(err))
			return err
		}
	}

	{
		cl := c.(app_control_launcher.ControlLauncher)
		l.Info("Verify recipes")
		for _, r := range cl.Catalogue() {
			spec := rc_spec.New(r)
			if spec == nil {
				continue
			}
			for _, m := range spec.Messages() {
				l.Debug("message", zap.String("key", m.Key()), zap.String("text", c.UI().Text(m.Key())))
			}
		}

		l.Info("Verify ingredients")
		for _, r := range cl.Ingredients() {
			spec := rc_spec.New(r)
			if spec == nil {
				continue
			}
			for _, m := range spec.Messages() {
				l.Debug("message", zap.String("key", m.Key()), zap.String("text", c.UI().Text(m.Key())))
			}
		}
	}

	l.Info("Verify message resources")
	qm := c.Messages().(app_msg_container.Quality)
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

	return qt_messages.VerifyMessages(c)
}
