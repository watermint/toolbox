package diag

import (
	"archive/zip"
	"errors"
	"github.com/watermint/toolbox/domain/model/mo_path"
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
	"path/filepath"
	"runtime"
	"time"
)

const (
	procmonDownloadUrl = "https://download.sysinternals.com/files/ProcessMonitor.zip"
	procmonExe32       = "ProcMon.exe"
	procmonExe64       = "ProcMon64.exe"
)

type Procmon struct {
	RepositoryPath mo_path.FileSystemPath
	DropboxPath    mo_path.DropboxPath
	Peer           rc_conn.ConnUserFile
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
		l.Debug("Try download", zap.String("url", procmonDownloadUrl))
		resp, err := http.Get(procmonDownloadUrl)
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
		l.Debug("Extract downloaded zip", zap.String("zip", procmonZip))
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

func (z *Procmon) runProcmon(c app_control.Control, exePath string) (cmd *exec.Cmd, err error) {
	l := c.Log()

	logName := c.Workspace().JobId()
	logPath := filepath.Join(z.RepositoryPath.Path(), "logs", logName)
	l.Debug("Creating log path", zap.String("path", logPath))
	err = os.MkdirAll(logPath, 0755)
	if err != nil {
		l.Debug("Unable to create log path", zap.Error(err))
		return nil, err
	}

	cmd = exec.Command(exePath,
		"/AcceptEula",
		"/Quiet",
		"/Minimized",
		"/BackingFile",
		logPath,
	)
	err = cmd.Start()
	if err != nil {
		l.Debug("Unable to start program", zap.Error(err), zap.Any("cmd", cmd))
		return nil, err
	}

	return cmd, nil
}

func (z *Procmon) terminateProcmon(c app_control.Control, exePath string, cmd *exec.Cmd) error {
	l := c.Log()

	l.Debug("Trying to terminate procmon")
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

	err := os.RemoveAll(logPath)
	if err != nil {
		l.Debug("Unable to clean up logs", zap.Error(err))
		return err
	}

	return nil
}

func (z *Procmon) Exec(c app_control.Control) error {
	l := c.Log()

	if z.Seconds < 10 {
		return errors.New("seconds must grater than 10 sec")
	}

	exe, err := z.ensureProcmon(c)
	if err != nil {
		return err
	}
	l.Debug("Procmon exe", zap.String("exe", exe))

	if runtime.GOOS != "windows" {
		l.Info("This command runs only on Windows")
		return nil
	}

	cmd, err := z.runProcmon(c, exe)
	if err != nil {
		return err
	}

	l.Info("Waiting for duration", zap.Int("seconds", z.Seconds))
	time.Sleep(time.Duration(z.Seconds) * time.Second)

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
		m.Seconds = 30
		m.RepositoryPath = mo_path.NewFileSystemPath(tmpDir)
	})
}

func (z *Procmon) Preset() {
	z.Seconds = 1800
}
