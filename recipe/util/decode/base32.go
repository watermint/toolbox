package decode

import (
	"encoding/base32"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
)

type Base32 struct {
	rc_recipe.RemarkTransient
	Text      string
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
	t, err := coder.DecodeString(z.Text)
	if err != nil {
		return err
	}
	ui_out.TextOut(c, string(t))

	return nil
}

func (z *Base32) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Base32{}, func(r rc_recipe.Recipe) {
		m := r.(*Base32)
		m.Text = "O5QXIZLSNVUW45BAORXW63DCN54A===="
	})
}
