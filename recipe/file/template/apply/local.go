package apply

import (
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_filesystem"
	"github.com/watermint/toolbox/essentials/file/es_filesystem_local"
	"github.com/watermint/toolbox/essentials/file/es_template"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io"
	"os"
	"path/filepath"
)

type Local struct {
	Template mo_path.ExistingFileSystemPath
	Path     mo_path.FileSystemPath
}

func (z *Local) Preset() {
}

func (z *Local) putFile(path es_filesystem.Path, f io.ReadSeeker) error {
	l := esl.Default().With(esl.String("path", path.Path()))
	l.Debug("Put the file")
	w, err := os.Create(path.Path())
	if err != nil {
		l.Debug("Unable to put the file", esl.Error(err))
		return err
	}
	defer func() {
		_ = w.Close()
	}()

	_, err = io.Copy(w, f)
	return err
}

func (z *Local) Exec(c app_control.Control) error {
	tmpl, err := os.ReadFile(z.Template.Path())
	if err != nil {
		return err
	}
	tmplData := es_template.Root{}
	if err := json.Unmarshal(tmpl, &tmplData); err != nil {
		return err
	}

	fs := es_filesystem_local.NewFileSystem()
	ap := es_template.NewApply(fs,
		es_template.ApplyOpts{
			HandlerPutFile: z.putFile,
		},
	)
	return ap.Apply(es_filesystem_local.NewPath(z.Path.Path()), tmplData)
}

func (z *Local) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("tmpl", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	tmplPath := filepath.Join(f, "template.json")
	err = os.WriteFile(tmplPath, []byte(`{"folders":[
               {"name":"test-001", "files":[
                   {"name":"test.txt", "content":"`+app_definitions.BuildId+`"}
                 ]
               }
             ]}`), 0644)
	if err != nil {
		return err
	}

	err = rc_exec.Exec(c, &Local{}, func(r rc_recipe.Recipe) {
		m := r.(*Local)
		m.Template = mo_path.NewExistingFileSystemPath(tmplPath)
		m.Path = mo_path.NewFileSystemPath(f)
	})
	if err != nil {
		return err
	}

	testData, err := os.ReadFile(filepath.Join(f, "test-001", "test.txt"))
	if err != nil {
		return err
	}
	if string(testData) != app_definitions.BuildId {
		c.Log().Warn("Invalid data", esl.ByteString("data", testData), esl.String("expected", app_definitions.BuildId))
		return errors.New("invalid data")
	}
	return nil
}
