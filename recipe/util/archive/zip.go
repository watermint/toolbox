package archive

import (
	"bytes"
	"errors"
	"github.com/watermint/toolbox/essentials/file/es_zip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
)

type Zip struct {
	rc_recipe.RemarkTransient
	Target  mo_path.ExistingFileSystemPath
	Out     mo_path.FileSystemPath
	Comment mo_string.OptionalString
}

func (z *Zip) Preset() {
}

func (z *Zip) Exec(c app_control.Control) error {
	return es_zip.CompressPath(z.Out.Path(), z.Target.Path(), z.Comment.Value())
}

func (z *Zip) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("zip", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(p)
	}()

	arcPath := filepath.Join(p, "zip-test.zip")
	err = rc_exec.Exec(c, &Zip{}, func(r rc_recipe.Recipe) {
		m := r.(*Zip)
		m.Target = mo_path.NewExistingFileSystemPath(".")
		m.Out = mo_path.NewFileSystemPath(arcPath)
	})
	if err != nil {
		return err
	}

	pp := filepath.Join(p, "unzip")
	err = rc_exec.Exec(c, &Unzip{}, func(r rc_recipe.Recipe) {
		m := r.(*Unzip)
		m.In = mo_path.NewExistingFileSystemPath(arcPath)
		m.Out = mo_path.NewFileSystemPath(pp)
	})
	if err != nil {
		return err
	}

	expected, err := os.ReadFile("zip.go")
	if err != nil {
		return err
	}
	content, err := os.ReadFile(filepath.Join(pp, "zip.go"))
	if err != nil {
		return err
	}

	if cp := bytes.Compare(expected, content); cp != 0 {
		c.Log().Error("content mismatch", esl.ByteString("content", content), esl.ByteString("expected", expected))
		return errors.New("content mismatch")
	}

	return nil
}
