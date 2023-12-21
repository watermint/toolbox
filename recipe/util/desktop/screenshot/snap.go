package screenshot

import (
	"errors"
	"github.com/kbinani/screenshot"
	screenshot2 "github.com/watermint/toolbox/domain/desktop/dd_screenshot"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"image/png"
	"os"
	"path/filepath"
)

type Snap struct {
	DisplayId    mo_int.RangeInt
	Path         mo_path.FileSystemPath
	ErrNoDisplay app_msg.Message
}

func (z *Snap) Preset() {
	z.DisplayId.SetRange(0, int64(screenshot.NumActiveDisplays()), 0)
}

func (z *Snap) Exec(c app_control.Control) error {
	l := c.Log()
	if err := screenshot2.CheckDisplayAvailability(z.DisplayId.Value()); err != nil {
		c.UI().Error(z.ErrNoDisplay)
		return err
	}
	displayId := z.DisplayId.Value()
	bounds := screenshot.GetDisplayBounds(displayId)
	img, err := screenshot.CaptureRect(bounds)
	if err != nil {
		l.Debug("Unable to capture screenshot", esl.Error(err))
		return err
	}
	f, err := os.Create(z.Path.Path())
	if err != nil {
		l.Debug("Unable to create file", esl.Error(err))
		return err
	}
	defer func() {
		_ = f.Close()
	}()

	if err := png.Encode(f, img); err != nil {
		l.Debug("Unable to encode image", esl.Error(err))
		return err
	}
	l.Debug("Screenshot saved", esl.String("path", z.Path.Path()))

	return nil
}

func (z *Snap) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("screenshot", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()
	err = rc_exec.ExecMock(c, &Snap{}, func(r rc_recipe.Recipe) {
		m := r.(*Snap)
		m.Path = mo_path.NewFileSystemPath(filepath.Join(f, "screenshot.png"))
	})
	switch {
	case errors.Is(err, screenshot2.ErrorNoDisplay), err == nil:
		return nil
	default:
		return err
	}
}
