package stage

import (
	"errors"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/text/es_encoding"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"golang.org/x/text/transform"
	"os"
	"strings"
)

type Encoding struct {
	rc_recipe.RemarkSecret
	Peer     dbx_conn.ConnScopedIndividual
	Path     mo_path.DropboxPath
	Name     string
	Encoding string
}

func (z *Encoding) Preset() {
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
}

func (z *Encoding) Exec(c app_control.Control) error {
	l := c.Log()
	ec := es_encoding.SelectEncoding(z.Encoding)
	if ec == nil {
		l.Error("No encoding found", esl.String("Encoding", z.Encoding))
		return errors.New("no encoding found")
	}
	en, el, err := transform.String(ec.NewEncoder(), z.Name)
	if err != nil {
		return err
	}
	en = strings.ReplaceAll(en, "\u0000", "?")
	l.Info("Convert", esl.Int("origLength", len(z.Name)), esl.Int("convertedLength", el))
	up := z.Path.ChildPath(en)

	f, err := qt_file.MakeTestFile("encoding", z.Encoding)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()

	ulf, err := sv_file_content.NewUpload(z.Peer.Context(), sv_file_content.UseCustomFileName(true)).Add(up, f)
	if err != nil {
		return err
	}
	l.Info("Upload completed", esl.Any("meta", ulf))
	return nil
}

func (z *Encoding) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Encoding{}, func(r rc_recipe.Recipe) {
		m := r.(*Encoding)
		m.Path = qtr_endtoend.NewTestDropboxFolderPath("encoding")
		m.Encoding = "shift-jis"
		m.Name = "エンコーディング"
	})
}
