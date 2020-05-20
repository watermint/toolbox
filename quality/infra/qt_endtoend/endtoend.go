package qt_endtoend

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/runtime/es_env"
	"github.com/watermint/toolbox/infra/app"
	"testing"
)

func IsSkipEndToEndTest() bool {
	if es_env.IsEnabled(app.EnvNameEndToEndSkipTest) {
		return true
	}

	// Recover: `testing: Short called before Init`
	defer func() {
		if r := recover(); r != nil {
			l := esl.Default()
			l.Debug("Recover", esl.Any("recover", r))
		}
	}()
	if testing.Short() {
		return true
	}
	return false
}
