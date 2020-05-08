package api_appkey

import (
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_resource"
	"github.com/watermint/toolbox/infra/security/sc_zap"
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
			return &resourceImpl{keys: keys}
		}
	}
	err = json.Unmarshal(kb, &keys)
	if err != nil {
		l.Debug("Skip loading app keys: unable to unmarshal resource", esl.Error(err))
		return &resourceImpl{keys: keys}
	}
	return &resourceImpl{keys: keys}
}

type resourceImpl struct {
	keys map[string]string
}

func (z resourceImpl) Key(scope string) (key, secret string) {
	var e bool
	if key, e = z.keys[scope+".key"]; !e {
		return "", ""
	}
	if secret, e = z.keys[scope+".secret"]; !e {
		return "", ""
	}
	return
}
