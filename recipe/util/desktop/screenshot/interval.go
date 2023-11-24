package screenshot

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/kbinani/screenshot"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"image"
	"image/png"
	"os"
	"path/filepath"
	"reflect"
	"text/template"
	"time"
)

type Interval struct {
	Interval       int64
	Path           mo_path.FileSystemPath
	Count          int64
	NamePattern    string
	SkipIfNoChange bool
	DisplayId      mo_int.RangeInt
	ErrNoDisplay   app_msg.Message
	ProgressSnap   app_msg.Message
	ProgressSkip   app_msg.Message
}

func (z *Interval) Preset() {
	z.NamePattern = "{{.Sequence}}_{{.Timestamp}}.png"
	z.DisplayId.SetRange(0, int64(screenshot.NumActiveDisplays()), 0)
	z.Count = -1
	z.Interval = 10
}

func snapName(nameTmpl *template.Template, displayId int, seq int64, bounds image.Rectangle) (string, error) {
	now := time.Now()
	nowUtc := now.UTC()
	nameValues := map[string]string{
		"Date":          now.Format("2006-01-02"),
		"DateUTC":       nowUtc.Format("2006-01-02"),
		"DisplayHeight": fmt.Sprintf("%d", bounds.Size().Y),
		"DisplayId":     fmt.Sprintf("%d", displayId),
		"DisplayWidth":  fmt.Sprintf("%d", bounds.Size().X),
		"DisplayX":      fmt.Sprintf("%d", bounds.Min.X),
		"DisplayY":      fmt.Sprintf("%d", bounds.Min.Y),
		"Sequence":      fmt.Sprintf("%05d", seq),
		"Time":          now.Format("15-04-05"),
		"TimeUTC":       nowUtc.Format("15-04-05"),
		"Timestamp":     now.Format("20060102-150405"),
		"TimestampUTC":  now.Format("20060102-150405Z"),
	}
	var nameBuf bytes.Buffer
	err := nameTmpl.Execute(&nameBuf, nameValues)
	if err != nil {
		return "", err
	}
	return nameBuf.String(), nil
}

func (z *Interval) shouldSave(lastImg, currentImg *image.RGBA) bool {
	if !z.SkipIfNoChange {
		return true
	}
	if lastImg == nil {
		return true
	}

	// compare images
	return !reflect.DeepEqual(lastImg.Pix, currentImg.Pix)
}

func (z *Interval) Exec(c app_control.Control) error {
	l := c.Log()
	if err := checkDisplayAvailability(z.DisplayId.Value()); err != nil {
		c.UI().Error(z.ErrNoDisplay)
		return err
	}
	bounds := screenshot.GetDisplayBounds(z.DisplayId.Value())

	var seq int64 = 0
	nameTmpl, err := template.New("name").Parse(z.NamePattern)
	var lastImg *image.RGBA
	if err != nil {
		return err
	}
	for {
		seq++
		name, err := snapName(nameTmpl, z.DisplayId.Value(), seq, bounds)
		if err != nil {
			l.Debug("Unable to create name", esl.Error(err))
			return err
		}
		c.UI().Progress(z.ProgressSnap.With("Name", name).With("Sequence", seq).With("Timestamp", time.Now().Format("2006-01-02 15:04:05")))
		img, err := screenshot.CaptureRect(bounds)
		if err != nil {
			l.Debug("Unable to capture screenshot", esl.Error(err))
			return err
		}

		if !z.shouldSave(lastImg, img) {
			c.UI().Progress(z.ProgressSkip.With("Name", name).With("Sequence", seq).With("Timestamp", time.Now().Format("2006-01-02 15:04:05")))
		} else {
			f, err := os.Create(filepath.Join(z.Path.Path(), name))
			if err != nil {
				l.Debug("Unable to create file", esl.Error(err))
				return err
			}
			if err := png.Encode(f, img); err != nil {
				l.Debug("Unable to encode image", esl.Error(err))
				return err
			}
			l.Debug("Screenshot saved", esl.String("path", z.Path.Path()))
			if err := f.Close(); err != nil {
				l.Debug("Unable to close file", esl.Error(err))
				return err
			}
			if err != nil {
				return err
			}
		}
		lastImg = img

		if z.Count > 0 && seq >= z.Count {
			c.Log().Debug("Reached the count", esl.Int64("count", z.Count))
			return nil
		}
		time.Sleep(time.Duration(z.Interval) * time.Second)
	}
}

func (z *Interval) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("screenshot", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()
	err = rc_exec.ExecMock(c, &Interval{}, func(r rc_recipe.Recipe) {
		m := r.(*Interval)
		m.Path = mo_path.NewFileSystemPath(f)
		m.Interval = 10
		m.Count = 1
	})
	switch {
	case errors.Is(err, ErrorNoDisplay), err == nil:
		return nil
	default:
		return err
	}
}
