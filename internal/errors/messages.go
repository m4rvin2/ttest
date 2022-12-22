package errors

var (
	// InvalidTestSubject indicates the test subject that has been set using the
	// Test.SetSubject method is not a function. This error message is printed
	// to stdout when the Test.Run method is called after the Test.SetSubject
	// has been called with an argument that is not a function
	InvalidTestSubject = []byte("test subject is not a function")
	// NilTestingT indicates the test cannot be run without a valid *testing.T
	// This error message is printed to stdout when the Test.Run method is
	// called, and if the Test instance was constructed with a nil *testing.T
	NilTestingT = []byte("test cannot be run without a valid *testing.T")
	// TestSubjectNotSet indicates the test subject has not been set before the
	// test execution. This error message is printed to stdout if the Test.Run
	// method is called before the Test.SetSubject method
	TestSubjectNotSet = []byte("test cannot be run without setting a test subject first")
)
