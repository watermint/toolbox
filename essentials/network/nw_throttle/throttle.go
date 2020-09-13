package nw_throttle

import (
	"github.com/watermint/toolbox/essentials/network/nw_ratelimit"
)

func Throttle(hash, endpoint string, f func()) {
	nw_ratelimit.WaitIfRequired(hash, endpoint)
	f()
}
