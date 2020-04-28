// +build windows

package es_terminal

import (
	"syscall"
)

// https://forum.golangbridge.org/t/no-println-output-with-go-build-ldflags-h-windowsgui/7633/2
func HideConsole() {
	getConsoleWindow := syscall.NewLazyDLL("kernel32.dll").NewProc("GetConsoleWindow")
	if getConsoleWindow.Find() != nil {
		return
	}

	showWindow := syscall.NewLazyDLL("user32.dll").NewProc("ShowWindow")
	if showWindow.Find() != nil {
		return
	}

	hwnd, _, _ := getConsoleWindow.Call()
	if hwnd == 0 {
		return
	}

	showWindow.Call(hwnd, 0)
}
