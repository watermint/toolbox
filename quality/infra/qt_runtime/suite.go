package qt_runtime

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_exit"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/doc/dc_license"
	"github.com/watermint/toolbox/infra/security/sc_zap"
)

func Suite(ctl app_control.Control) {
	checkResource(ctl)
	checkZap(ctl)
	checkLicense(ctl)
}

func checkResource(ctl app_control.Control) {
	l := ctl.Log()
	if app.Hash == "" {
		l.Error("Hash is empty")
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
	if app.BuilderKey == "" {
		l.Error("BuilderKey is empty")
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
	if app.Zap == "" {
		l.Error("Zap is empty")
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
}

func checkZap(ctl app_control.Control) {
	l := ctl.Log()
	b, err := sc_zap.Unzap(ctl)
	if err != nil {
		l.Error("Unzap failed", esl.Error(err))
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
	var keys map[string]string
	err = json.Unmarshal(b, &keys)
	if err != nil {
		l.Error("Unable to unmarshal", esl.Error(err))
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
}

func checkLicense(ctl app_control.Control) {
	l := ctl.Log()
	licenseData, err := app_resource.Bundle().Data().Bytes("licenses.json")
	if err != nil {
		l.Error("License file not found")
		app_exit.Abort(app_exit.FatalResourceUnavailable)
	}
	licenses := &dc_license.Licenses{}
	err = json.Unmarshal(licenseData, licenses)
	if err != nil {
		l.Error("Unable to parse", esl.Error(err), esl.ByteString("data", licenseData))
		app_exit.Abort(app_exit.FatalResourceUnavailable)
	}
}
