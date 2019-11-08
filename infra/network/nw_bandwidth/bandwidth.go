package nw_bandwidth

import (
	"github.com/watermint/bwlimit"
	"io"
)

var (
	throttle = bwlimit.NewBwlimit(0, false)
)

// Set bandwidth limit in Kilo Bytes per second.
func SetBandwidth(kps int) {
	throttle.SetRateLimit(kps * 1024)
}

func WrapReader(r io.Reader) io.Reader {
	return throttle.Reader(r)
}

func WrapWriter(w io.Writer) io.Writer {
	return throttle.Writer(w)
}
