package ttest

import (
	"reflect"
	"strings"
	"testing"
)

// TestCase represents a single test case for a given function or method. It is
// used to define the input (or "given") values and the expected (or "wanted")
// output values for that specific test case. Each TestCase is associated with a
// Test and can be used to run and evaluate the outcome of a single test
// scenario
type TestCase struct {
	t    *testing.T
	have any
	skip bool
	Give any // Give is the input (or "given") values for this test case
	Want any // Want is the expected (or "wanted") output values for this test case
}

// Skip skips this test case and returns the test case instance itself. If the
// test case fails and is then skipped, it is still considered to have failed
func (c TestCase) Skip() TestCase {
	if c.t != nil {
		c.skip = true
		c.t.Skip()
	}
	return c
}

// String returns the textual representation of this test case
func (c TestCase) String() string {
	// TODO
	// return fmt.Sprintf("func(%+v) -> (%+v == %+v)", c.Give, c.Want, c.have)

	// table
	// detailed (paginate)
	return "func(c.Give) -> (c.Want == c.have)"
}

func (c TestCase) expect() bool {
	return reflect.DeepEqual(c.Want, c.have)
}

// name returns the name of this test case. If the embedded test (c.t) is nil,
// an empty string is returned
func (c TestCase) name() string {
	if c.t == nil {
		return ""
	}
	ss := strings.Split(c.t.Name(), "#")
	return ss[1]
}
