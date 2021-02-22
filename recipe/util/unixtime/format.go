package unixtime

import (
	"errors"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/ui_out"
	"sort"
	"time"
)

var (
	supportedFormats = map[string]string{
		"iso8601":      time.RFC3339,
		"rfc3339":      time.RFC3339,
		"rfc3339_nano": time.RFC3339Nano,
		"rfc822":       time.RFC822,
		"rfc822z":      time.RFC822Z,
		"rfc1123":      time.RFC1123,
		"rfc1123z":     time.RFC1123Z,
	}
)

type Format struct {
	rc_recipe.RemarkTransient
	Time      int64
	Precision mo_string.SelectString
	Format    mo_string.SelectString
}

func (z *Format) Preset() {
	z.Precision.SetOptions(
		PrecisionSecond,
		PrecisionSecond,
		PrecisionMilliSecond,
		PrecisionNanoSecond,
	)
	formatKeys := make([]string, 0)
	for k := range supportedFormats {
		formatKeys = append(formatKeys, k)
	}
	sort.Strings(formatKeys)
	z.Format.SetOptions(
		"iso8601",
		formatKeys...,
	)
}

func (z *Format) Exec(c app_control.Control) error {
	l := c.Log()
	var outFormat string
	var found bool
	if outFormat, found = supportedFormats[z.Format.Value()]; !found {
		l.Error("Undefined format", esl.String("format", z.Format.Value()))
		return errors.New("undefined format")
	}

	switch z.Precision.Value() {
	case PrecisionSecond:
		ui_out.TextOut(c, time.Unix(z.Time, 0).Format(outFormat))
	case PrecisionMilliSecond:
		ui_out.TextOut(c, time.Unix(z.Time/1000, z.Time%1000*1000).Format(outFormat))
	case PrecisionNanoSecond:
		ui_out.TextOut(c, time.Unix(z.Time/1_000_000_000, z.Time%1_000_000_000).Format(outFormat))
	default:
		l.Error("Undefined precision", esl.String("precision", z.Precision.Value()))
		return errors.New("undefined precision")
	}
	return nil
}

func (z *Format) Test(c app_control.Control) error {
	{
		now := time.Now()
		err := rc_exec.Exec(c, &Format{}, func(r rc_recipe.Recipe) {
			m := r.(*Format)
			m.Time = now.Unix()
		})
		if err != nil {
			return err
		}
		if x := ui_out.CapturedText(c); x != now.Format(time.RFC3339) {
			return errors.New("invalid format: " + x)
		}
	}
	return nil
}
