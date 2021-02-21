package ui_out

import (
	"github.com/watermint/toolbox/essentials/io/es_stdout"
	"github.com/watermint/toolbox/infra/control/app_control"
	"github.com/watermint/toolbox/infra/ui/app_msg"
)

var (
	textOutCapture string
)

func TextOut(c app_control.Control, text string) {
	if c.Feature().IsTest() {
		textOutCapture = textOutCapture + text
	}
	if c.Feature().IsQuiet() {
		_, _ = es_stdout.NewDirectOut().Write([]byte(text))
	} else {
		c.UI().Info(app_msg.Raw(text))
	}
}

func CapturedText() string {
	return textOutCapture
}
func ClearCaptureText() {
	textOutCapture = ""
}
