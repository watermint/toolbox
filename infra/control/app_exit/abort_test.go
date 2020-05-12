package app_exit

import "testing"

func TestAbort(t *testing.T) {
	SetTestMode(true)
	expectedCode := FatalGeneral
	defer func() {
		err := recover()
		if err != expectedCode {
			t.Error(err)
		}
	}()

	Abort(expectedCode)
}

func TestExitSuccess(t *testing.T) {
	SetTestMode(true)

	expectedCode := Success
	defer func() {
		err := recover()
		if err != expectedCode {
			t.Error(err)
		}
	}()

	ExitSuccess()
}
