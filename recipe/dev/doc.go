package dev

import (
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/util/ut_doc"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg"
)

type Doc struct {
	rc_recipe.RemarkSecret
	Badge          bool
	MarkdownReadme bool
	Lang           mo_string.OptionalString
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

	if z.Lang.IsExists() {
		ctl = ctl.WithLang(z.Lang.Value())
	}

	rme := ut_doc.NewReadme(ctl, z.Filename, z.Badge, z.MarkdownReadme, z.CommandPath)
	cmd := ut_doc.NewCommandWithPath(z.CommandPath)
	if err := rme.Generate(); err != nil {
		l.Error("Failed to generate README", es_log.Error(err))
		return err
	}
	if err := cmd.GenerateAll(ctl); err != nil {
		l.Error("Failed to generate command manuals", es_log.Error(err))
		return err
	}

	missing := qt_missingmsg.Record().Missing()
	if len(missing) > 0 {
		return qt_messages.VerifyMessages(ctl)
	}
	return nil
}

func (z *Doc) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Doc{}, func(r rc_recipe.Recipe) {
		rr := r.(*Doc)
		rr.Badge = false
	})
}
