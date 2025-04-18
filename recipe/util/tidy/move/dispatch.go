package move

import (
	"bytes"
	"fmt"
	"html/template"
	"os"
	"path/filepath"
	"regexp"
	"strings"

	"github.com/watermint/toolbox/essentials/file/es_filemove"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/feed/fd_file"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/report/rp_writer_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type MsgLocal struct {
	ErrorUnableToReadSourcePath       app_msg.Message
	ErrorUnableToMove                 app_msg.Message
	ErrorInvalidSourceFileNamePattern app_msg.Message
	ErrorInvalidDestPathPattern       app_msg.Message
	ErrorInvalidDestNamePattern       app_msg.Message
	ExecRule                          app_msg.Message
	ExecPreview                       app_msg.Message
	ExecMove                          app_msg.Message
}

var (
	MLocal = app_msg.Apply(&MsgLocal{}).(*MsgLocal)
)

func newNamePattern(pattern string) (*namePattern, error) {
	srcRe, err := regexp.Compile(pattern)
	if err != nil {
		return nil, err
	}
	return &namePattern{
		pattern: pattern,
		re:      srcRe,
	}, nil
}

type namePattern struct {
	pattern string
	re      *regexp.Regexp
}

func (z *namePattern) Match(name string) bool {
	return z.re.MatchString(name)
}

func (z *namePattern) MatchValues(name string) map[string]string {
	data := make(map[string]string)
	sm := z.re.FindStringSubmatch(name)
	for i, s := range sm {
		data[fmt.Sprintf("M%d", i)] = s
	}
	return data
}

func (z *namePattern) Compile(name, dstPattern string) (string, error) {
	data := z.MatchValues(name)
	t, err := template.New("").Parse(dstPattern)
	if err != nil {
		return "", err
	}
	buf := &bytes.Buffer{}
	if err = t.Execute(buf, data); err != nil {
		return "", err
	}
	return buf.String(), nil
}

type DispatchPattern struct {
	Suffix            string `json:"suffix"`
	SourcePath        string `json:"source_path"`
	SourceFilePattern string `json:"source_file_pattern"`
	DestPathPattern   string `json:"dest_path_pattern"`
	DestFilePattern   string `json:"dest_file_pattern"`
}

func (z *DispatchPattern) preview(src, dst string, c app_control.Control) error {
	ui := c.UI()
	ui.Info(MLocal.ExecPreview.With("Src", src).With("Dest", dst))
	return nil
}

func (z *DispatchPattern) move(src, dst string, c app_control.Control) error {
	l := c.Log().With(esl.String("src", src), esl.String("dst", dst))
	ui := c.UI()

	dstPath := filepath.Dir(dst)
	l.Debug("Prepare directory", esl.String("dstPath", dstPath))
	if err := os.MkdirAll(dstPath, 0755); err != nil {
		l.Debug("Unable to create a directory", esl.Error(err), esl.String("dstPath", dstPath))
		ui.Error(MLocal.ExecMove.With("Src", src).With("Dst", dst).With("Error", err))
		return err
	}

	l.Debug("Moving")
	err := es_filemove.Move(src, dst)
	if err != nil {
		l.Debug("Unable to move", esl.Error(err))
		ui.Error(MLocal.ExecMove.With("Src", src).With("Dst", dst).With("Error", err))
		return err
	}
	ui.Progress(MLocal.ExecMove.With("Src", src).With("Dst", dst))
	return nil
}

