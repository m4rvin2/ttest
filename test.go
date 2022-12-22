package ttest

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/celicoo/ttest/internal/errors"
	"github.com/celicoo/ttest/internal/function"
	"github.com/celicoo/ttest/internal/slice"
	"github.com/celicoo/ttest/internal/tabwriter"
	"github.com/celicoo/ttest/internal/unicode"
)

// New constructs and returns a Test struct
func New(t *testing.T) Test {
	b := bytes.NewBuffer(nil)
	if t == nil {
		b.Write(errors.NilTestingT)
	}
	return Test{
		t: t,
		b: b,
	}
}

// Test represents a collection of test cases that can be run against a given
// test subject
type Test struct {
	t       *testing.T
	b       *bytes.Buffer
	cases   []TestCase
	subject *function.Value
}

// Parallel signals that this test is to be run in parallel with other
// parallel tests and returns the test itself. When a test is run multiple times
// due to use of -test.count or -test.cpu, multiple instances of a single test
// never run in parallel with each other
func (t Test) Parallel() Test {
	if t.t != nil && t.b.Len() == 0 { // error precedence
		t.t.Parallel()
	}
	return t
}

// Run runs the given test cases and sends the results to stdout. The results
// are only printed if any test case fails or the -test.v flag is set. If the
// test subject is not set, this test fails with the error errors.TestSubjectNotSet
func (t Test) Run(cc ...TestCase) {
	if t.t != nil && t.b.Len() == 0 {
		t.t.Helper()
		if t.subject == nil {
			t.b.Write(errors.TestSubjectNotSet)
		} else {
			t.cases = cc
			slice.ForEach(t.cases, func(c *TestCase, _i int, _cases []TestCase) {
				t.t.Run("", func(subt *testing.T) { // #{id}
					defer func() {
						c.t = subt
						x := recover()
						if x == nil && c.expect() {
							return
						}
						subt.Fail()
					}()
					// The reason for this is to ensure that the correct number of
					// arguments is passed to the Call method (dur). The Call
					// method is variadic, which means it can accept any number of
					// arguments. If c.Give is nil, then passing it as an argument
					// would result in []any{nil} being passed to the Call method,
					// which is not what the code intends. By checking if c.Give is
					// nil and omitting it as an argument if it is, the code ensures
					// that only the intended arguments are passed to the Call method
					if c.Give == nil {
						c.have = t.subject.Call()
					} else {
						c.have = t.subject.Call(c.Give)
					}
				})
			})
		}
	}
	t.log()
}

// SetSubject sets a given value as the test subject and returns the test instance
// itself. If the given value is not a function or method, this test will fail
// with an error message when the Run method is called
func (t Test) SetSubject(a any) Test {
	if t.t != nil && t.b.Len() == 0 {
		t.subject = function.New(a)
		if t.subject == nil {
			t.b.Write(errors.InvalidTestSubject)
		}
	}
	return t
}

// Skip skips this test and returns the test instance itself. If the test fails
// and is then skipped, it is still considered to have failed
func (t Test) Skip() Test {
	if t.t != nil && t.b.Len() == 0 {
		t.t.SkipNow()
	}
	return t
}

// String returns the textual representation of this test. It formats the output
// of the test cases and prints whether they have passed or failed. If there are
// any errors that occur during the set-up of this test, the error message is
// printed instead
func (t Test) String() string {
	// TODO
	if t.b.Len() != 0 {
		return t.b.String()
	}
	w := tabwriter.New(t.b, 0, 0, 2, ' ', 0)
	fmt.Fprintf(w, "[internal]\n")
	fmt.Fprintf(w, "[internal] id\teval\tresult\n")
	slice.ForEach(t.cases, func(c *TestCase, _i int, _cases []TestCase) {
		fmt.Fprintf(w, "[internal] %s\t%s\t", c.name(), c)
		if c.t.Failed() {
			fmt.Fprintf(w, "%s\n", unicode.CharacterDatabase["heavy cross mark"])
		} else {
			fmt.Fprintf(w, "%s\n", unicode.CharacterDatabase["heavy check mark"])
		}
		// TODO
	})
	w.MustFlush()
	return t.b.String()
}

// TODO
func (t Test) log() {
	var log func(...any)
	if t.t == nil {
		log = func(arguments ...any) { // default implementation

		}
	} else {
		log = t.t.Log
		t.t.Helper()
	}
	log(t)
}
