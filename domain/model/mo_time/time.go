package mo_time

import (
	"errors"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"time"
)

type Time interface {
	Iso8601() string
	String() string
	Time() time.Time
	IsZero() bool
}

var (
	InvalidTimeFormat = errors.New("invalid time format")
)

func Zero() (tm Time) {
	return &timeImpl{time: time.Time{}}
}

func New(t string) (tm Time, err error) {
	ts, valid := ut_time.ParseTimestamp(t)
	if !valid {
		return nil, InvalidTimeFormat
	}
	return &timeImpl{time: ts}, nil
}

type timeImpl struct {
	time time.Time
}

func (z *timeImpl) IsZero() bool {
	return z.time.IsZero()
}

func (z *timeImpl) Time() time.Time {
	return z.time
}

func (z *timeImpl) Iso8601() string {
	return api_util.RebaseAsString(z.time)
}

func (z *timeImpl) String() string {
	return z.Iso8601()
}
