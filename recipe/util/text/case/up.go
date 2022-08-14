package _case

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"strings"
)

type Up struct {
	rc_recipe.RemarkTransient
	Text da_text.TextInput
}

func (z *Up) Preset() {
}

func (z *Up) Exec(c app_control.Control) error {
	content, err := z.Text.Content()
	if err != nil {
		return err
	}
	ui_out.TextOut(c, strings.ToUpper(string(content)))
	return nil
}

func (z *Up) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("case", "Hello World")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.Exec(c, &Up{}, func(r rc_recipe.Recipe) {
		m := r.(*Up)
		m.Text.SetFilePath(f)
	})
}
