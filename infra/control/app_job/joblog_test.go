package app_job

import (
	"github.com/watermint/toolbox/infra/control/app_workspace"
	"testing"
	"time"
)

func TestTimeFromLog(t *testing.T) {
	{
		ts := StartLog{}
		ts.TimeStart = "2020-07-02T14:52:45Z"

		if tm, ok := TimeFromLog(ts, ""); !ok || !tm.Equal(time.Unix(1593701565, 0)) {
			t.Error(tm, ok)
		}
	}

	{
		tf := ResultLog{}
		tf.TimeFinish = "2020-07-02T14:52:45Z"

		if tm, ok := TimeFromLog(tf, ""); !ok || !tm.Equal(time.Unix(1593701565, 0)) {
			t.Error(tm, ok)
		}
	}

	{
		now := time.Now().UTC().Truncate(time.Second)
		jobId := now.Format(app_workspace.JobIdFormat)

		if tm, ok := TimeFromLog(nil, jobId); !ok || !tm.Equal(now) {
			t.Error(tm, ok)
		}
	}

	{
		if tm, ok := TimeFromLog(nil, ""); ok {
			t.Error(tm, ok)
		}
	}
}
