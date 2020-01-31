// +build windows

package ut_ui

func HideWindow() {
	cw := w32.GetConsoleWindow()
	if cw != 0 {
		_, pid := w32.GetWindowThreadProcessId(cw)
		if w32.GetCurrentProcessId() == pid {
			w32.ShowWindowAsync(cw, w32.SW_HIDE)
		}
	}
}
