package nw_bandwidth

import (
	"github.com/watermint/bwlimit"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"io"
)

var (
	throttle = bwlimit.NewBwlimit(0, false)
)

// Set bandwidth limit in Kilo Bytes per second.
func SetBandwidth(kps int) {
	app_root.Log().Debug("Bandwidth limit", zap.Int("kbs", kps))
	throttle.SetRateLimit(kps * 1024)
}

func WrapReader(r io.Reader) io.Reader {
	return throttle.Reader(r)
}

func WrapWriter(w io.Writer) io.Writer {
	return throttle.Writer(w)
}
