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
	return &TimeImpl{time: time.Time{}}
}

func New(t string) (tm Time, err error) {
	ts, valid := ut_time.ParseTimestamp(t)
	if !valid {
		return nil, InvalidTimeFormat
	}
	return &TimeImpl{time: ts}, nil
}

type TimeImpl struct {
	time time.Time
}

func (z *TimeImpl) IsZero() bool {
	return z.time.IsZero()
}

func (z *TimeImpl) Time() time.Time {
	return z.time
}

func (z *TimeImpl) Iso8601() string {
	return api_util.RebaseAsString(z.time)
}

func (z *TimeImpl) String() string {
	return z.Iso8601()
}

func (z *TimeImpl) UpdateTime(dateTime string) error {
	ts, valid := ut_time.ParseTimestamp(dateTime)
	if !valid {
		return InvalidTimeFormat
	}
	z.time = ts
	return nil
}
