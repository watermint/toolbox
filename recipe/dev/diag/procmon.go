package diag

import (
	"archive/zip"
	"encoding/json"
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_path"
	"github.com/watermint/toolbox/domain/model/mo_time"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_conn"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/ingredient/file"
	"go.uber.org/zap"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"os/exec"
	"os/user"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
	"time"
)

const (
	procmonDownloadUrl = "https://download.sysinternals.com/files/ProcessMonitor.zip"
	procmonExe32       = "ProcMon.exe"
	procmonExe64       = "ProcMon64.exe"
	procmonLogPrefix   = "monitor"
	procmonLogSummary  = "info.json"
)

type Procmon struct {
	ProcmonUrl     string
	RepositoryPath mo_path.FileSystemPath
	DropboxPath    mo_path.DropboxPath
	Peer           rc_conn.ConnUserFile
	RunUntil       mo_time.Time
	RetainLogs     int
	Seconds        int
	Upload         *file.Upload
}

func (z *Procmon) downloadProcmon(c app_control.Control) error {
	l := c.Log()

	err := os.MkdirAll(z.RepositoryPath.Path(), 0755)
	if err != nil {
		l.Debug("Unable to create repository path", zap.Error(err))
		return err
	}
	procmonZip := filepath.Join(z.RepositoryPath.Path(), "procmon.zip")

	// Download
	{
		l.Info("Try download", zap.String("url", z.ProcmonUrl))
		resp, err := http.Get(z.ProcmonUrl)
		if err != nil {
			l.Debug("Unable to create download request")
			return err
		}
		defer resp.Body.Close()

		out, err := os.Create(procmonZip)
		if err != nil {
			l.Debug("Unable to create download file")
			return err
		}
		defer out.Close()

		_, err = io.Copy(out, resp.Body)
		if err != nil {
			l.Debug("Unable to copy from response", zap.Error(err))
			return err
		}
	}

	// Extract
	{
		l.Info("Extract downloaded zip", zap.String("zip", procmonZip))
		r, err := zip.OpenReader(procmonZip)
		if err != nil {
			l.Debug("Unable to open zip file", zap.Error(err))
			return err
		}

		for _, f := range r.File {
			compressed, err := f.Open()
			if err != nil {
				l.Debug("Unable to open compressed file", zap.Error(err), zap.String("name", f.Name))
				return err
			}

			extractPath := filepath.Join(z.RepositoryPath.Path(), filepath.Base(f.Name))
			l.Debug("Extract file", zap.String("extractPath", extractPath))
			extracted, err := os.Create(extractPath)
			if err != nil {
				l.Debug("Unable to create extract file", zap.Error(err))
				return err
			}

			_, err = io.Copy(extracted, compressed)
			if err != nil {
				l.Debug("Unable to copy from zip", zap.Error(err))
				return err
			}
			extracted.Close()
			compressed.Close()
		}
	}

	return err
}

func (z *Procmon) ensureProcmon(c app_control.Control) (exePath string, err error) {
	l := c.Log()
	if runtime.GOARCH == "amd64" {
		exePath = filepath.Join(z.RepositoryPath.Path(), procmonExe64)
	} else {
		exePath = filepath.Join(z.RepositoryPath.Path(), procmonExe32)
	}

	info, err := os.Lstat(exePath)
	if err != nil {
		l.Debug("Unable to find exe", zap.Error(err), zap.String("exe", exePath))
		err = z.downloadProcmon(c)
		if err != nil {
			return "", err
		}
	}
	l.Debug("Exe info", zap.Any("info", info))
	return exePath, err
}

