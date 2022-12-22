package log

import (
	"bufio"
	"bytes"
	"fmt"
	"regexp"

	"github.com/celicoo/ttest/internal/log/internal/os"
	"github.com/celicoo/ttest/internal/slice"
)

// CaptureStdout captures the output printed to stdout by fn that matches the
// specified pattern, writes it to the given *bytes.Buffer, and returns the
// value returned by fn
//
// The pattern is applied to each line of output received from the standard
// output. If a line matches the pattern, it will be captured and written to
// the buffer, even if another line before or after does not match the pattern
func CaptureStdout[T any](fn func() T, b *bytes.Buffer, pattern *regexp.Regexp) T {
	// Set the standard output to write to the write end of the pipe
	r, w := os.MustPipe()
	os.SetStdout(w)
	// ...
	defer func() {
		w.MustClose()
		s := bufio.NewScanner(r)
		// ...
		slice.While(s.Scan, func() {
			slice.ForEach(
				pattern.FindStringSubmatch(s.Text()),
				func(submatch *string, i int, _submatches []string) {
					if submatch != nil && i > 0 {
						fmt.Fprintf(b, "%s\n", *submatch)
					}
				})
		})
		// Reset the stdout to the original value before the function returns
		os.ResetStdout()
	}()
	return fn()
}
