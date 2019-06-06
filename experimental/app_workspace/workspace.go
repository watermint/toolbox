package app_workspace

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/experimental/app_root"
	"go.uber.org/zap"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

type Workspace interface {
	// Toolbox home path
	Home() string

	// Secrets path
	Secrets() string

	// Path for job
	Job() string

	// Job ID
	JobId() string

	// Log path for job
	Log() string

	// Create or get child folder of job folder
	Descendant(name string) (path string, err error)
}

const (
	nameSecrets = "secrets"
)

func NewJobId() string {
	return fmt.Sprintf(time.Now().Format("20060102-150405.000"))
}

func NewTempWorkspace() Workspace {
	home := os.TempDir()
	ws := &wsImpl{
		home:  home,
		jobId: NewJobId(),
	}
	err := ws.setup()
	if err != nil {
		panic("Unable to create temporary workspace")
	}
	return ws
}

func DefaultPath() (path string, err error) {
	for _, e := range os.Environ() {
		v := strings.Split(e, "=")
		if v[0] == "TOOLBOX_HOME" && len(v) > 1 {
			return v[1], nil
		}
	}

	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, ".toolbox"), nil
}

func NewWorkspace(home string) (Workspace, error) {
	if home == "" {
		var err error
		home, err = DefaultPath()
		if err != nil {
			return nil, err
		}
	}

	ws := &wsImpl{
		home:  home,
		jobId: NewJobId(),
	}
	err := ws.setup()
	if err != nil {
		return nil, err
	}
	return ws, nil
}

type wsImpl struct {
	home  string
	jobId string
}

func (z *wsImpl) JobId() string {
	return z.jobId
}

// create or get fully qualified path
func (z *wsImpl) getOrCreate(fqp string) (path string, err error) {
	l := app_root.Log().With(zap.String("path", fqp))
	st, err := os.Stat(fqp)
	switch {
	case err != nil && os.IsNotExist(err):
		err = os.MkdirAll(fqp, 0701)
		if err != nil {
			l.Error("Unable to create workspace path", zap.Error(err))
			return "", err
		}
	case err != nil:
		l.Error("Unable to setup path", zap.Error(err))
		return "", err

	case !st.IsDir():
		l.Error("Workspace path is not a directory")
		return "", errors.New("workspace path is not a directory")

	case st.Mode()&0700 == 0:
		l.Error("No permission to read and write at workspace path", zap.Any("mode", st.Mode()))
		return "", errors.New("no permission")
	}
	return fqp, nil
}

func (z *wsImpl) setup() (err error) {
	_, err = z.getOrCreate(z.home)
	if err != nil {
		return err
	}
	_, err = z.getOrCreate(z.Secrets())
	if err != nil {
		return err
	}
	_, err = z.getOrCreate(z.Job())
	if err != nil {
		return err
	}
	_, err = z.getOrCreate(z.Log())
	if err != nil {
		return err
	}
	return nil
}

func (z *wsImpl) Home() string {
	return z.home
}

func (z *wsImpl) Log() string {
	return filepath.Join(z.Job(), "logs")
}

func (z *wsImpl) Secrets() string {
	return filepath.Join(z.home, nameSecrets)
}

func (z *wsImpl) Job() string {
	return filepath.Join(z.home, "jobs", z.jobId)
}

func (z *wsImpl) Descendant(name string) (path string, err error) {
	return z.getOrCreate(filepath.Join(z.Job(), name))
}
