package mo_time

import (
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"time"
)

type Time interface {
	Iso8601() string
	String() string
	Time() time.Time
	Normalize() (t Time, valid bool)
}

type TimeImpl struct {
	DateTime string
	time     time.Time
	timeSet  bool
}

func (z *TimeImpl) Time() time.Time {
	if z.timeSet {
		return z.time
	}
	t, v := z.Normalize()
	if v {
		return t.Time()
	}
	return z.time
}

func (z *TimeImpl) Normalize() (t Time, valid bool) {
	tm, valid := ut_time.ParseTimestamp(z.DateTime)
	return &TimeImpl{
		DateTime: z.DateTime,
		time:     tm,
		timeSet:  true,
	}, valid
}

func (z *TimeImpl) Iso8601() string {
	if z.timeSet {
		return api_util.RebaseAsString(z.time)
	}
	t, v := z.Normalize()
	if v {
		return t.Iso8601()
	}
	return z.DateTime
}

func (z *TimeImpl) String() string {
	return z.Iso8601()
}
