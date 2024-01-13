//go:build windows

package cmd

import (
	"syscall"

	"golang.org/x/term"
)

func readPasswordWrapper() ([]byte, error) {
	return term.ReadPassword(int(syscall.Stdin))
}
