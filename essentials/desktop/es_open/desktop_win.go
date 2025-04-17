//go:build windows

package es_open

import (
	"os"
	"path/filepath"
)

var (
	sys32RunDll = filepath.Join(os.Getenv("SYSTEMROOT"), "SYSTEM32", "rundll32.exe")
)

func desktopOpen(path string) error {
	return desktopOpenExec(sys32RunDll, "url.dll,FileProtocolHandler", path)
}
