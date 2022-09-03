package app_apikey

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/security/sc_zap"
)

const (
	suffixKey    = ".key"
	suffixSecret = ".secret"
)

func Resolve(ctl app_control.Control, appKey string) (clientId, clientSecret string) {
	l := ctl.Log().With(esl.String("appKey", appKey))

	// Retrieve from extra
	extra := ctl.Feature().Extra()
	var e bool
	if clientId, e = extra.AppKey(appKey + suffixKey); e {
		if clientSecret, e = extra.AppKey(appKey + suffixSecret); e {
			return clientId, clientSecret
		}
	}

	// Retrieve from obfuscated resource
	keys := make(map[string]string)

	kb, err := sc_zap.Unzap(ctl)
	if err != nil {
		kb, err = app_resource.Bundle().Keys().Bytes("toolbox.appkeys")
		if err != nil {
			l.Debug("Skip loading app keys")
			return "", ""
		}
	}
	if err = json.Unmarshal(kb, &keys); err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		return "", ""
	}

	if clientId, e = keys[appKey+suffixKey]; e {
		if clientSecret, e = keys[appKey+suffixSecret]; e {
			return clientId, clientSecret
		}
	}
	l.Debug("Client ID/Secret not found")
	return "", ""
}
