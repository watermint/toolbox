package app_workspace

import (
	"fmt"
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
)

func NewJobId() string {
	return fmt.Sprintf(time.Now().Format("20060102-150405.000"))
}

func NewTempWorkspace() Workspace {
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

func NewSingleUser(home string) (Workspace, error) {
	if home == "" {
		var err error
		home, err = DefaultPath()
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
