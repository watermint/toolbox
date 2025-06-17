package build

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/doc/dc_index"
	"github.com/watermint/toolbox/infra/doc/dc_readme"
	"github.com/watermint/toolbox/infra/doc/dc_section"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type Readme struct {
	rc_recipe.RemarkSecret
	Path mo_path.FileSystemPath
}

func (z *Readme) Preset() {
}

func (z *Readme) genDoc(path string, doc string, c app_control.Control) error {
	l := c.Log()
	
	if c.Feature().IsTest() {
		l.Debug("Generating README to stdout (test mode)")
		out := es_stdout.NewDefaultOut(c.Feature())
		if out == nil {
			l.Error("Failed to create stdout output")
			return fmt.Errorf("failed to create stdout output")
		}
		_, err := fmt.Fprintln(out, doc)
		if err != nil {
			l.Error("Failed to write README to stdout", esl.Error(err))
			return err
		}
		return nil
	} else {
		l.Debug("Writing README to file", esl.String("path", path))
		err := os.WriteFile(path, []byte(doc), 0644)
		if err != nil {
			l.Error("Failed to write README file", esl.Error(err), esl.String("path", path))
			return err
		}
		l.Debug("README file written successfully", esl.String("path", path))
		return nil
	}
}

func (z *Readme) Exec(c app_control.Control) error {
	l := c.Log()
	l.Info("Generating README", esl.String("path", z.Path.Path()))
	
	// Add defensive error handling for CI environment
	defer func() {
		if r := recover(); r != nil {
			l.Error("README generation panicked", esl.Any("panic", r))
		}
	}()
	
	// Generate documentation sections with error handling
	sec := dc_readme.New(dc_index.MediaRepository, c.Messages(), false)
	if sec == nil {
		l.Error("Failed to create README sections")
		return fmt.Errorf("failed to create README sections")
	}
	
	doc := dc_section.Generate(dc_index.MediaRepository, dc_section.LayoutPage, c.Messages(), sec)
	if doc == "" {
		l.Error("Generated README document is empty")
		return fmt.Errorf("generated README document is empty")
	}
	
	l.Debug("README document generated", esl.Int("length", len(doc)))
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
