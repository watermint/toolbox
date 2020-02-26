package mo_time

import (
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_essential"
	"github.com/watermint/toolbox/infra/api/api_util"
	"github.com/watermint/toolbox/infra/util/ut_time"
	"time"
)

type Time interface {
	// Returns Time in ISO8601 (yyyy-MM-ddThh:mm:ssZ) format as UTC.
	// Returns an empty string if an impl. has TimeOptional, and the instance marked as unset.
	Iso8601() string

	// Same as Iso8601
	String() string

	// Returns time instance
	Time() time.Time

	// True when the time is zero
	IsZero() bool
}

type TimeOptional interface {
	Time
	mo_essential.Optional
}

var (
	InvalidTimeFormat = errors.New("invalid time format")
)

func Zero() (tm Time) {
	return &TimeImpl{time: time.Time{}, isSet: true}
}

func New(t string) (tm Time, err error) {
	ts, valid := ut_time.ParseTimestamp(t)
	if !valid {
		return nil, InvalidTimeFormat
	}
	return &TimeImpl{time: ts, isSet: true}, nil
}

type TimeImpl struct {
	time  time.Time
	isSet bool
}

func (z *TimeImpl) Unset() {
	z.isSet = false
}

func (z *TimeImpl) Ok() bool {
	return !z.IsZero() && z.isSet
}

func (z *TimeImpl) IsZero() bool {
	return z.time.IsZero()
}

func (z *TimeImpl) Time() time.Time {
	return z.time
}

func (z *TimeImpl) Iso8601() string {
	if z.Ok() {
		return api_util.RebaseAsString(z.time)
	} else {
		return ""
	}
}

func (z *TimeImpl) String() string {
	if z.Ok() {
		return z.Iso8601()
	} else {
		return ""
	}
}

func (z *TimeImpl) UpdateTime(dateTime string) error {
	ts, valid := ut_time.ParseTimestamp(dateTime)
	if !valid {
		return InvalidTimeFormat
	}
	z.time = ts
	z.isSet = true
	return nil
}
