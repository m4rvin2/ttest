package tabwriter

import (
	"testing"

	"github.com/celicoo/ttest"
)

func TestNew(t *testing.T) {
	ttest.New(t).
		Skip().
		SetSubject(New).
		Run([]ttest.TestCase{}...)
}

func TestMustFlush(t *testing.T) {
	ttest.New(t).
		Skip().
		// SetSubject(New().MustFlush).
		Run([]ttest.TestCase{}...)
}
