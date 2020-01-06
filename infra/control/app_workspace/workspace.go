package app_workspace

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"
)

type Workspace interface {
	Application
	User
	Job
}

const (
	nameSecrets = "secrets"
	nameUser    = "user"
	nameJobs    = "jobs"
	nameLogs    = "logs"
	nameReport  = "report"
	nameTest    = "test"
	nameKvs     = "kvs"
)

func NewJobId() string {
	return fmt.Sprintf(time.Now().Format("20060102-150405.000"))
}

func NewTempAppWorkspace() Workspace {
	home := os.TempDir()
	ws := &singleWorkspace{
		home:  home,
		jobId: NewJobId(),
	}
	err := ws.setup()
	if err != nil {
		panic("Unable to create temporary workspace")
	}
	return ws
}

func DefaultAppPath() (path string, err error) {
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

func NewSingleUser(home string) (Workspace, error) {
	if home == "" {
		var err error
		home, err = DefaultAppPath()
		if err != nil {
			return nil, err
		}
	}

	ws := &singleWorkspace{
		home:  home,
		jobId: NewJobId(),
	}
	err := ws.setup()
	if err != nil {
		return nil, err
	}
	return ws, nil
}

// create or get fully qualified path
func getOrCreate(fqp string) (path string, err error) {
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
