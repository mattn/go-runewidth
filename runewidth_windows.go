//go:build windows && !appengine
// +build windows,!appengine

package runewidth

import (
	"os"
	"syscall"
)

var (
	kernel32               = syscall.NewLazyDLL("kernel32")
	procGetConsoleOutputCP = kernel32.NewProc("GetConsoleOutputCP")
)

// IsEastAsian return true if the current locale is CJK
func IsEastAsian() bool {
	if os.Getenv("WT_SESSION") != "" {
		// Windows Terminal is always UTF-8
		return false
	}

	r1, _, _ := procGetConsoleOutputCP.Call()
	if r1 == 0 {
		return false
	}

	switch int(r1) {
	case 932, 51932, 936, 949, 950:
		return true
	}

	return false
}
