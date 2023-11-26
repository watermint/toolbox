//go:build linux

package es_open

func desktopOpen(path string) OpenOutcome {
	return desktopOpenExec("xdg-open", path)
}
