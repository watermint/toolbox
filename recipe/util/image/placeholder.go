package image

import (
	"errors"
	"github.com/watermint/toolbox/essentials/islet/egraphic/ecolor"
	"github.com/watermint/toolbox/essentials/islet/egraphic/edraw"
	"github.com/watermint/toolbox/essentials/islet/egraphic/egeom"
	"github.com/watermint/toolbox/essentials/islet/egraphic/eimage"
	"github.com/watermint/toolbox/essentials/islet/egraphic/etext"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"io/ioutil"
	"os"
	"path/filepath"
)

type Placeholder struct {
	Width                 int
	Height                int
	Color                 string
	Text                  mo_string.OptionalString
	TextColor             string
	TextPosition          string
	TextAlign             mo_string.SelectString
	FontSize              int
	FontPath              mo_string.OptionalString
	Path                  mo_path.FileSystemPath
	ErrorFontPathRequired app_msg.Message
	ErrorCantLoadFont     app_msg.Message
}

func (z *Placeholder) Preset() {
	z.Width = 640
	z.Height = 400
	z.Color = "white"
	z.TextColor = "black"
	z.FontSize = 12
	z.TextPosition = "center"
	z.TextAlign.SetOptions(
		"left",
		"left", "center", "right",
	)
}

func (z *Placeholder) Exec(c app_control.Control) error {
	ui := c.UI()
	bgColor, oc := ecolor.ParseColor(z.Color)
	if oc.IsError() {
		return oc.Cause()
	}

	img := eimage.NewRgba(z.Width, z.Height)
	imgDraw := edraw.NewImageDrawer(img)
	imgDraw.FillRectangle(img.Bounds(), bgColor)

	if z.Text.IsExists() {
		if !z.FontPath.IsExists() {
			ui.Error(z.ErrorFontPathRequired)
			return errors.New("font path required to draw text")
		}
		fontData, err := ioutil.ReadFile(z.FontPath.Value())
		if err != nil {
			ui.Error(z.ErrorCantLoadFont.With("Path", z.FontPath.Value()).With("Error", err))
			return err
		}
		ttf, oc := etext.NewTrueTypeParse(fontData)
		if oc.IsError() {
			ui.Error(z.ErrorCantLoadFont.With("Path", z.FontPath.Value()).With("Error", oc.Cause()))
			return oc.Cause()
		}

		txtColor, oc := ecolor.ParseColor(z.TextColor)
		if oc.IsError() {
			return oc.Cause()
		}
		var txtAlign etext.Alignment
		switch z.TextAlign.Value() {
		case "center":
			txtAlign = etext.AlignCenter
		case "right":
			txtAlign = etext.AlignRight
		default:
			txtAlign = etext.AlignLeft
		}
		txtStyle := etext.NewStyle(ttf.WithSize(z.FontSize), txtColor).WithAlignment(txtAlign)
		txtPos, oc := egeom.ParsePosition(z.TextPosition)
		if oc.IsError() {
			return oc.Cause()
		}

		imgDraw.DrawString(
			txtPos.Locate(
				img.Bounds(),
				txtStyle.Bound(z.Text.Value()),
				egeom.NewPaddingFixed(z.FontSize, z.FontSize),
			),
			z.Text.Value(),
			txtStyle,
		)
	}

	if oc := img.ExportTo(eimage.FormatPng, z.Path.Path()); oc.IsError() {
		return oc.Cause()
	}
	return nil
}

func (z *Placeholder) Test(c app_control.Control) error {
	p, err := qt_file.MakeTestFolder("placeholder", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(p)
	}()

	return rc_exec.Exec(c, &Placeholder{}, func(r rc_recipe.Recipe) {
		m := r.(*Placeholder)
		m.Path = mo_path.NewFileSystemPath(filepath.Join(p, "test.png"))
		//		m.Text = mo_string.NewOptional("watermint toolbox")
		m.Color = "marker(b18)"
		m.TextColor = "marker(w00)"
		m.TextPosition = "top-right"
		m.TextAlign.SetSelect("right")
	})
}
