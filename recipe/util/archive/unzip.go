package archive

import (
	"bytes"
	"errors"
	"github.com/watermint/toolbox/essentials/io/es_zip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Unzip struct {
	In  mo_path.ExistingFileSystemPath
	Out mo_path.FileSystemPath
}

func (z *Unzip) Preset() {
}

func (z *Unzip) Exec(c app_control.Control) error {
	return es_zip.Extract(c.Log(), z.In.Path(), z.Out.Path())
}

func (z *Unzip) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("zip", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(p)
	}()

	err = rc_exec.Exec(c, &Unzip{}, func(r rc_recipe.Recipe) {
		m := r.(*Unzip)
		m.In = mo_path.NewExistingFileSystemPath("unzip-test.zip")
		m.Out = mo_path.NewFileSystemPath(p)
	})
	if err != nil {
		return err
	}

	content, err := os.ReadFile(filepath.Join(p, "unzip-test/unzip-test.txt"))
	if err != nil {
		return err
	}
	if cp := bytes.Compare(content, []byte("unzip\n")); cp != 0 {
		c.Log().Error("Content mismatch", esl.ByteString("unarchived", content))
		return errors.New("content mismatch")
	}
	return nil
}
