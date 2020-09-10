package log

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

type Last struct {
	rc_recipe.RemarkTransient
	Path              mo_string.OptionalString
	Kind              mo_string.SelectString
	NoticeNoLogFound  app_msg.Message
	ErrorUnableToRead app_msg.Message
}

func (z *Last) Preset() {
	z.Kind.SetOptions(string(app_job.LogFileTypeToolbox), app_job.LogFileTypes...)
}

func (z *Last) Exec(c app_control.Control) error {
	l := c.Log()

	histories, err := app_job_impl.GetHistories(z.Path)
	if err != nil {
		return err
	}
	if len(histories) < 1 {
		return nil
	}

	last := histories[len(histories)-1]
	l.Debug("Last job", esl.String("jobId", last.JobId()))

	logs, err := last.Logs()
	if err != nil {
		l.Debug("Unable to retrieve logs", esl.Error(err))
		return err
	}

	var out io.WriteCloser
	if c.Feature().IsTest() {
		out = es_stdout.NewDiscard()
	} else {
		out = es_stdout.NewDirectOut()
	}

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
	return nil
}

func (z *Last) Test(c app_control.Control) error {
	return rc_exec.Exec(c, &Last{}, rc_recipe.NoCustomValues)
}
