package dd_screenshot

import (
	"errors"
	"github.com/kbinani/screenshot"
	"github.com/watermint/toolbox/essentials/log/esl"
)

var (
	ErrorNoDisplay = errors.New("no display")
)

func CheckDisplayAvailability(displayId int) error {
	l := esl.Default()
	numDisplays := screenshot.NumActiveDisplays()
	if numDisplays < 0 {
		l.Debug("No display available")
		return ErrorNoDisplay
	}
	if displayId < 0 || displayId >= numDisplays {
		// This should not happen because displayId bound by RangeInt
		l.Error("Invalid display id", esl.Int("displayId", displayId), esl.Int("numDisplays", numDisplays))
		return ErrorNoDisplay
	}
	return nil
}
