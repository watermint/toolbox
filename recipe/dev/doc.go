package dev

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_control_launcher"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg_container"
	"github.com/watermint/toolbox/infra/util/ut_doc"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"go.uber.org/zap"
)

type Doc struct {
	Badge          bool
	MarkdownReadme bool
	Lang           string
	Filename       string
	CommandPath    string
}

func (z *Doc) Preset() {
	z.Badge = true
	z.Filename = "README.md"
	z.CommandPath = "doc/generated/"
}

func (z *Doc) Exec(ctl app_control.Control) error {
	l := ctl.Log()

	if z.Lang != "" {
		if c, ok := app_control_launcher.ControlWithLang(z.Lang, ctl); ok {
			ctl = c
		}
	}

	rme := ut_doc.NewReadme(ctl, z.Filename, z.Badge, z.MarkdownReadme, z.CommandPath)
	cmd := ut_doc.NewCommandWithPath(z.CommandPath)
	if err := rme.Generate(); err != nil {
		l.Error("Failed to generate README", zap.Error(err))
		return err
	}
	if err := cmd.GenerateAll(ctl); err != nil {
		l.Error("Failed to generate command manuals", zap.Error(err))
		return err
	}

	qm := ctl.Messages().(app_msg_container.Quality)
	missing := qm.MissingKeys()
	if len(missing) > 0 {
		return qt_messages.VerifyMessages(ctl)
	}
	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Doc{}, func(r rc_recipe.Recipe) {
		rr := r.(*Doc)
		rr.Badge = false
		rr.Filename = ""
	})
}
