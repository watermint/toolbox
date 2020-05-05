package nw_bandwidth

import (
	"github.com/watermint/bwlimit"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"io"
)

var (
	throttle     = bwlimit.NewBwlimit(0, false)
	currentLimit = 0
)

// Set bandwidth limit in Kilo Bytes per second.
func SetBandwidth(kps int) {
	es_log.Default().Debug("Bandwidth limit", es_log.Int("kbs", kps))
	throttle.SetRateLimit(kps * 1024)
	currentLimit = kps
}

func WrapReader(r io.Reader) io.Reader {
	if currentLimit == 0 {
		return r
	}
	es_log.Default().Debug("Create new bandwidth limited reader", es_log.Int("kbs", currentLimit))
	return throttle.Reader(r)
}

func WrapWriter(w io.Writer) io.Writer {
	if currentLimit == 0 {
		return w
	}
	es_log.Default().Debug("Create new bandwidth limited writer", es_log.Int("kbs", currentLimit))
	return throttle.Writer(w)
}
