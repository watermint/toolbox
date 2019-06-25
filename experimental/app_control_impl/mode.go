package app_control_impl

import "github.com/watermint/toolbox/app"

func isProduction() bool {
	return app.AppHash != ""
}
