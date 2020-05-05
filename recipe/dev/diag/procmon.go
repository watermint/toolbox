package diag

import (
	"archive/zip"
	"encoding/json"
	"github.com/watermint/toolbox/domain/common/model/mo_int"
	mo_path2 "github.com/watermint/toolbox/domain/common/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/api/dbx_conn"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_path"
	"github.com/watermint/toolbox/domain/dropbox/model/mo_time"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_file_content"
	"github.com/watermint/toolbox/domain/dropbox/service/sv_profile"
	"github.com/watermint/toolbox/essentials/file/es_zip"
	"github.com/watermint/toolbox/essentials/http/es_download"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/essentials/log/es_process"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/quality/infra/qt_endtoend"
	"github.com/watermint/toolbox/quality/infra/qt_recipe"
	"io"
	"io/ioutil"
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
	procmonLogFinish   = "info_finish.json"
)

type Procmon struct {
	rc_recipe.RemarkSecret
	ProcmonUrl     string
	RepositoryPath mo_path2.FileSystemPath
	DropboxPath    mo_path.DropboxPath
	Peer           dbx_conn.ConnUserFile
	RunUntil       mo_time.TimeOptional
	RetainLogs     mo_int.RangeInt
	Seconds        mo_int.RangeInt
}

func (z *Procmon) downloadProcmon(c app_control.Control) error {
	l := c.Log()

	err := os.MkdirAll(z.RepositoryPath.Path(), 0755)
	if err != nil {
		l.Debug("Unable to create repository path", es_log.Error(err))
		return err
	}
	procmonZip := filepath.Join(z.RepositoryPath.Path(), "procmon.zip")

	// Download
	if err := es_download.Download(l, z.ProcmonUrl, procmonZip); err != nil {
		return err
	}

	// Extract
	{
		l.Info("Extract downloaded zip", es_log.String("zip", procmonZip))
		r, err := zip.OpenReader(procmonZip)
		if err != nil {
			l.Debug("Unable to open zip file", es_log.Error(err))
			return err
		}

		for _, f := range r.File {
			compressed, err := f.Open()
			if err != nil {
				l.Debug("Unable to open compressed file", es_log.Error(err), es_log.String("name", f.Name))
				return err
			}

			extractPath := filepath.Join(z.RepositoryPath.Path(), filepath.Base(f.Name))
			l.Debug("Extract file", es_log.String("extractPath", extractPath))
			extracted, err := os.Create(extractPath)
			if err != nil {
				l.Debug("Unable to create extract file", es_log.Error(err))
				return err
			}

			_, err = io.Copy(extracted, compressed)
			if err != nil {
				l.Debug("Unable to copy from zip", es_log.Error(err))
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
		l.Debug("Unable to find exe", es_log.Error(err), es_log.String("exe", exePath))
		err = z.downloadProcmon(c)
		if err != nil {
			return "", err
		}
	}
	l.Debug("Exe info", es_log.Any("info", info))
	return exePath, err
}

func (z *Procmon) runProcmon(c app_control.Control, exePath string) (cmd *exec.Cmd, logPath string, err error) {
	l := c.Log()

	logName := c.Workspace().JobId()
	logPath = filepath.Join(z.RepositoryPath.Path(), "logs", logName)
	l.Debug("Creating log path", es_log.String("path", logPath))
	err = os.MkdirAll(logPath, 0755)
	if err != nil {
		l.Debug("Unable to create log path", es_log.Error(err))
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
			l.Error("Unable to create info file", es_log.Error(err))
		}

		err = ioutil.WriteFile(filepath.Join(logPath, procmonLogSummary), content, 0644)
		if err != nil {
			l.Error("Unable to write info file", es_log.Error(err))
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
	l.Info("Run Process monitor", es_log.String("exe", exePath), es_log.Strings("args", cmd.Args))

	pl := es_process.NewLogger(cmd, c)
	pl.Start()
	defer pl.Close()

	err = cmd.Start()
	if err != nil {
		l.Debug("Unable to start program", es_log.Error(err), es_log.Any("cmd", cmd))
		return nil, logPath, err
	}

	return cmd, logPath, nil
}

func (z *Procmon) watchProcmon(c app_control.Control, exePath string, cmd *exec.Cmd, logPath string) error {
	l := c.Log().With(es_log.String("logPath", logPath))

	if cmd == nil || !app.IsWindows() {
		l.Info("skip watching")
		return nil
	}

	go func() {
		for {
			time.Sleep(1 * 1000 * time.Millisecond)
			l.Debug("Process", es_log.Any("state", cmd.ProcessState), es_log.Any("process", cmd.Process), es_log.Any("sysAttr", cmd.SysProcAttr))

			entries, err := ioutil.ReadDir(logPath)
			if err != nil {
				l.Debug("Unable to list dir", es_log.Error(err))
				continue
			}
			if z.RetainLogs.Value() == 0 {
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
					l.Debug("Log file found", es_log.Any("entry", f), es_log.String("modTime", mt))
				}
			}

			if len(modTimes) <= z.RetainLogs.Value() {
				l.Debug("Log files is less than threshold")
				continue
			}

			sort.Strings(modTimes)
			thresholdIndex := len(modTimes) - z.RetainLogs.Value()
			thresholdTime := modTimes[thresholdIndex]

			for _, f := range logEntries {
				et := f.ModTime().Format(cmpTimeFormat)
				if strings.Compare(et, thresholdTime) < 0 {
					l.Debug("Remove log", es_log.Any("entry", f))
					lf := filepath.Join(logPath, f.Name())
					err = os.Remove(lf)
					l.Debug("Removed", es_log.Error(err), es_log.String("logFile", lf))
				} else {
					l.Debug("Retain file", es_log.Any("entry", f))
				}
			}
		}
	}()

	l.Info("Waiting for duration", es_log.Int("seconds", z.Seconds.Value()))
	time.Sleep(time.Duration(z.Seconds.Value()) * 1000 * time.Millisecond)

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
	pl := es_process.NewLogger(cmd, c)
	pl.Start()
	defer pl.Close()

	err := termCmd.Start()
	if err != nil {
		l.Debug("Unable to invoke procmon", es_log.Error(err), es_log.Any("cmd", cmd))
		l.Debug("Trying to terminate thru cmd")
		err2 := cmd.Process.Kill()
		l.Debug("Kill sent", es_log.Error(err2))
		return err
	}
	if err := termCmd.Wait(); err != nil {
		l.Debug("Terminate wait returned an error", es_log.Error(err))
		return nil
	}

	{
		logName := c.Workspace().JobId()
		logPath := filepath.Join(z.RepositoryPath.Path(), "logs", logName)
		l.Debug("Creating info_finish", es_log.String("path", logPath))

		info := struct {
			TimeLocal string `json:"time_local"`
			TimeUTC   string `json:"time_utc"`
		}{
			TimeLocal: time.Now().Local().Format(time.RFC3339),
			TimeUTC:   time.Now().UTC().Format(time.RFC3339),
		}
		content, err := json.Marshal(&info)
		if err != nil {
			l.Error("Unable to create info file", es_log.Error(err))
		}

		err = ioutil.WriteFile(filepath.Join(logPath, procmonLogFinish), content, 0644)
		if err != nil {
			l.Error("Unable to write info file", es_log.Error(err))
		}
	}

	l.Info("Waiting for termination", es_log.Int("seconds", 60))
	time.Sleep(60 * 1000 * time.Millisecond)

	return nil
}

func (z *Procmon) compressProcmonLogs(c app_control.Control) (arcPath string, err error) {
	logPath := filepath.Join(z.RepositoryPath.Path(), "logs")
	l := c.Log().With(es_log.String("logPath", logPath))

	lstat, err := os.Lstat(logPath)
	if err != nil || !lstat.IsDir() {
		l.Debug("No logs folder found", es_log.Error(err), es_log.Any("lstat", lstat))
		return "", nil
	}

	arcName := c.Workspace().JobId()
	arcPath = filepath.Join(z.RepositoryPath.Path(), arcName+".zip")

	l.Info("Start compress logs", es_log.String("archive", arcPath))
	if err := es_zip.CompressPath(arcPath, logPath, arcName); err != nil {
		l.Debug("Unable to create archive file", es_log.Error(err))
		return "", err
	}

	return arcPath, nil
}

func (z *Procmon) uploadProcmonLogs(c app_control.Control, arcPath string) error {
	if arcPath == "" {
		return nil
	}

	logPath := filepath.Join(z.RepositoryPath.Path(), "logs")
	l := c.Log().With(es_log.String("logPath", logPath))
	l.Info("Start uploading logs", es_log.String("archive", arcPath))

	prof, err := sv_profile.NewProfile(z.Peer.Context()).Current()
	if err != nil {
		l.Error("Unable to retrieve profile", es_log.Error(err))
		return err
	}
	l.Info("Upload to the account", es_log.Any("account", prof))

	e, err := sv_file_content.NewUpload(z.Peer.Context()).Add(z.DropboxPath, arcPath)
	if err != nil {
		l.Error("Unable to upload file", es_log.Error(err))
		return err
	}
	l.Info("Uploaded", es_log.Any("entry", e))
	err = os.Remove(arcPath)
	l.Debug("Removed", es_log.Error(err))

	return nil
}

func (z *Procmon) cleanupProcmonLogs(c app_control.Control) error {
	logPath := filepath.Join(z.RepositoryPath.Path(), "logs")
	l := c.Log().With(es_log.String("logPath", logPath))
	l.Debug("Start clean up logs")

	for i := 0; i < 10; i++ {
		err := os.RemoveAll(logPath)
		if err != nil {
			l.Debug("Unable to clean up logs", es_log.Error(err))
			time.Sleep(10 * 1000 * time.Millisecond)
			continue
		}
		return nil
	}
	return nil
}

func (z *Procmon) Exec(c app_control.Control) error {
	l := c.Log()

	if z.RunUntil.Ok() && z.RunUntil.Time().Before(time.Now()) {
		l.Info("Skip run")
		return nil
	}

	processLogs := func() error {
		logArc, err := z.compressProcmonLogs(c)
		if err != nil {
			return err
		}
		if err = z.uploadProcmonLogs(c, logArc); err != nil {
			return err
		}
		if err = z.cleanupProcmonLogs(c); err != nil {
			return err
		}
		return nil
	}

	exe, err := z.ensureProcmon(c)
	if err != nil {
		return err
	}
	l.Debug("Procmon exe", es_log.String("exe", exe))
	if err = processLogs(); err != nil {
		return err
	}

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
	if err = processLogs(); err != nil {
		return err
	}

	return nil
}

func (z *Procmon) Test(c app_control.Control) error {
	if qt_endtoend.IsSkipEndToEndTest() {
		return nil
	}

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
		m.RepositoryPath = mo_path2.NewFileSystemPath(tmpDir)
		m.DropboxPath = qt_recipe.NewTestDropboxFolderPath("diag-procmon")
		m.Seconds.SetValue(30)
		m.RetainLogs.SetValue(4)
	})
}

func (z *Procmon) Preset() {
	z.ProcmonUrl = procmonDownloadUrl
	z.Seconds.SetRange(10, 86400, 1800)
	z.RetainLogs.SetRange(0, 10000, 4)
}
