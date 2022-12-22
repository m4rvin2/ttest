package ttest

import (
	"bytes"
	"fmt"
	"regexp"
	"testing"

	"github.com/celicoo/ttest/internal/log"
)

// Run runs the tests. If the given value is nil -1 is returned. Although not
// strictly necessary, it is recommended to use this function to ensure the
// test output is properly formatted
func Run(m *testing.M) (code int) {
	if m == nil {
		return -1
	} else {
		var b bytes.Buffer
		defer fmt.Printf("%s\n", &b)
		include := regexp.MustCompile("\\[internal\\] (.*)")
		return log.CaptureStdout(m.Run, &b, include)
	}
}
