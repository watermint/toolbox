package ui_out

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

var (
	textOutCapture = make(map[string]string)
)

func TextOut(c app_control.Control, text string) {
	if c.Feature().IsTest() {
		id := c.UI().Id()
		if t, ok := textOutCapture[id]; ok {
			textOutCapture[id] = t + text
		} else {
			textOutCapture[id] = text
		}
	}
	if c.Feature().IsQuiet() {
		_, _ = es_stdout.NewDirectOut().Write([]byte(text))
	} else {
		c.UI().Info(app_msg.Raw(text))
	}
}

func CapturedText(c app_control.Control) string {
	if t, ok := textOutCapture[c.UI().Id()]; ok {
		return t
	} else {
		return ""
	}
}
func ClearCaptureText(c app_control.Control) {
	delete(textOutCapture, c.UI().Id())
}
