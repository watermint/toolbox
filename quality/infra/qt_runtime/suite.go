package qt_runtime

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_exit"
	"github.com/watermint/toolbox/infra/security/sc_zap"
	"github.com/watermint/toolbox/recipe"
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
		l.Error("Unzap failed", es_log.Error(err))
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
	var keys map[string]string
	err = json.Unmarshal(b, &keys)
	if err != nil {
		l.Error("Unable to unmarshal", es_log.Error(err))
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
}

func checkLicense(ctl app_control.Control) {
	l := ctl.Log()
	_, _, _, err := recipe.LoadLicense(ctl)
	if err != nil {
		l.Error("Unable to load license", es_log.Error(err))
		app_exit.Abort(app_exit.FatalResourceUnavailable)
		return
	}
}
