package api_appkey

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

type Resource interface {
	// Key/secret
	Key(scope string) (key, secret string)
}

func New(ctl app_control.Control) Resource {
	keys := make(map[string]string)

	l := ctl.Log()
	kb, err := sc_zap.Unzap(ctl)
	if err != nil {
		kb, err = app_resource.Bundle().Keys().Bytes("toolbox.appkeys")
		if err != nil {
			l.Debug("Skip loading app keys")
			return &resourceImpl{keys: keys, ctl: ctl}
		}
	}
	err = json.Unmarshal(kb, &keys)
	if err != nil {
		l.Debug("Skip loading app keys: unable to unmarshal resource", esl.Error(err))
		return &resourceImpl{keys: keys, ctl: ctl}
	}
	return &resourceImpl{keys: keys, ctl: ctl}
}

type resourceImpl struct {
	keys map[string]string
	ctl  app_control.Control
}

func (z resourceImpl) Key(scope string) (key, secret string) {
	extra := z.ctl.Feature().Extra()
	var e bool
	if key, e = extra.AppKey(scope + suffixKey); e {
		if secret, e = extra.AppKey(scope + suffixSecret); e {
			return
		}
	}
	if key, e = z.keys[scope+suffixKey]; !e {
		return "", ""
	}
	if secret, e = z.keys[scope+suffixSecret]; !e {
		return "", ""
	}
	return
}
