package lang

import (
	"github.com/cloudfoundry/jibber_jabber"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
)

// Detect & select language in select
func Detect(supported []Lang) Lang {
	bcp47, err := jibber_jabber.DetectIETF()
	if err != nil {
		app_root.Log().Debug("unable to detect language", zap.Error(err))
		return Default
	}

	return Select(bcp47, supported)
}
