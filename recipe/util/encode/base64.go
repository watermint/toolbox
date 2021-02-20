package encode

import (
	"encoding/base64"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
)

type Base64 struct {
	rc_recipe.RemarkTransient
	Text      string
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
	ui_out.TextOut(c, coder.EncodeToString([]byte(z.Text)))

	return nil
}

func (z *Base64) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Base64{}, func(r rc_recipe.Recipe) {
		m := r.(*Base64)
		m.Text = "watermint toolbox"
	})
}
