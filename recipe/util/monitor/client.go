package monitor

import (
	"encoding/json"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_auth"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	mo_path2 "github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/file/es_gzip"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_int"
	"github.com/watermint/toolbox/essentials/model/mo_path"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

const (
	monitorFilePrefix = "tbx-monitor-"
)

type Client struct {
	DataPath        mo_path.FileSystemPath
	SyncPath        mo_path2.DropboxPath
	Peer            dbx_conn.ConnScopedIndividual
	Name            string
	MonitorInterval mo_int.RangeInt
	SyncInterval    mo_int.RangeInt
	MonitorEnd      mo_time.TimeOptional

	sentErrors      map[string]bool
	currentJournal  *os.File
	currentPath     string
	currentStart    time.Time
	currentDeadline time.Time
}

func (z *Client) Preset() {
	z.MonitorInterval.SetRange(1, 86400, 10)
	z.SyncInterval.SetRange(10, 86400, 3600)
	z.sentErrors = make(map[string]bool)
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
}
func (z *Client) sendError(c app_control.Control, eventType string, err error) {
	if _, ok := z.sentErrors[eventType]; ok {
		return
	}
	l := c.Log()
	l.Warn("Unable to retrieve event data", esl.String("type", eventType), esl.Error(err))
	z.sentErrors[eventType] = true
}

func (z *Client) openJournal(c app_control.Control) error {
	l := c.Log()
	if z.currentJournal != nil {
		return nil
	}
	name := monitorFilePrefix + es_filepath.Escape(z.Name) + "-" + strconv.FormatInt(time.Now().Unix(), 16) + ".log"
	path := filepath.Join(z.DataPath.Path(), name)
	l = l.With(esl.String("path", path))
	f, err := os.Create(path)
	if err != nil {
		l.Debug("Unable to create the log file", esl.Error(err))
		return err
	}
	z.currentJournal = f
	z.currentPath = path
	z.currentStart = time.Now()
	z.currentDeadline = z.currentStart.Add(time.Duration(z.SyncInterval.Value()) * time.Second)
	l.Debug("Journal created", esl.Time("deadline", z.currentDeadline))
	return nil
}

func (z *Client) syncJournal(c app_control.Control) error {
	l := c.Log()
	if err := z.currentJournal.Close(); err != nil {
		l.Debug("Unable to close", esl.Error(err))
	}
	if _, err := es_gzip.Compress(z.currentPath); err != nil {
		l.Debug("Unable to compress", esl.Error(err))
		return err
	}
	z.currentJournal = nil
	z.currentPath = ""

	files, err := os.ReadDir(z.DataPath.Path())
	if err != nil {
		l.Debug("Unable to read directory entry")
		return err
	}

	sv := sv_file_content.NewUpload(z.Peer.Context())
	basePath := z.SyncPath.ChildPath(es_filepath.Escape(z.Name), z.currentStart.Format("2006-01"), z.currentStart.Format("2006-01-02"))
	for _, f := range files {
		if !strings.HasPrefix(f.Name(), monitorFilePrefix) {
			continue
		}
		fp := filepath.Join(z.DataPath.Path(), f.Name())
		l.Info("Syncing journal file", esl.String("name", f.Name()))
		entry, err := sv.Add(basePath, fp)
		if err != nil {
			l.Debug("Unable to upload", esl.Error(err))
			return err
		}
		l.Debug("Upload completed", esl.Any("entry", entry))
		_ = os.Remove(fp)
	}
	return nil
}

func (z *Client) sendEvent(c app_control.Control, eventType string, data interface{}) {
	ev := Event{
		Time: time.Now().Format(time.RFC3339),
		Type: eventType,
		Data: data,
	}
	evs, err := json.Marshal(&ev)
	if err != nil {
		z.sendError(c, eventType, err)
		return
	}

	if z.currentJournal != nil {
		_, err0 := z.currentJournal.Write(evs)
		_, err1 := z.currentJournal.Write([]byte("\n"))
		if err0 == nil && err1 == nil && z.currentDeadline.Before(time.Now()) {
			if err := z.syncJournal(c); err != nil {
				z.sendError(c, eventType, err)
			}
			if err := z.openJournal(c); err != nil {
				z.sendError(c, eventType, err)
			}
		}
	}
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
	l := c.Log()
	if err := os.MkdirAll(z.DataPath.Path(), 0755); err != nil {
		return err
	}
	if err := z.openJournal(c); err != nil {
		return err
	}

	z.eventCpuInfo(c)
	z.eventHostInfo(c)
	z.eventDiskPartition(c)

	// Periodical events
	for {
		z.eventLoadAverage(c)
		z.eventCpuTime(c)
		z.eventMemoryStat(c)
		z.eventDiskUsage(c)

		if !z.MonitorEnd.IsZero() && z.MonitorEnd.Time().Before(time.Now()) {
			return z.syncJournal(c)
		}
		time.Sleep(time.Duration(z.MonitorInterval.Value()) * time.Second)
		l.Info("Monitor", esl.Time("t", time.Now()))
	}
}

func (z *Client) Test(c app_control.Control) error {
	f, err := qt_file.MakeTestFolder("monitor", false)
	if err != nil {
		return err
	}
	defer func() {
		_ = os.RemoveAll(f)
	}()

	return rc_exec.ExecMock(c, &Client{}, func(r rc_recipe.Recipe) {
		m := r.(*Client)
		m.Name = "mango"
		m.DataPath = mo_path.NewFileSystemPath(f)
		m.SyncPath = qtr_endtoend.NewTestDropboxFolderPath("monitor")
		m.MonitorEnd = mo_time.NewOptional(time.Now())
	})
}
