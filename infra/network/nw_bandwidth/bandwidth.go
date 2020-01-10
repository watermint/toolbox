package nw_bandwidth

import (
	"github.com/watermint/bwlimit"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"io"
)

var (
	throttle     = bwlimit.NewBwlimit(0, false)
	currentLimit = 0
)

// Set bandwidth limit in Kilo Bytes per second.
func SetBandwidth(kps int) {
	app_root.Log().Debug("Bandwidth limit", zap.Int("kbs", kps))
	throttle.SetRateLimit(kps * 1024)
	currentLimit = kps
}

func WrapReader(r io.Reader) io.Reader {
	if currentLimit == 0 {
		return r
	}
	app_root.Log().Debug("Create new bandwidth limited reader", zap.Int("kbs", currentLimit))
	return throttle.Reader(r)
}

func WrapWriter(w io.Writer) io.Writer {
	if currentLimit == 0 {
		return w
	}
	app_root.Log().Debug("Create new bandwidth limited writer", zap.Int("kbs", currentLimit))
	return throttle.Writer(w)
}
