package lang

import (
	"github.com/cloudfoundry/jibber_jabber"
	"github.com/watermint/toolbox/essentials/log/esl"
)

// Detect & select language in select
func Detect(supported []Lang) Lang {
	l := esl.Default()
	bcp47, err := jibber_jabber.DetectIETF()
	if err != nil {
		l.Debug("unable to detect language", esl.Error(err))
		return Default
	}

	return Select(bcp47, supported)
}
