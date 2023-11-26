//go:build darwin

package es_open

func desktopOpen(path string) OpenOutcome {
	return desktopOpenExec("open", path)
}
