//go:build darwin

package es_open

func desktopOpen(path string) error {
	return desktopOpenExec("open", path)
}
