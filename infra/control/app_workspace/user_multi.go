package app_workspace

import (
	"github.com/watermint/toolbox/essentials/log/esl"
	"path/filepath"
)

func NewMultiApp(home string) (Application, error) {
	if home == "" {
		var err error
		home, err = DefaultAppPath()
		if err != nil {
			return nil, err
		}
	}
	ma := &multiApp{
		home: home,
	}
	err := ma.setup()
	if err != nil {
		return nil, err
	}
	return ma, nil
}

func NewMultiUser(app Application, userHash string) (MultiUser, error) {
	mu := &multiUser{
		home: filepath.Join(app.Home(), NameUser, userHash),
	}
	err := mu.setup()
	if err != nil {
		return nil, err
	}
	return mu, nil
}

func NewMultiJob(user MultiUser) (Workspace, error) {
	mj := &multiJob{
		user:  user,
		jobId: NewJobId(),
	}
	err := mj.setup()
	if err != nil {
		return nil, err
	}
	return mj, nil
}

type multiApp struct {
	home string
}

func (z *multiApp) Home() string {
	return z.home
}

func (z *multiApp) setup() error {
	_, err := getOrCreate(z.home)
	return err
}

type multiUser struct {
	home string
}

func (z *multiUser) Cache() string {
	return filepath.Join(z.UserHome(), NameCache)
}

func (z *multiUser) UserHome() string {
	return z.home
}

func (z *multiUser) Secrets() string {
	return filepath.Join(z.UserHome(), NameSecrets)
}

func (z *multiUser) setup() error {
	_, err := getOrCreate(z.UserHome())
	if err != nil {
		return err
	}
	_, err = getOrCreate(z.Secrets())
	if err != nil {
		return err
	}
	return nil
}

type multiJob struct {
	user  MultiUser
	jobId string
}

func (z *multiJob) Cache() string {
	return z.user.Cache()
}

func (z *multiJob) KVS() string {
	t, err := z.Descendant(NameKvs)
	if err != nil {
		esl.Default().Error("Unable to create KVS folder", esl.Error(err))
		t = filepath.Join(z.Job(), NameKvs)
	}
	return t
}

func (z *multiJob) Test() string {
	t, err := z.Descendant(NameTest)
	if err != nil {
		esl.Default().Error("Unable to create test folder", esl.Error(err))
		t = filepath.Join(z.Job(), NameTest)
	}
	return t
}

func (z *multiJob) Home() string {
	return z.user.UserHome()
}

func (z *multiJob) Secrets() string {
	return z.user.Secrets()
}

func (z *multiJob) Job() string {
	return filepath.Join(z.user.UserHome(), NameJobs, z.JobId())
}

func (z *multiJob) JobId() string {
	return z.jobId
}

func (z *multiJob) Report() string {
	return filepath.Join(z.Job(), NameReport)
}

func (z *multiJob) Log() string {
	return filepath.Join(z.Job(), NameLogs)
}

func (z *multiJob) Descendant(name string) (path string, err error) {
	return getOrCreate(filepath.Join(z.Job(), name))
}

func (z *multiJob) setup() error {
	_, err := getOrCreate(z.Job())
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
