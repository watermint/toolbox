package gui

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/ui/app_template"
)

type Page struct {
	Name    string
	Layouts []string
}

func (z Page) Apply(htp app_template.Template) error {
	l := esl.Default().With(esl.String("name", z.Name), esl.Strings("layouts", z.Layouts))

	if err := htp.Define(z.Name, z.Layouts...); err != nil {
		l.Debug("Unable to prepare templates", esl.Error(err))
		return err
	}
	l.Debug("The layout defined")
	return nil
}
