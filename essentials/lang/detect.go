package lang

import (
	"github.com/cloudfoundry/jibber_jabber"
	"github.com/watermint/toolbox/essentials/log/es_fallback"
	"go.uber.org/zap"
)

// Detect & select language in select
func Detect(supported []Lang) Lang {
	l := es_fallback.Fallback()
	bcp47, err := jibber_jabber.DetectIETF()
	if err != nil {
		l.Debug("unable to detect language", zap.Error(err))
		return Default
	}

	return Select(bcp47, supported)
}
