package encode

import (
	"encoding/base32"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Base32 struct {
	rc_recipe.RemarkTransient
	Text      da_text.TextInput
	NoPadding bool
}

func (z *Base32) Preset() {
}

func (z *Base32) Exec(c app_control.Control) error {
	var coder *base32.Encoding
	if z.NoPadding {
		coder = base32.StdEncoding.WithPadding(base32.NoPadding)
	} else {
		coder = base32.StdEncoding
	}
	tx, err := z.Text.Content()
	if err != nil {
		return err
	}
	ui_out.TextOut(c, coder.EncodeToString(tx))
	return nil
}

func (z *Base32) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("encode", "watermint toolbox")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.Exec(c, &Base32{}, func(r rc_recipe.Recipe) {
		m := r.(*Base32)
		m.Text.SetFilePath(f)
	})
}
