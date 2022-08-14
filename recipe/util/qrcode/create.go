package qrcode

import (
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/data/da_text"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"math"
	"os"
	"path/filepath"
)

type Create struct {
	Text                 da_text.TextInput
	ErrorCorrectionLevel mo_string.SelectString
	Mode                 mo_string.SelectString
	Size                 mo_int.RangeInt
	Out                  mo_path.FileSystemPath
}

func (z *Create) Preset() {
	z.ErrorCorrectionLevel.SetOptions(qrCodeErrorCorrectionLevelM, qrCodeErrorCorrectionLevels...)
	z.Mode.SetOptions(qrCodeEncodeAuto, qrCodeEncodes...)
	z.Size.SetRange(25, math.MaxInt16, 256)
}

func (z *Create) Exec(c app_control.Control) error {
	tx, err := z.Text.Content()
	if err != nil {
		return err
	}
	return createQrCodeImage(c.Log(), z.Out.Path(), string(tx), z.Size.Value(), z.ErrorCorrectionLevel.Value(), z.Mode.Value())
}

func (z *Create) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFile("qr", "watermint toolbox")
	if err != nil {
		return err
	}
	defer func() {
		_ = os.Remove(f)
	}()
	return rc_exec.Exec(c, &Create{}, func(r rc_recipe.Recipe) {
		m := r.(*Create)
		m.Out = mo_path.NewFileSystemPath(filepath.Join(c.Workspace().Report(), "out.png"))
		m.Text.SetFilePath(f)
	})
}
