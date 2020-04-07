package api_appkey

import (
	"encoding/json"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/security/sc_zap"
	"go.uber.org/zap"
)

type Resource interface {
	// Key/secret
	Key(scope string) (key, secret string)
}

func New(ctl app_control.Control) Resource {
	return &resourceImpl{
		ctl: ctl,
	}
}

type resourceImpl struct {
	ctl  app_control.Control
	keys map[string]string
}

func (z *resourceImpl) load() {
	l := z.ctl.Log()
	kb, err := sc_zap.Unzap(z.ctl)
	if err != nil {
		kb, err = z.ctl.Resource("toolbox.appkeys")
		if err != nil {
			l.Debug("Skip loading app keys")
			return
		}
	}
	err = json.Unmarshal(kb, &z.keys)
	if err != nil {
		l.Debug("Skip loading app keys: unable to unmarshal resource", zap.Error(err))
		return
	}
}

func (z *resourceImpl) Key(scope string) (key, secret string) {
	if z.keys == nil {
		z.load()
	}

	var e bool
	if key, e = z.keys[scope+".key"]; !e {
		return "", ""
	}
	if secret, e = z.keys[scope+".secret"]; !e {
		return "", ""
	}
	return
}
