package sc_zap

import (
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_obfuscate"
)

func Unzap(ctl app_control.Control) (b []byte, err error) {
	tas, err := ctl.Resource("toolbox.appkeys.secret")
	if err != nil {
		return nil, err
	}
	return sc_obfuscate.Deobfuscate(ctl.Log(), []byte(app.Zap), tas)
}
