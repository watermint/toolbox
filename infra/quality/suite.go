package quality

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_zap"
	"github.com/watermint/toolbox/recipe"
	"go.uber.org/zap"
)

func Suite(ctl app_control.Control) {
	resourceCheck(ctl)
	zapCheck(ctl)
	licenseCheck(ctl)
}

func resourceCheck(ctl app_control.Control) {
	l := ctl.Log()
	if app.Hash == "" {
		l.Error("Hash is empty")
		ctl.Abort(app_control.Reason(app_control.FatalResourceUnavailable))
		return
	}
	if app.BuilderKey == "" {
		l.Error("BuilderKey is empty")
		ctl.Abort(app_control.Reason(app_control.FatalResourceUnavailable))
		return
	}
	if app.Zap == "" {
		l.Error("Zap is empty")
		ctl.Abort(app_control.Reason(app_control.FatalResourceUnavailable))
		return
	}
}

func zapCheck(ctl app_control.Control) {
	l := ctl.Log()
	b, err := sc_zap.Unzap(ctl)
	if err != nil {
		l.Error("Unzap failed", zap.Error(err))
		ctl.Abort(app_control.Reason(app_control.FatalResourceUnavailable))
		return
	}
	var keys map[string]string
	err = json.Unmarshal(b, &keys)
	if err != nil {
		l.Error("Unable to unmarshal", zap.Error(err))
		ctl.Abort(app_control.Reason(app_control.FatalResourceUnavailable))
		return
	}
}

func licenseCheck(ctl app_control.Control) {
	l := ctl.Log()
	_, _, err := recipe.LoadLicense(ctl)
	if err != nil {
		l.Error("Unable to load license", zap.Error(err))
		ctl.Abort(app_control.Reason(app_control.FatalResourceUnavailable))
		return
	}
}
