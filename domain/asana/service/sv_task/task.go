package sv_task

import (
	"github.com/watermint/toolbox/domain/asana/api/as_client"
	"github.com/watermint/toolbox/domain/asana/api/as_pagination"
	"github.com/watermint/toolbox/domain/asana/model/mo_project"
	"github.com/watermint/toolbox/domain/asana/model/mo_task"
	"github.com/watermint/toolbox/essentials/api/api_request"
	"github.com/watermint/toolbox/essentials/encoding/es_json"
)

type Task interface {
	Resolve(gid string) (task *mo_task.Task, err error)
	List(opts ...Opt) (tasks []*mo_task.Task, err error)
}

type Opts struct {
	Project string `url:"project,omitempty"`
}

type Opt func(o Opts) Opts

func (z Opts) Apply(opts ...Opt) Opts {
	switch len(opts) {
	case 0:
		return z
	case 1:
		return opts[0](z)
	default:
		return opts[0](z).Apply(opts[1:]...)
	}
}

func Project(prj *mo_project.Project) Opt {
	return func(o Opts) Opts {
		o.Project = prj.Gid
		return o
	}
}
func New(ctx as_client.Client) Task {
	return &taskImpl{
		ctx: ctx,
	}
}

type taskImpl struct {
	ctx as_client.Client
}

func (z taskImpl) Resolve(gid string) (task *mo_task.Task, err error) {
	task = &mo_task.Task{}
	res := z.ctx.Get("tasks/" + gid)
	if err, fail := res.Failure(); fail {
		return nil, err
	}
	err = res.Success().Json().FindModel("data", task)
	return
}

func (z taskImpl) List(opts ...Opt) (tasks []*mo_task.Task, err error) {
	o := Opts{}.Apply(opts...)
	pg := as_pagination.New(z.ctx).WithEndpoint("tasks").WithData(api_request.Query(o))
	tasks = make([]*mo_task.Task, 0)
	err = pg.OnData(func(entry es_json.Json) error {
		t := &mo_task.Task{}
		if errM := entry.Model(t); errM != nil {
			return errM
		} else {
			tasks = append(tasks, t)
			return nil
		}
	})
	return
}
