package dev

import (
	"fmt"
	"github.com/watermint/toolbox/domain/common/model/mo_string"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_command"
	"github.com/watermint/toolbox/infra/doc/dc_readme"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_messages"
	"github.com/watermint/toolbox/quality/infra/qt_missingmsg"
	"io/ioutil"
)

type Doc struct {
	rc_recipe.RemarkSecret
	Badge       bool
	Lang        mo_string.OptionalString
	Filename    string
	CommandPath string
}

func (z *Doc) Preset() {
	z.Badge = true
	z.Filename = "README.md"
	z.CommandPath = "doc/generated/"
}

func (z *Doc) genReadme(c app_control.Control) error {
	l := c.Log()
	l.Info("Generating README", esl.String("file", z.Filename))
	sec := dc_readme.New(z.Badge, z.CommandPath)
	doc := dc_section.Document(c.Messages(), sec...)

	if c.Feature().IsTest() {
		out := es_stdout.NewDefaultOut(c.Feature().IsTest())
		_, _ = fmt.Fprintln(out, doc)
		return nil
	} else {
		return ioutil.WriteFile(z.Filename, []byte(doc), 0644)
	}
}

func (z *Doc) Exec(ctl app_control.Control) error {
	l := ctl.Log()

	if z.Lang.IsExists() {
		ctl = ctl.WithLang(z.Lang.Value())
	}
	if err := z.genReadme(ctl); err != nil {
		l.Error("Failed to generate README", esl.Error(err))
		return err
	}

	cmd := dc_command.NewCommandWithPath(z.CommandPath)
	if err := cmd.GenerateAll(ctl); err != nil {
		l.Error("Failed to generate command manuals", esl.Error(err))
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
