//go:build darwin

package edesktop

func desktopOpen(path string) OpenOutcome {
	return desktopOpenExec("open", path)
}
