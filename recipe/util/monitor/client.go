package monitor

import (
	"encoding/json"
	"fmt"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/load"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/net"
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
	"github.com/watermint/toolbox/infra/control/app_definitions"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_file"
	"github.com/watermint/toolbox/quality/recipe/qtr_endtoend"
	"os"
	"os/user"
	"path/filepath"
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
	Display         bool

	sentErrors       map[string]int64 // eventType -> num error sent
	sentCount        map[string]int64 // eventType -> num event sent
	currentJournal   *os.File
	currentPath      string
	currentStart     time.Time
	currentDeadline  time.Time
	rotateInProgress bool
}

func (z *Client) Preset() {
	z.MonitorInterval.SetRange(1, 86400, 10)
	z.SyncInterval.SetRange(10, 86400, 3600)
	z.sentErrors = make(map[string]int64)
	z.sentCount = make(map[string]int64)
	z.Peer.SetScopes(
		dbx_auth.ScopeFilesContentWrite,
	)
}
func (z *Client) sendError(c app_control.Control, eventType string, err error) {
	if n, ok := z.sentErrors[eventType]; ok {
		z.sentErrors[eventType] = n + 1
		return
	}
	l := c.Log()
	l.Warn("Unable to retrieve event data", esl.String("type", eventType), esl.Error(err))
	z.sentErrors[eventType] = 1
}

func (z *Client) openJournal(c app_control.Control) error {
	l := c.Log()
	if z.currentJournal != nil {
		return nil
	}
	name := monitorFilePrefix + es_filepath.Escape(z.Name) + "-" + fmt.Sprintf("%08x", time.Now().Unix()) + ".log"
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

	sv := sv_file_content.NewUpload(z.Peer.Client())
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
	if z.Display {
		c.Log().Info("event", esl.Any("data", ev))
	}

	if z.currentJournal != nil {
		_, err0 := z.currentJournal.Write(evs)
		_, err1 := z.currentJournal.Write([]byte("\n"))

		if z.rotateInProgress {
			return
		}

		if err0 == nil && err1 == nil && z.currentDeadline.Before(time.Now()) {
			z.rotateInProgress = true
			if err := z.syncJournal(c); err != nil {
				z.sendError(c, eventType, err)
			}
			if err := z.openJournal(c); err != nil {
				z.sendError(c, eventType, err)
			}
			z.headEvents(c)
			z.rotateInProgress = false
		}
	}

	if ns, ok := z.sentCount[eventType]; ok {
		z.sentCount[eventType] = ns + 1
	} else {
		z.sentCount[eventType] = 1
	}
}

func (z *Client) shouldAbort(eventType string) bool {
	if ne, oke := z.sentErrors[eventType]; oke {
		if ns, okc := z.sentCount[eventType]; okc {
			// Abort when three more errors without success
			if 2 < ne && ns < 1 {
				return true
			}
		}
	}
	return false
}

func (z *Client) handleEvent(c app_control.Control, eventType string, f func() (data interface{}, err error)) {
	if z.shouldAbort(eventType) {
		return
	}

	data, err := f()
	if err != nil {
		z.sendError(c, eventType, err)
	} else {
		z.sendEvent(c, eventType, data)
	}
}

func (z *Client) headEventCpuInfo(c app_control.Control) {
	z.handleEvent(c, EventCpuInfo, func() (data interface{}, err error) {
		return cpu.Info()
	})
}

func (z *Client) eventCpuTime(c app_control.Control) {
	z.handleEvent(c, EventCpuTime, func() (data interface{}, err error) {
		return cpu.Times(true)
	})
}

func (z *Client) eventCpuPercent(c app_control.Control) {
	z.handleEvent(c, EventCpuPercent, func() (data interface{}, err error) {
		return cpu.Percent(0, true)
	})
}

func (z *Client) headEventHostInfo(c app_control.Control) {
	z.handleEvent(c, EventHostInfo, func() (data interface{}, err error) {
		return host.Info()
	})
}

func (z *Client) headEventDiskPartition(c app_control.Control) {
	z.handleEvent(c, EventDiskPartition, func() (data interface{}, err error) {
		return disk.Partitions(true)
	})
}

func (z *Client) eventDiskUsage(c app_control.Control) {
	if z.shouldAbort(EventDiskUsage) {
		return
	}

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
	z.handleEvent(c, EventLoadAverage, func() (data interface{}, err error) {
		return load.Avg()
	})
}

func (z *Client) eventMemoryStat(c app_control.Control) {
	z.handleEvent(c, EventMemoryStat, func() (data interface{}, err error) {
		return mem.VirtualMemory()
	})
}

func (z *Client) eventNetIO(c app_control.Control) {
	z.handleEvent(c, EventNetIO, func() (data interface{}, err error) {
		return net.IOCounters(true)
	})
}

func (z *Client) eventNetProtocol(c app_control.Control) {
	z.handleEvent(c, EventNetProtocol, func() (data interface{}, err error) {
		return net.ProtoCounters([]string{})
	})
}

func (z *Client) headEventMonitorInfo(c app_control.Control) {
	var userUserName, userDisplayName, userUid string
	if usr, err := user.Current(); err == nil {
		userUserName = usr.Username
		userDisplayName = usr.Name
		userUid = usr.Uid
	}

	z.sendEvent(c, EventMonitorInfo, struct {
		AppVersion      string `json:"app_version"`
		MonitorName     string `json:"monitor_name"`
		IntervalMonitor int    `json:"interval_monitor"`
		IntervalSync    int    `json:"interval_sync"`
		UserDisplayName string `json:"user_display_name"`
		UserUid         string `json:"user_uid"`
		UserName        string `json:"user_name"`
	}{
		AppVersion:      app_definitions.BuildId,
		MonitorName:     z.Name,
		IntervalMonitor: z.MonitorInterval.Value(),
		IntervalSync:    z.SyncInterval.Value(),
		UserDisplayName: userDisplayName,
		UserUid:         userUid,
		UserName:        userUserName,
	})
}

func (z *Client) headEvents(c app_control.Control) {
	z.headEventMonitorInfo(c)
	z.headEventCpuInfo(c)
	z.headEventHostInfo(c)
	z.headEventDiskPartition(c)
}

func (z *Client) Exec(c app_control.Control) error {
	l := c.Log()
	if err := os.MkdirAll(z.DataPath.Path(), 0755); err != nil {
		return err
	}
	if err := z.openJournal(c); err != nil {
		return err
	}

	// Head events
	z.headEvents(c)

	// Periodical events
	for {
		z.eventCpuPercent(c)
		z.eventCpuTime(c)
		z.eventDiskUsage(c)
		z.eventLoadAverage(c)
		z.eventMemoryStat(c)
		z.eventNetIO(c)
		z.eventNetProtocol(c)

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
