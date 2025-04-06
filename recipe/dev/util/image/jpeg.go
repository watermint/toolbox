package image

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"math"
	"math/rand"
	"os"
	"path/filepath"

	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

const (
	maxJpegResolution = 65535
)

type Jpeg struct {
	rc_recipe.RemarkSecret
	Path       mo_path.FileSystemPath
	NamePrefix string
	Width      mo_int.RangeInt
	Height     mo_int.RangeInt
	Quality    mo_int.RangeInt
	Count      mo_int.RangeInt
	Seed       int
	Created    app_msg.Message
}

func (z *Jpeg) Preset() {
	z.Width.SetRange(1, maxJpegResolution, 1920)
	z.Height.SetRange(1, maxJpegResolution, 1080)
	z.Quality.SetRange(1, 100, jpeg.DefaultQuality)
	z.Count.SetRange(1, math.MaxInt16, 10)
	z.Seed = 1
	z.NamePrefix = "test_image"
}

func (z *Jpeg) create(rs *rand.Rand, index int, c app_control.Control) error {
	w, h := z.Width.Value(), z.Height.Value()
	q := z.Quality.Value()
	l := c.Log().With(esl.Int("index", index), esl.Int("width", w), esl.Int("height", h), esl.Int("quality", q))
	plane := image.NewCMYK(image.Rectangle{
		Min: image.Point{},
		Max: image.Point{X: w, Y: h},
	})
	l.Debug("Creating image plane")
	for x := 0; x < w; x++ {
		for y := 0; y < h; y++ {
			plane.Set(x, y, color.CMYK{
				C: uint8(rs.Intn(math.MaxUint8)),
				M: uint8(rs.Intn(math.MaxUint8)),
				Y: uint8(rs.Intn(math.MaxUint8)),
				K: uint8(rs.Intn(math.MaxUint8)),
			})
		}
	}

	name := fmt.Sprintf("%s_%d.jpg", z.NamePrefix, index)
	path := filepath.Join(z.Path.Path(), name)
	l.Debug("Creating file", esl.String("path", path))
	f, err := os.Create(path)
	if err != nil {
		l.Debug("Unable to create file", esl.Error(err))
		return err
	}
	err = jpeg.Encode(f, plane, &jpeg.Options{Quality: q})
	if err != nil {
		l.Debug("Unable to encode image", esl.Error(err))
		return err
	}
	c.UI().Progress(z.Created.With("Index", index).With("Total", z.Count.Value()))
	return nil
}

func (z *Jpeg) Exec(c app_control.Control) error {
	rs := rand.New(rand.NewSource(int64(z.Seed)))

	for i := 0; i < z.Count.Value(); i++ {
		if err := z.create(rs, i, c); err != nil {
			return err
		}
	}
	return nil
}

func (z *Jpeg) Test(c app_control.Control) error {
	tmp, err := os.MkdirTemp("", "jpeg")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(tmp)
	}()
	return rc_exec.Exec(c, &Jpeg{}, func(r rc_recipe.Recipe) {
		m := r.(*Jpeg)
		m.Width.SetValue(100)
		m.Height.SetValue(100)
		m.Count.SetValue(3)
		m.Path = mo_path.NewFileSystemPath(tmp)
	})
}
