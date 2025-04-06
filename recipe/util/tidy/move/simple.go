package move

import (
	"errors"
	"os"
	"path/filepath"

	"github.com/watermint/toolbox/essentials/file/es_filemove"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
)

type Simple struct {
	IncludeSystemFiles bool
	ExcludeFolders     bool
	Preview            bool
	Src                mo_path.ExistingFileSystemPath
	Dst                mo_path.FileSystemPath
	ProgressMoving     app_msg.Message
	ProgressSkip       app_msg.Message
	ErrorUnableToMove  app_msg.Message
}

func (z *Simple) Preset() {

}

func (z *Simple) move(c app_control.Control) (err error) {
	l := c.Log().With(esl.String("targetPath", z.Src.Path()))
	ui := c.UI()

	entries, err := os.ReadDir(z.Src.Path())
	if err != nil {
		l.Debug("Unable to read path", esl.Error(err))
		return err
	}

	dstInfo, err := os.Lstat(z.Dst.Path())
	if err == nil {
		l.Debug("The dest folder exists")
		if !dstInfo.IsDir() {
			l.Debug("The dest path is not a folder")
			return errors.New("the dest path is not a folder")
		}
	} else if os.IsNotExist(err) {
		l.Debug("The dest folder is not exist, try create folders")
		err = os.MkdirAll(z.Dst.Path(), 0755)
		if err != nil {
			l.Debug("Unable to create folder", esl.Error(err))
			return err
		}
	} else {
		l.Debug("Unknown error for dest folder", esl.Error(err))
		return err
	}

	var lastErr error
	for _, e := range entries {
		d := filepath.Join(z.Dst.Path(), e.Name())
		s := filepath.Join(z.Src.Path(), e.Name())
		ll := l.With(esl.String("sourceEntry", s), esl.String("destEntry", d))
		if e.IsDir() && z.ExcludeFolders {
			ll.Debug("Skip folder", esl.String("name", e.Name()))
			continue
		}
		if !z.IncludeSystemFiles && es_filepath.IsSystemFile(s) {
			ll.Debug("Skip system files", esl.String("name", e.Name()))
			continue
		}

		ui.Progress(z.ProgressMoving.With("Source", s).With("Dest", d))
		if z.Preview {
			ui.Progress(z.ProgressSkip.With("Source", s).With("Dest", d))
			continue
		}

		err = es_filemove.Move(s, d)
		if err != nil {
			ui.Error(z.ErrorUnableToMove.With("Source", s).With("Dest", d).With("Error", err))
			ll.Debug("Unable to move", esl.Error(err))
			lastErr = err
			continue
		}
	}
	return lastErr
}

func (z *Simple) Exec(c app_control.Control) error {
	return z.move(c)
}

func (z *Simple) Test(c app_control.Control) error {
	src, err := qt_file.MakeTestFolder("source", true)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(src)
	}()
	dst, err := qt_file.MakeTestFolder("dest", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(dst)
	}()

	return rc_exec.ExecMock(c, &Simple{}, func(r rc_recipe.Recipe) {
		m := r.(*Simple)
		m.Src = mo_path.NewExistingFileSystemPath(src)
		m.Dst = mo_path.NewFileSystemPath(dst)
	})
}
