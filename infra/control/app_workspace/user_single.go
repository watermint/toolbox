package app_workspace

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"path/filepath"
)

func newWorkspace(home string) (Workspace, error) {
	sw := &singleWorkspace{
		home:  home,
		jobId: NewJobId(),
	}
	if err := sw.setup(); err != nil {
		return nil, err
	}
	return sw, nil
}

func newWorkspaceWithJobIdNoSetup(home, jobId string) Workspace {
	return &singleWorkspace{
		home:  home,
		jobId: jobId,
	}
}

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
	return filepath.Join(z.Job(), NameLogs)
}

func (z *singleWorkspace) Cache() string {
	return filepath.Join(z.home, NameCache)
}

func (z *singleWorkspace) Secrets() string {
	return filepath.Join(z.home, NameSecrets)
}

func (z *singleWorkspace) Job() string {
	return filepath.Join(z.home, NameJobs, z.jobId)
}

func (z *singleWorkspace) Descendant(name string) (path string, err error) {
	return getOrCreate(filepath.Join(z.Job(), name))
}

func (z *singleWorkspace) Report() string {
	return filepath.Join(z.Job(), NameReport)
}

func (z *singleWorkspace) KVS() string {
	t, err := z.Descendant(NameKvs)
	if err != nil {
		esl.Default().Error("Unable to create KVS folder", esl.Error(err))
		t = filepath.Join(z.Job(), NameKvs)
	}
	return t
}

func (z *singleWorkspace) Test() string {
	t, err := z.Descendant(NameTest)
	if err != nil {
		esl.Default().Error("Unable to create test folder", esl.Error(err))
		t = filepath.Join(z.Job(), NameTest)
	}
	return t
}
