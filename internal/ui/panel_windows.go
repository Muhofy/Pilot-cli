//go:build windows

package ui

import (
	"os"

	"golang.org/x/sys/windows"
)

// enableVirtualTerminal enables ANSI escape code processing on Windows.
// Must be called before any colored output.
func init() {
	stdout := windows.Handle(os.Stdout.Fd())
	var mode uint32
	if err := windows.GetConsoleMode(stdout, &mode); err != nil {
		return
	}
	// ENABLE_VIRTUAL_TERMINAL_PROCESSING = 0x0004
	windows.SetConsoleMode(stdout, mode|0x0004)

	stderr := windows.Handle(os.Stderr.Fd())
	if err := windows.GetConsoleMode(stderr, &mode); err != nil {
		return
	}
	windows.SetConsoleMode(stderr, mode|0x0004)
}