func (z *Procmon) runProcmon(c app_control.Control, exePath string) (cmd *exec.Cmd, logPath string, err error) {
	l := c.Log()

	logName := c.Workspace().JobId()
	logPath = filepath.Join(z.RepositoryPath.Path(), "logs", logName)
	l.Debug("Creating log path", zap.String("path", logPath))
	err = os.MkdirAll(logPath, 0755)
	if err != nil {
		l.Debug("Unable to create log path", zap.Error(err))
		return nil, "", err
	}

	{
		hostname, _ := os.Hostname()
		usr, _ := user.Current()

		info := struct {
			TimeLocal string `json:"time_local"`
			TimeUTC   string `json:"time_utc"`
			Hostname  string `json:"hostname"`
			Username  string `json:"username"`
			UserHome  string `json:"user_home"`
			UserUID   string `json:"user_uid"`
			UserGID   string `json:"user_gid"`
		}{
			TimeLocal: time.Now().Local().Format(time.RFC3339),
			TimeUTC:   time.Now().UTC().Format(time.RFC3339),
			Hostname:  hostname,
			Username:  usr.Name,
			UserHome:  usr.HomeDir,
			UserUID:   usr.Uid,
			UserGID:   usr.Gid,
		}
		content, err := json.Marshal(&info)
		if err != nil {
			l.Error("Unable to create info file", zap.Error(err))
		}

		err = ioutil.WriteFile(filepath.Join(logPath, procmonLogSummary), content, 0644)
		if err != nil {
			l.Error("Unable to write info file", zap.Error(err))
		}
	}

	if !app.IsWindows() {
		l.Warn("Skip run procmon (Reason; not on Windows)")
		return nil, logPath, nil
	}

	cmd = exec.Command(exePath,
		"/AcceptEula",
		"/Quiet",
		"/Minimized",
		"/BackingFile",
		filepath.Join(logPath, procmonLogPrefix),
	)
	l.Info("Run Process monitor", zap.String("exe", exePath), zap.Strings("args", cmd.Args))

	err = cmd.Start()
	if err != nil {
		l.Debug("Unable to start program", zap.Error(err), zap.Any("cmd", cmd))
		return nil, logPath, err
	}

	return cmd, logPath, nil
}

func (z *Procmon) watchProcmon(c app_control.Control, exePath string, cmd *exec.Cmd, logPath string) error {
	l := c.Log().With(zap.String("logPath", logPath))

	if cmd == nil || !app.IsWindows() {
		l.Info("skip watching")
		return nil
	}

	go func() {
		for {
			time.Sleep(1 * time.Second)
			l.Debug("Process", zap.Any("status", cmd.ProcessState))

			entries, err := ioutil.ReadDir(logPath)
			if err != nil {
				l.Debug("Unable to list dir", zap.Error(err))
				continue
			}
			if z.RetainLogs == 0 {
				continue
			}

			logEntries := make([]os.FileInfo, 0)
			modTimes := make([]string, 0)
			cmpTimeFormat := "20060102-150405.000"

			for _, f := range entries {
				if f.IsDir() {
					continue
				}
				if strings.HasPrefix(strings.ToLower(f.Name()), procmonLogPrefix) {
					logEntries = append(logEntries, f)
					mt := f.ModTime().Format(cmpTimeFormat)
					modTimes = append(modTimes, mt)
					l.Debug("Log file found", zap.Any("entry", f), zap.String("modTime", mt))
				}
			}

			if len(modTimes) <= z.RetainLogs {
				l.Debug("Log files is less than threshold")
				continue
			}

			sort.Strings(modTimes)
			thresholdIndex := len(modTimes) - z.RetainLogs
			thresholdTime := modTimes[thresholdIndex]

			for _, f := range logEntries {
				et := f.ModTime().Format(cmpTimeFormat)
				if strings.Compare(et, thresholdTime) < 0 {
					l.Debug("Remove log", zap.Any("entry", f))
					lf := filepath.Join(logPath, f.Name())
					err = os.Remove(lf)
					l.Debug("Removed", zap.Error(err), zap.String("logFile", lf))
				} else {
					l.Debug("Retain file", zap.Any("entry", f))
				}
			}
		}
	}()

	l.Info("Waiting for duration", zap.Int("seconds", z.Seconds))
	time.Sleep(time.Duration(z.Seconds) * time.Second)

	return nil
}

