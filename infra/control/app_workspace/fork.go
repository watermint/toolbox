package app_workspace

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"path/filepath"
	"strings"
	"time"
)

func Fork(ws Workspace, name string) (Workspace, error) {
	if strings.HasPrefix(name, "-") {
		name = strings.TrimPrefix(name, "-")
	}
	fws := &forkWorkspace{
		name:   name,
		parent: ws,
	}
	if err := fws.setup(); err != nil {
		return nil, err
	}
	return fws, nil
}

type forkWorkspace struct {
	name   string
	parent Workspace
}

func (z *forkWorkspace) setup() (err error) {
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

func (z *forkWorkspace) Cache() string {
	return z.parent.Cache()
}
func (z *forkWorkspace) Home() string {
	return z.parent.Home()
}

func (z *forkWorkspace) Secrets() string {
	return z.parent.Secrets()
}

func (z *forkWorkspace) Job() string {
	return filepath.Join(z.parent.Job(), z.name)
}

func (z *forkWorkspace) JobStartTime() time.Time {
	return z.parent.JobStartTime()
}

func (z *forkWorkspace) JobId() string {
	return z.parent.JobId() + "-" + z.name
}

func (z *forkWorkspace) Log() string {
	return filepath.Join(z.Job(), NameLogs)
}

func (z *forkWorkspace) Test() string {
	t, err := z.Descendant(NameTest)
	if err != nil {
		esl.Default().Error("Unable to create test folder", esl.Error(err))
		t = filepath.Join(z.Job(), NameTest)
	}
	return t
}

func (z *forkWorkspace) KVS() string {
	t, err := z.Descendant(NameKvs)
	if err != nil {
		esl.Default().Error("Unable to create KVS folder", esl.Error(err))
		t = filepath.Join(z.Job(), NameKvs)
	}
	return t
}

func (z *forkWorkspace) Report() string {
	return filepath.Join(z.Job(), NameReport)
}

func (z *forkWorkspace) Descendant(name string) (path string, err error) {
	return getOrCreate(filepath.Join(z.Job(), name))
}
