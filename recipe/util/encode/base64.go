package encode

import (
	"encoding/base64"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

type Base64 struct {
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
	t := coder.EncodeToString([]byte(z.Text))
	if c.Feature().IsQuiet() {
		_, _ = es_stdout.NewDirectOut().Write([]byte(t))
	} else {
		c.UI().Info(app_msg.Raw(t))
	}

	return nil
}

func (z *Base64) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Base64{}, func(r rc_recipe.Recipe) {
		m := r.(*Base64)
		m.Text = "watermint toolbox"
	})
}
