package app_workspace

import (
	"github.com/watermint/toolbox/infra/control/app_root"
	"go.uber.org/zap"
	"path/filepath"
)

type singleWorkspace struct {
	home  string
	jobId string
}

func (z *singleWorkspace) JobId() string {
	return z.jobId
}

func (z *singleWorkspace) setup() (err error) {
	_, err = getOrCreate(z.home)
	if err != nil {
		return err
	}
	_, err = getOrCreate(z.Secrets())
	if err != nil {
		return err
	}
	_, err = getOrCreate(z.Job())
	if err != nil {
		return err
	}
	_, err = getOrCreate(z.Log())
	if err != nil {
		return err
	}
	_, err = getOrCreate(z.Report())
	if err != nil {
		return err
	}
	return nil
}

func (z *singleWorkspace) Home() string {
	return z.home
}

func (z *singleWorkspace) Log() string {
	return filepath.Join(z.Job(), nameLogs)
}

func (z *singleWorkspace) Secrets() string {
	return filepath.Join(z.home, nameSecrets)
}

func (z *singleWorkspace) Job() string {
	return filepath.Join(z.home, nameJobs, z.jobId)
}

func (z *singleWorkspace) Descendant(name string) (path string, err error) {
	return getOrCreate(filepath.Join(z.Job(), name))
}

func (z *singleWorkspace) Report() string {
	return filepath.Join(z.Job(), nameReport)
}

func (z *singleWorkspace) Test() string {
	t, err := z.Descendant(nameTest)
	if err != nil {
		app_root.Log().Error("Unable to create test folder", zap.Error(err))
		t = filepath.Join(z.Job(), nameTest)
	}
	return t
}
