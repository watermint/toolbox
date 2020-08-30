package rp_writer_impl

import (
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/report/rp_writer"
)

func New(name string, ctl app_control.Control) rp_writer.Writer {
	caw := NewCascade(name, ctl)
	if ctl.Feature().IsTest() {
		return caw
	}
	scw := NewSmallCache(name, caw)
	return scw
}
