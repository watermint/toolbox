package encode

import (
	"encoding/base32"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Base32 struct {
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
	t := coder.EncodeToString([]byte(z.Text))
	if c.Feature().IsQuiet() {
		_, _ = es_stdout.NewDirectOut().Write([]byte(t))
	} else {
		c.UI().Info(app_msg.Raw(t))
	}
	return nil
}

func (z *Base32) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Base32{}, func(r rc_recipe.Recipe) {
		m := r.(*Base32)
		m.Text = "watermint toolbox"
	})
}
