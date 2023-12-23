package cat

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/recipe/rc_exec"
	"github.com/watermint/toolbox/infra/recipe/rc_recipe"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"io"
)

type Job struct {
	rc_recipe.RemarkTransient
	Path              mo_string.OptionalString
	Id                string
	Kind              mo_string.SelectString
	ErrorUnableToRead app_msg.Message
}

func (z *Job) Preset() {
	z.Kind.SetOptions(string(app_job.LogFileTypeToolbox), app_job.LogFileTypes...)
}

func (z *Job) Exec(c app_control.Control) error {
	l := c.Log()

	histories, err := app_job_impl.GetHistories(z.Path)
	if err != nil {
		return err
	}

	var out io.WriteCloser
	if c.Feature().IsTest() {
		out = es_stdout.NewDiscard()
	} else {
		out = es_stdout.NewDirectOut()
	}

	for _, h := range histories {
		if h.JobId() != z.Id {
			l.Debug("Skip", esl.String("jobId", h.JobId()))
			continue
		}
		logs, err := h.Logs()
		if err != nil {
			l.Debug("Unable to retrieve logs", esl.Error(err))
			return err
		}
		l.Debug("Last job", esl.String("jobId", h.JobId()))

		for _, lf := range logs {
			if app_job.LogFileType(z.Kind.Value()) != lf.Type() {
				l.Debug("skip non target log type", esl.String("name", lf.Name()), esl.Any("type", lf.Type()))
				continue
			}

			l.Debug("Copying", esl.String("name", lf.Name()))
			if err := lf.CopyTo(out); err != nil {
				l.Debug("Failed copy", esl.Error(err), esl.String("name", lf.Name()))
				c.UI().Error(z.ErrorUnableToRead.With("Name", lf.Name()).With("Error", err))
				continue
			}
		}
	}
	return nil
}

func (z *Job) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Job{}, func(r rc_recipe.Recipe) {
		m := r.(*Job)
		m.Id = "20200512-011129.010"
	})
}
