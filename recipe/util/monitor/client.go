package monitor

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/quality/infra/qt_errors"
)

type Client struct {
	Peer            dbx_conn.ConnScopedIndividual
	Name            string
	MonitorInterval mo_int.RangeInt
	SyncInterval    mo_int.RangeInt
}

func (z *Client) Preset() {
	z.MonitorInterval.SetRange(1, 86400, 10)
	z.SyncInterval.SetRange(10, 86400, 300)
}
func (z *Client) sendError(c app_control.Control, eventType string, err error) {
	l := c.Log()
	l.Warn("Unable to retrieve event data", esl.String("type", eventType), esl.Error(err))
}

func (z *Client) sendEvent(c app_control.Control, eventType string, data interface{}) {
	ev := Event{
		Type: eventType,
		Data: data,
	}
	c.Log().Info("event", esl.Any("event", ev))
}

func (z *Client) eventCpuInfo(c app_control.Control) {
	cpuInfo, err := cpu.Info()
	if err != nil {
		z.sendError(c, EventCpuInfo, err)
		return
	}
	z.sendEvent(c, EventCpuInfo, cpuInfo)
}

func (z *Client) eventCpuTime(c app_control.Control) {
	stat, err := cpu.Times(true)
	if err != nil {
		z.sendError(c, EventCpuTime, err)
		return
	}
	z.sendEvent(c, EventCpuTime, stat)
}

func (z *Client) eventHostInfo(c app_control.Control) {
	hostInfo, err := host.Info()
	if err != nil {
		z.sendError(c, EventHostInfo, err)
		return
	}
	z.sendEvent(c, EventHostInfo, hostInfo)
}

func (z *Client) eventDiskPartition(c app_control.Control) {
	info, err := disk.Partitions(true)
	if err != nil {
		z.sendError(c, EventDiskPartition, err)
		return
	}
	z.sendEvent(c, EventDiskPartition, info)
}

func (z *Client) eventDiskUsage(c app_control.Control) {
	partitions, err := disk.Partitions(true)
	if err != nil {
		z.sendError(c, EventDiskUsage, err)
		return
	}
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			z.sendError(c, EventDiskUsage, err)
			continue
		}
		z.sendEvent(c, EventDiskUsage, usage)
	}
}

func (z *Client) eventLoadAverage(c app_control.Control) {
	la, err := load.Avg()
	if err != nil {
		z.sendError(c, EventLoadAverage, err)
		return
	}
	z.sendEvent(c, EventLoadAverage, la)
}

func (z *Client) eventMemoryStat(c app_control.Control) {
	l := c.Log()
	vm, err := mem.VirtualMemory()
	if err != nil {
		l.Warn("Unable to retrieve memory stat", esl.Error(err))
		return
	}
	z.sendEvent(c, EventMemoryStat, vm)
}

func (z *Client) Exec(c app_control.Control) error {
	z.eventCpuInfo(c)
	z.eventHostInfo(c)
	z.eventDiskPartition(c)

	// Periodical events
	z.eventLoadAverage(c)
	z.eventCpuTime(c)
	z.eventMemoryStat(c)
	z.eventDiskUsage(c)
	return nil
}

func (z *Client) Test(c app_control.Control) error {
	return qt_errors.ErrorMock
}
