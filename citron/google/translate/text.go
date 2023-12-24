package translate

import (
	"bufio"
	"bytes"
	"github.com/watermint/toolbox/domain/google/api/goog_auth"
	"github.com/watermint/toolbox/domain/google/api/goog_conn"
	"github.com/watermint/toolbox/domain/google/translate/service/sv_translate_v3"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
)

type Text struct {
	rc_recipe.RemarkSecret
	Peer      goog_conn.ConnGoogleTranslate
	Text      da_text.TextInput
	ProjectId string
	Source    mo_string.OptionalString
	Target    string
}

func (z *Text) Preset() {
	z.Peer.SetScopes(goog_auth.ScopeTranslateCloudTranslate)
}

func (z *Text) Exec(c app_control.Control) error {
	content, err := z.Text.Content()
	if err != nil {
		return err
	}

	tr := sv_translate_v3.New(z.ProjectId, z.Peer.Client())

	s := bufio.NewScanner(bytes.NewReader(content))
	for s.Scan() {
		line := s.Text()
		if result, err := tr.Translate(line, z.Source.Value(), z.Target); err != nil {
			return err
		} else {
			ui_out.TextOut(c, result)
		}
	}
	return s.Err()
}

func (z *Text) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("translate", "Hello, world!")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	return rc_exec.ExecMock(c, &Text{}, func(r rc_recipe.Recipe) {
		m := r.(*Text)
		m.Text.SetFilePath(f)
		m.Target = "ja"
	})
}
