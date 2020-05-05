package app_workspace

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/es_log"
	"github.com/watermint/toolbox/infra/app"
	"go.uber.org/atomic"
	"os"
	"os/user"
	"path/filepath"
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
	JobIdFormat = "20060102-150405"
)

var (
	jobSequence = atomic.Int64{}
)

func NewJobId() string {
	return fmt.Sprintf("%s.%03d", time.Now().Format(JobIdFormat), jobSequence.Add(1))
}

func DefaultAppPath() (path string, err error) {
	if eh := os.Getenv(app.EnvNameToolboxHome); eh != "" {
		return "", nil
	}

	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return filepath.Join(u.HomeDir, ".toolbox"), nil
}

func newWorkspaceWithPath(path string) (ws Workspace, err error) {
	sws := &singleWorkspace{
		home:  path,
		jobId: NewJobId(),
	}
	err = sws.setup()
	return sws, err
}

func NewWorkspace(home string) (Workspace, error) {
	if home == "" {
		if path, err := DefaultAppPath(); err != nil {
			return nil, err
		} else {
			return newWorkspaceWithPath(path)
		}
	} else {
		if path, err := es_filepath.FormatPathWithPredefinedVariables(home); err != nil {
			return nil, err
		} else {
			return newWorkspaceWithPath(path)
		}
	}
}

// create or get fully qualified path
func getOrCreate(fqp string) (path string, err error) {
	l := es_log.Default().With(es_log.String("path", fqp))
	st, err := os.Stat(fqp)
	switch {
	case err != nil && os.IsNotExist(err):
		err = os.MkdirAll(fqp, 0701)
		if err != nil {
			l.Error("Unable to create workspace path", es_log.Error(err))
			return "", err
		}
	case err != nil:
		l.Error("Unable to setup path", es_log.Error(err))
		return "", err

	case !st.IsDir():
		l.Error("Workspace path is not a directory")
		return "", errors.New("workspace path is not a directory")

	case st.Mode()&0700 == 0:
		l.Error("No permission to read and write at workspace path", es_log.Any("mode", st.Mode()))
		return "", errors.New("no permission")
	}
	return fqp, nil
}
