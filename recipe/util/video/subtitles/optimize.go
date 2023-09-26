package subtitles

import (
	"github.com/asticode/go-astisub"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"os"
	"path/filepath"
	"time"
)

type Optimize struct {
	In           mo_path.ExistingFileSystemPath
	Out          mo_path.FileSystemPath
	OffsetMillis int64
}

func (z *Optimize) Preset() {
}

func (z *Optimize) Exec(c app_control.Control) error {
	s, err := astisub.OpenFile(z.In.Path())
	if err != nil {
		return err
	}
	s.Optimize()
	if z.OffsetMillis != 0 {
		s.Add(time.Duration(z.OffsetMillis) * time.Millisecond)
	}

	return s.Write(z.Out.Path())
}

func (z *Optimize) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("subtitles", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()
	f, err := qt_file.MakeTestFile("subtitles", "1\\n00:01:00.000 --> 00:02:00.000\\n watermint toolbox")
	if err != nil {
		return err
	}

	return rc_exec.Exec(c, &Optimize{}, func(r rc_recipe.Recipe) {
		m := r.(*Optimize)
		m.In = mo_path.NewExistingFileSystemPath(f)
		m.Out = mo_path.NewFileSystemPath(filepath.Join(p, "optimized.srt"))
	})
}