func (z *DispatchPattern) Exec(c app_control.Control, op func(src, dst string, c app_control.Control) error) error {
	ui := c.UI()
	l := c.Log()

	ui.Progress(MLocal.ExecRule.
		With("Suffix", z.Suffix).
		With("SourcePath", z.SourcePath).
		With("SourceFile", z.SourceFilePattern).
		With("DistPath", z.DestPathPattern).
		With("DestFile", z.DestFilePattern))

	srcPattern, err := newNamePattern(z.SourceFilePattern)
	if err != nil {
		ui.Error(MLocal.ErrorInvalidSourceFileNamePattern.
			With("Pattern", srcPattern).
			With("Error", err))
		return err
	}

	srcPath, err := es_filepath.FormatPathWithPredefinedVariables(z.SourcePath)
	if err != nil {
		return err
	}
	entries, err := os.ReadDir(srcPath)
	if err != nil {
		ui.Error(MLocal.ErrorUnableToReadSourcePath.With("Path", srcPath))
		return err
	}

	for _, entry := range entries {
		if !srcPattern.Match(entry.Name()) {
			l.Debug("Skip unmatched file",
				esl.String("pattern", z.SourceFilePattern),
				esl.String("name", entry.Name()))
			continue
		}
		ext := strings.ToLower(filepath.Ext(entry.Name()))
		if "."+strings.ToLower(z.Suffix) != ext {
			l.Debug("Skip unmatched suffix",
				esl.String("suffix", z.Suffix),
				esl.String("name", entry.Name()))
			continue
		}

		matchValues := srcPattern.MatchValues(entry.Name())
		matchPairs := make([]es_filepath.Pair, 0)
		for k, v := range matchValues {
			matchPairs = append(matchPairs, es_filepath.Pair{
				Key:   k,
				Value: v,
			})
		}
		destPath, err := es_filepath.FormatPathWithPredefinedVariables(z.DestPathPattern, matchPairs...)
		if err != nil {
			ui.Error(MLocal.ErrorInvalidDestPathPattern.
				With("Path", z.DestPathPattern).
				With("Error", err))
			return err
		}
		destName, err := srcPattern.Compile(entry.Name(), z.DestFilePattern)
		if err != nil {
			ui.Error(MLocal.ErrorInvalidDestNamePattern.
				With("Name", z.DestFilePattern).
				With("Error", err))
			return err
		}

		dest := filepath.Join(destPath, destName+"."+z.Suffix)
		src := filepath.Join(srcPath, entry.Name())

		if err := op(src, dest, c); err != nil {
			return err
		}
	}
	return err
}

type Dispatch struct {
	rc_recipe.RemarkIrreversible
	File    fd_file.RowFeed
	Preview bool
}

func (z *Dispatch) Preset() {
	z.File.SetModel(&DispatchPattern{})
}

func (z *Dispatch) Exec(c app_control.Control) error {
	return z.File.EachRow(func(m interface{}, rowIndex int) error {
		p := m.(*DispatchPattern)
		// ignore errors
		if z.Preview {
			p.Exec(c, p.preview)
		} else {
			p.Exec(c, p.move)
		}
		return nil
	})
}

func (z *Dispatch) Test(c app_control.Control) error {
	src, err := qt_file.MakeTestFolder("src", false)
	if err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(src)
	}()
	dst, err := qt_file.MakeTestFolder("dst", false)
	if err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(dst)
	}()

	name := "TBX-2020-04-01.txt"
	srcFile := filepath.Join(src, name)

	err = os.WriteFile(srcFile, []byte(app_definitions.BuildId), 0644)
	if err != nil {
		return err
	}
	lp := &DispatchPattern{
		Suffix:            "txt",
		SourcePath:        src,
		SourceFilePattern: `TBX-(\d{4})-(\d{2})-(\d{2})`,
		DestPathPattern:   dst + "/{{.M1}}",
		DestFilePattern:   "{{.M1}}-{{.M2}}-{{.M3}}",
	}
	cw := rp_writer_impl.NewCsvWriter("local", c)
	if err = cw.Open(c, &DispatchPattern{}); err != nil {
		return err
	}
	cw.Row(lp)
	cw.Close()
	filePath := filepath.Join(c.Workspace().Report(), "local.csv")

	// Preview
	err = rc_exec.Exec(c, &Dispatch{}, func(r rc_recipe.Recipe) {
		m := r.(*Dispatch)
		m.File.SetFilePath(filePath)
		m.Preview = true
	})
	if err != nil {
		return err
	}

	// Move
	err = rc_exec.Exec(c, &Dispatch{}, func(r rc_recipe.Recipe) {
		m := r.(*Dispatch)
		m.File.SetFilePath(filePath)
		m.Preview = false
	})
	if err != nil {
		return err
	}
	_, err = os.Lstat(filepath.Join(dst, "2020", "2020-04-01.txt"))
	if err != nil {
		return err
	}
	return nil
}
