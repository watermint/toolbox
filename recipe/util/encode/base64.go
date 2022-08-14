package encode

import (
	"encoding/base64"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Base64 struct {
	rc_recipe.RemarkTransient
	Text      da_text.TextInput
	NoPadding bool
}

func (z *Base64) Preset() {
}

func (z *Base64) Exec(c app_control.Control) error {
	var coder *base64.Encoding
	if z.NoPadding {
		coder = base64.StdEncoding.WithPadding(base64.NoPadding)
	} else {
		coder = base64.StdEncoding
	}
	content, err := z.Text.Content()
	if err != nil {
		return err
	}
	ui_out.TextOut(c, coder.EncodeToString(content))

	return nil
}

func (z *Base64) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("encode", "watermint toolbox")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.Exec(c, &Base64{}, func(r rc_recipe.Recipe) {
		m := r.(*Base64)
		m.Text.SetFilePath(f)
	})
}