func (z *Procmon) terminateProcmon(c app_control.Control, exePath string, cmd *exec.Cmd) error {
	l := c.Log()

	if !app.IsWindows() {
		l.Warn("Skip run procmon (Reason; not on Windows)")
		return nil
	}

	l.Info("Trying to terminate procmon")
	termCmd := exec.Command(exePath,
		"/Terminate",
	)
	err := termCmd.Start()
	if err != nil {
		l.Debug("Unable to invoke procmon", zap.Error(err), zap.Any("cmd", cmd))
		l.Debug("Trying to terminate thru cmd")
		err2 := cmd.Process.Kill()
		l.Debug("Kill sent", zap.Error(err2))
		return err
	}
	if err := termCmd.Wait(); err != nil {
		l.Debug("Terminate wait returned an error", zap.Error(err))
		return nil
	}

	l.Info("Waiting for termination", zap.Int("seconds", 60))
	time.Sleep(60 * time.Second)

	return nil
}

func (z *Procmon) uploadProcmonLogs(c app_control.Control) error {
	logPath := filepath.Join(z.RepositoryPath.Path(), "logs")
	l := c.Log().With(zap.String("logPath", logPath))
	l.Debug("Start uploading logs")

	err := rc_exec.Exec(c, z.Upload, func(r rc_recipe.Recipe) {
		ru := r.(*file.Upload)
		ru.EstimateOnly = false
		ru.LocalPath = mo_path.NewFileSystemPath(logPath)
		ru.DropboxPath = z.DropboxPath
		ru.Overwrite = true
		ru.CreateFolder = true
		ru.Context = z.Peer.Context()
	})
	if err != nil {
		l.Debug("Unable to upload logs", zap.Error(err))
		return err
	}
	return nil
}

func (z *Procmon) cleanupProcmonLogs(c app_control.Control) error {
	logPath := filepath.Join(z.RepositoryPath.Path(), "logs")
	l := c.Log().With(zap.String("logPath", logPath))
	l.Debug("Start clean up logs")

	for i := 0; i < 10; i++ {
		err := os.RemoveAll(logPath)
		if err != nil {
			l.Debug("Unable to clean up logs", zap.Error(err))
			time.Sleep(10 * time.Second)
			continue
		}
		return nil
	}
	return nil
}

func (z *Procmon) Exec(c app_control.Control) error {
	l := c.Log()

	if z.Seconds < 10 {
		return errors.New("seconds must grater than 10 sec")
	}
	if z.RunUntil.Time().Before(time.Now()) {
		l.Info("Skip run")
		return nil
	}

	exe, err := z.ensureProcmon(c)
	if err != nil {
		return err
	}
	l.Debug("Procmon exe", zap.String("exe", exe))

	cmd, logPath, err := z.runProcmon(c, exe)
	if err != nil {
		return err
	}

	if err = z.watchProcmon(c, exe, cmd, logPath); err != nil {
		return err
	}
	if err = z.terminateProcmon(c, exe, cmd); err != nil {
		return err
	}
	if err = z.uploadProcmonLogs(c); err != nil {
		return err
	}
	if err = z.cleanupProcmonLogs(c); err != nil {
		return err
	}
	return nil
}

func (z *Procmon) Test(c app_control.Control) error {
	tmpDir, err := ioutil.TempDir("", "procmon")
	if err != nil {
		return err
	}
	defer func() {
		os.RemoveAll(tmpDir)
	}()

	return rc_exec.Exec(c, &Procmon{}, func(r rc_recipe.Recipe) {
		m := r.(*Procmon)
		m.ProcmonUrl = procmonDownloadUrl
		m.Seconds = 30
		m.RetainLogs = 4
		m.RepositoryPath = mo_path.NewFileSystemPath(tmpDir)
	})
}

func (z *Procmon) Preset() {
	ru, err := mo_time.New(time.Now().Add(7 * 24 * time.Hour).Format("2006-01-02"))
	if err != nil {
		panic(err)
	}
	z.ProcmonUrl = procmonDownloadUrl
	z.Seconds = 1800
	z.RunUntil = ru
	z.RetainLogs = 4
}
