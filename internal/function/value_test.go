package function

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

func TestFunction_Call(t *testing.T) {
	ttest.New(t).
		Skip().
		SetSubject(New(nil).Call).
		Run([]ttest.TestCase{}...)
}

func TestFunction_Name(t *testing.T) {
	ttest.New(t).
		Skip().
		SetSubject(New(nil).Name).
		Run([]ttest.TestCase{}...)
}
