package translate

import (
	"encoding/json"
	"github.com/watermint/toolbox/domain/deepl/api/deepl_conn"
	"github.com/watermint/toolbox/domain/deepl/service/sv_translate"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
)

type Text struct {
	Peer       deepl_conn.ConnDeeplApi
	SourceLang mo_string.OptionalString
	TargetLang string
	Text       string
}

func (z *Text) Preset() {
}

func (z *Text) Exec(c app_control.Control) error {
	res, err := sv_translate.New(z.Peer.Client()).Translate(z.Text, z.SourceLang.Value(), z.TargetLang)
	if err != nil {
		return err
	}
	out, err := json.Marshal(res)
	if err != nil {
		return err
	}
	ui_out.TextOut(c, string(out))
	return nil
}

func (z *Text) Test(c app_control.Control) error {
	return rc_exec.ExecMock(c, &Text{}, func(r rc_recipe.Recipe) {
		m := r.(*Text)
		m.Text = "Hello, World!"
		m.TargetLang = "ja"
	})
}
