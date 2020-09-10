package build

import (
	"fmt"
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_readme"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Readme struct {
	rc_recipe.RemarkSecret
	Path mo_path.FileSystemPath
}

func (z *Readme) Preset() {
}

func (z *Readme) genDoc(path string, doc string, c app_control.Control) error {
	if c.Feature().IsTest() {
		out := es_stdout.NewDefaultOut(c.Feature())
		_, _ = fmt.Fprintln(out, doc)
		return nil
	} else {
		return ioutil.WriteFile(path, []byte(doc), 0644)
	}
}

func (z *Readme) Exec(c app_control.Control) error {
	l := c.Log()
	l.Info("Generating README", esl.String("path", z.Path.Path()))
	sec := dc_readme.New(false, "")
	doc := dc_section.Document(c.Messages(), sec...)

	return z.genDoc(z.Path.Path(), doc, c)
}

func (z *Readme) Test(c app_control.Control) error {
	path, err := qt_file.MakeTestFolder("readme", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(path)
	}()

	return rc_exec.Exec(c, &Readme{}, func(r rc_recipe.Recipe) {
		m := r.(*Readme)
		m.Path = mo_path.NewFileSystemPath(filepath.Join(path, "README.txt"))
	})
}
