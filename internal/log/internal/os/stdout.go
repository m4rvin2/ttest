package os

import (
	"os"
	"syscall"
)

// SetStdout sets the stdout of the process to the underlying file descriptor of
// the given file. This can be used to redirect the output of a program to a
// different file or device
func SetStdout(f File) {
	os.Stdout = f.File
}

// ResetStdout resets the stdout of the process to its default value
func ResetStdout() {
	os.Stdout = os.NewFile(uintptr(syscall.Stdout), "/dev/stdout")
}
