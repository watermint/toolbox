//go:build linux

package es_open

func desktopOpen(path string) error {
	return desktopOpenExec("xdg-open", path)
}
