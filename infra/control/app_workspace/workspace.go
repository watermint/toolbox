package app_workspace

import (
	"errors"
	"fmt"
	"github.com/watermint/toolbox/essentials/file/es_filepath"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/infra/app"
	"github.com/watermint/toolbox/infra/security/sc_random"
	"go.uber.org/atomic"
	"io/ioutil"
	"os"
	"os/user"
	"path/filepath"
	"time"
)

type Application interface {
	// Toolbox home path
	Home() string
}

type Job interface {
	// Path for job
	Job() string

	// Time of the job started
	JobStartTime() time.Time

	// Job ID
	JobId() string

	// Log path for job
	Log() string

	// Test
	Test() string

	// Report path for job
	Report() string

	// Path for KVS storage
	KVS() string

	// Database Path for database
	Database() string

	// Create or get child folder of job folder
	Descendant(name string) (path string, err error)
}

type User interface {
	// Secrets path
	Secrets() string

	// Cache
	Cache() string
}

type MultiUser interface {
	User

	// User home path
	UserHome() string
}

type Workspace interface {
	Application
	User
	Job
}

const (
	NameSecrets  = "secrets"
	NameCache    = "cache"
	NameUser     = "user"
	NameJobs     = "jobs"
	NameLogs     = "logs"
	NameReport   = "report"
	NameTest     = "test"
	NameKvs      = "kvs"
	NameDatabase = "database"
	JobIdFormat  = "20060102-150405"
)

var (
	jobSequence = atomic.Int64{}
)

func NewJobId() string {
	return fmt.Sprintf("%s.%s", time.Now().Format(JobIdFormat), sc_random.MustGetSecureRandomString(3))
}

func UserHomePath() (path string, err error) {
	u, err := user.Current()
	if err != nil {
		return "", err
	}
	return u.HomeDir, nil
}

func DefaultAppPath() (path string, err error) {
	if eh := os.Getenv(app.EnvNameToolboxHome); eh != "" {
		return eh, nil
	}

	p, err := UserHomePath()
	if err != nil {
		return "", err
	}
	return filepath.Join(p, ".toolbox"), nil
}

func DefaultAppConfigPath() (path string, err error) {
	p, err := UserHomePath()
	if err != nil {
		return "", err
	}
	return filepath.Join(p, ".config", "watermint-toolbox"), nil
}

func GetOrCreateDefaultAppConfigPath() (path string, err error) {
	l := esl.Default()
	path, err = DefaultAppConfigPath()
	if err != nil {
		return "", err
	}
	_, err = os.Lstat(path)
	switch {
	case os.IsNotExist(err):
		l.Debug("Create default app config path", esl.String("path", path))
		mkErr := os.MkdirAll(path, 0755)
		if mkErr != nil {
			l.Debug("Unable to create default app config path", esl.String("path", path), esl.Error(mkErr))
			return "", mkErr
		}
	case err != nil:
		l.Debug("Unable to retrieve path metadata", esl.Error(err))
		return "", err
	}

	return path, nil
}

func NewWorkspace(home string, transient bool) (Workspace, error) {
	if transient {
		if path, err := ioutil.TempDir("", "transient"); err != nil {
			return nil, err
		} else {
			return newWorkspaceWithJobIdNoSetup(path, NewJobId()), nil
		}
	}

	if home == "" {
		if path, err := DefaultAppPath(); err != nil {
			return nil, err
		} else {
			return newWorkspace(path)
		}
	} else {
		if path, err := es_filepath.FormatPathWithPredefinedVariables(home); err != nil {
			return nil, err
		} else {
			return newWorkspace(path)
		}
	}
}

// Create workspace instance by job path.
// This will not create directories. Just matches to existing folder structure.
// Returns error if a
func NewWorkspaceByJobPath(home Application, jobId string) (ws Workspace, err error) {
	ws = newWorkspaceWithJobIdNoSetup(home.Home(), jobId)
	p, err := os.Lstat(ws.Log())
	if err != nil {
		return nil, err
	}
	if p.IsDir() {
		return ws, nil
	}
	return nil, errors.New("given path look like not a workspace path")
}

// create or get fully qualified path
func getOrCreate(fqp string) (path string, err error) {
	l := esl.Default().With(esl.String("path", fqp))
	st, err := os.Stat(fqp)
	switch {
	case err != nil && os.IsNotExist(err):
		err = os.MkdirAll(fqp, 0701)
		if err != nil {
			l.Error("Unable to create workspace path", esl.Error(err))
			return "", err
		}
	case err != nil:
		l.Error("Unable to setup path", esl.Error(err))
		return "", err

	case !st.IsDir():
		l.Error("Workspace path is not a directory")
		return "", errors.New("workspace path is not a directory")

	case st.Mode()&0700 == 0:
		l.Error("No permission to read and write at workspace path", esl.Any("mode", st.Mode()))
		return "", errors.New("no permission")
	}
	return fqp, nil
}
