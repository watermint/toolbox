package qt_endtoend

import (
	"errors"
	"os"
	"strconv"
)

const (
	EndToEndPeer        = "end_to_end_test"
	EndToEndTestSkipEnv = "TOOLBOX_SKIPENDTOENDTEST"
)

func IsSkipEndToEndTest() bool {
	if p, found := os.LookupEnv(EndToEndTestSkipEnv); found {
		if b, _ := strconv.ParseBool(p); b {
			return true
		}
	}
	return false
}

var (
	ErrorSkipEndToEndTest = errors.New("skip end to end test")
)

func HumanInteractionRequired() error {
	return &ErrorHumanInteractionRequired{}
}

type ErrorHumanInteractionRequired struct {
}

func (z *ErrorHumanInteractionRequired) Error() string {
	return "human interaction require"
}

func NoTestRequired() error {
	return &ErrorNoTestRequired{}
}

type ErrorNoTestRequired struct {
}

func (z *ErrorNoTestRequired) Error() string {
	return "no test required"
}

func ScenarioTest() error {
	return &ErrorScenarioTest{}
}

type ErrorScenarioTest struct {
}

func (z *ErrorScenarioTest) Error() string {
	return "scenario test"
}

func ImplementMe() error {
	return &ErrorImplementMe{}
}

type ErrorImplementMe struct {
}

func (z *ErrorImplementMe) Error() string {
	return "implement me"
}

func NotEnoughResource() error {
	return &ErrorNotEnoughResource{}
}

type ErrorNotEnoughResource struct {
}

func (z *ErrorNotEnoughResource) Error() string {
	return "not enough resource"
}
