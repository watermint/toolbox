package screenshot

import (
	"errors"
	"github.com/kbinani/screenshot"
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

var (
	ErrorNoDisplay = errors.New("no display")
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
	numDisplays := screenshot.NumActiveDisplays()
	if numDisplays < 0 {
		c.UI().Error(z.ErrNoDisplay)
		return ErrorNoDisplay
	}
	displayId := z.DisplayId.Value()
	if displayId < 0 || displayId >= numDisplays {
		// This should not happen because displayId bound by RangeInt
		l.Error("Invalid display id", esl.Int("displayId", displayId), esl.Int("numDisplays", numDisplays))
		return ErrorNoDisplay
	}
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
	case errors.Is(err, ErrorNoDisplay), err == nil:
		return nil
	default:
		return err
	}
}
