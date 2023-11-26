//go:build linux

package edesktop

func desktopOpen(path string) OpenOutcome {
	return desktopOpenExec("xdg-open", path)
}
