package app_workspace

import "path/filepath"

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
		app:      app,
		userHash: userHash,
	}
	err := mu.setup()
	if err != nil {
		return nil, err
	}
	return mu, nil
}

func NewMultiJob(user MultiUser) (Job, error) {
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
	app      Application
	userHash string
}

func (z *multiUser) UserHome() string {
	return filepath.Join(z.app.Home(), nameUser, z.userHash)
}

func (z *multiUser) Secrets() string {
	return filepath.Join(z.UserHome(), nameSecrets)
}

func (z *multiUser) setup() error {
	_, err := getOrCreate(z.userHash)
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

func (z *multiJob) Job() string {
	return filepath.Join(z.user.UserHome(), nameJobs)
}

func (z *multiJob) JobId() string {
	return z.jobId
}

func (z *multiJob) Log() string {
	return filepath.Join(z.Job(), nameLogs, z.JobId())
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
	return nil
}
