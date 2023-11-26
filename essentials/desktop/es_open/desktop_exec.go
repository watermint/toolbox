package es_open

import "os/exec"

func desktopOpenExec(executable string, args ...string) OpenOutcome {
	cmd := exec.Command(executable, args...)
	rErr := cmd.Run()
	if rErr != nil {
		return NewOpenOutcomeOpenFailure(rErr)
	}
	return NewOpenOutcomeSuccess()
}
