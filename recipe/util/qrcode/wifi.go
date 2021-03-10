package qrcode

import (
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"math"
	"path/filepath"
)

// Generate Wifi configuration QR code
// @see https://github.com/zxing/zxing/blob/b1c85db64e0ef13e7d8a4c9de32bd94c76eea5d8/core/src/main/java/com/google/zxing/client/result/WifiResultParser.java#L22
//
// WIFI:T:[network type];S:[network SSID];P:[network password];H:[hidden?];;
type Wifi struct {
	ErrorCorrectionLevel mo_string.SelectString
	Hidden               mo_string.SelectString
	Mode                 mo_string.SelectString
	NetworkType          mo_string.SelectString
	Out                  mo_path.FileSystemPath
	Size                 mo_int.RangeInt
	Ssid                 string
	AskPassword          app_msg.Message
	OperationCancelled   app_msg.Message
}

func (z *Wifi) Preset() {
	z.ErrorCorrectionLevel.SetOptions(qrCodeErrorCorrectionLevelM, qrCodeErrorCorrectionLevels...)
	z.Hidden.SetOptions("", "", "true", "false")
	z.Mode.SetOptions(qrCodeEncodeAuto, qrCodeEncodes...)
	z.NetworkType.SetOptions("WPA", "WPA", "WEP", "")
	z.Size.SetRange(25, math.MaxInt16, 256)
}

func (z *Wifi) Exec(c app_control.Control) error {
	password, cancel := c.UI().AskSecure(z.AskPassword.With("SSID", z.Ssid))
	if cancel {
		c.UI().Info(z.OperationCancelled)
		return nil
	}

	wifiText := "WIFI:T:" + z.NetworkType.Value() + ";S:" + z.Ssid + ";P:" + password + ";H:" + z.Hidden.Value() + ";;"

	return createQrCodeImage(c.Log(), z.Out.Path(), wifiText, z.Size.Value(), z.ErrorCorrectionLevel.Value(), z.Mode.Value())
}

func (z *Wifi) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Wifi{}, func(r rc_recipe.Recipe) {
		m := r.(*Wifi)
		m.Out = mo_path.NewFileSystemPath(filepath.Join(c.Workspace().Report(), "out.png"))
		m.Ssid = "test"
	})
}
