package app_workspace

import (
	"errors"
	"github.com/watermint/toolbox/experimental/app_root"
	"go.uber.org/zap"
	"os"
	"path/filepath"
)

type singleWorkspace struct {
	home  string
	jobId string
}

func (z *singleWorkspace) JobId() string {
	return z.jobId
}

// create or get fully qualified path
func (z *singleWorkspace) getOrCreate(fqp string) (path string, err error) {
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

func (z *singleWorkspace) setup() (err error) {
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

func (z *singleWorkspace) Home() string {
	return z.home
}

func (z *singleWorkspace) Log() string {
	return filepath.Join(z.Job(), "logs")
}

func (z *singleWorkspace) Secrets() string {
	return filepath.Join(z.home, nameSecrets)
}

func (z *singleWorkspace) Job() string {
	return filepath.Join(z.home, "jobs", z.jobId)
}

func (z *singleWorkspace) Descendant(name string) (path string, err error) {
	return z.getOrCreate(filepath.Join(z.Job(), name))
}
