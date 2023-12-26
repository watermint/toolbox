package qt_runtime

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_definitions"
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
	if app_definitions.BuildInfo.Hash == "" {
		l.Error("Hash is empty")
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
	if app_definitions.BuildInfo.Xap == "" {
		l.Error("BuilderKey is empty")
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
	if app_definitions.BuildInfo.Zap == "" {
		l.Error("Zap is empty")
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
	if !app_definitions.BuildInfo.Production {
		l.Error("The build is not for the production")
		app_exit.Abort(app_exit.FatalResourceUnavailable)
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
	_, err = app_resource.Bundle().Keys().Bytes("toolbox.appkeys")
	if err == nil {
		l.Debug("Dev app keys still exists", esl.Error(err))
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
