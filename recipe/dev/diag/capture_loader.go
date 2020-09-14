package diag

import (
	"bufio"
	"bytes"
	"encoding/json"
	"github.com/watermint/toolbox/essentials/log/esl"
	"github.com/watermint/toolbox/essentials/model/mo_string"
	"github.com/watermint/toolbox/essentials/network/nw_capture"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/control/app_job"
	"github.com/watermint/toolbox/infra/control/app_job_impl"
	"github.com/watermint/toolbox/infra/ui/app_msg"
	"io"
)

type MsgLoader struct {
	ProgressLoadJob app_msg.Message
}

var (
	MLoader = app_msg.Apply(&MsgLoader{}).(*MsgLoader)
)

type CaptureHandler func(history app_job.History, rec nw_capture.Record)

type CaptureLoader struct {
	Ctl   app_control.Control
	JobId mo_string.OptionalString
	Path  mo_string.OptionalString
}

func (z CaptureLoader) Load(handler CaptureHandler) error {
	l := z.Ctl.Log()
	histories, err := app_job_impl.GetHistories(z.Path)
	if err != nil {
		l.Debug("Unable to retrieve histories", esl.Error(err))
		return err
	}

	for _, history := range histories {
		ll := l.With(esl.String("jobId", history.JobId()))
		if z.JobId.IsExists() && history.JobId() != z.JobId.Value() {
			ll.Debug("Skip jobs")
			continue
		}
		if err := z.loadCaptures(history, handler); err != nil {
			ll.Debug("Unable to load capture", esl.Error(err))
		}
	}
	return nil
}

func (z CaptureLoader) loadCaptures(h app_job.History, handler CaptureHandler) error {
	l := z.Ctl.Log().With(esl.String("jobId", h.JobId()))
	z.Ctl.UI().Progress(MLoader.ProgressLoadJob.With("JobId", h.JobId()))

	logs, err := h.Logs()
	if err != nil {
		l.Debug("Unable to list logs", esl.Error(err))
		return err
	}

	for _, log := range logs {
		if log.Type() != app_job.LogFileTypeCapture {
			continue
		}
		if err := z.loadCapture(h, log, handler); err != nil {
			l.Debug("Unable to load capture", esl.Error(err))
		}
	}
	return nil
}

func (z CaptureLoader) loadCapture(history app_job.History, log app_job.LogFile, handler CaptureHandler) error {
	l := z.Ctl.Log().With(esl.String("path", log.Path()))

	var buf bytes.Buffer
	if err := log.CopyTo(&buf); err != nil {
		l.Debug("Unable to read log", esl.Error(err))
		return err
	}

	br := bufio.NewReader(&buf)

	prefix := &bytes.Buffer{}
	for {
		line, isPrefix, err := br.ReadLine()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			l.Debug("Error on read", esl.Error(err))
			return err
		}
		if isPrefix {
			_, err := prefix.Write(line)
			if err != nil {
				l.Debug("Unable to append prefix", esl.Error(err))

				// reset prefix and continue
				prefix.Reset()
				continue
			}
			continue
		}

		if prefix.Len() < 1 {
			if err := z.handleLine(history, line, handler); err != nil {
				l.Debug("Failed process line", esl.Error(err))
			}
		} else {
			_, err := prefix.Write(line)
			if err != nil {
				l.Debug("Unable to append prefix", esl.Error(err))

				// reset prefix and continue
				prefix.Reset()
				continue
			}
			if err := z.handleLine(history, prefix.Bytes(), handler); err != nil {
				l.Debug("Failed process line", esl.Error(err))
			}
			prefix.Reset()
		}
	}
}

func (z CaptureLoader) handleLine(history app_job.History, line []byte, handler CaptureHandler) error {
	l := z.Ctl.Log()
	rec := nw_capture.Record{}
	if err := json.Unmarshal(line, &rec); err != nil {
		l.Debug("Unable to unmarshal", esl.Error(err))
		return err
	}

	handler(history, rec)
	return nil
}
