package ttest

import (
	"bytes"
	"os"
	"testing"

	"github.com/celicoo/ttest/internal/function"
)

func TestMain(m *testing.M) {
	code := Run(m)
	os.Exit(code)
}

func TestNew(t *testing.T) {
	New(t).
		SetSubject(New).
		Run(TestCase{
			Give: t,
			Want: Test{
				t,
				bytes.NewBuffer(nil),
				[]TestCase(nil),
				function.New(nil),
			},
		})
}

// func TestTest_Parallel(t *testing.T) {
// 	New(t).
// 		SetSubject(New(t).Parallel).
// 		Run(TestCase{
// 			Want: Test{
// 				t,
// 				bytes.NewBuffer(nil),
// 				[]TestCase(nil),
// 				function.New(nil),
// 			},
// 		})
// }
//
// func TestTest_Run(t *testing.T) {
// 	New(t).
// 		Skip().
// 		SetSubject(New(t).Run).
// 		Run(TestCase{})
// }
//
// func TestTest_SetSubject(t *testing.T) {
// 	New(t).
// 		SetSubject(New(t).SetSubject).
// 		Run([]TestCase{
// 			{
// 				Give: "",
// 				Want: Test{
// 					t,
// 					func() *bytes.Buffer {
// 						var b bytes.Buffer
// 						b.Write(errors.InvalidTestSubject)
// 						return &b
// 					}(),
// 					[]TestCase(nil),
// 					function.New(nil),
// 				},
// 			},
// 			{
// 				Give: TestTest_SetSubject,
// 				Want: Test{
// 					t,
// 					bytes.NewBuffer([]byte{}),
// 					[]TestCase(nil),
// 					function.New(TestTest_SetSubject),
// 				},
// 			},
// 		}...)
// }
//
// func TestTest_Skip(t *testing.T) {
// 	// INFO(celicoo): the Skip method cannot be tested as the underlying
// 	//                implementation stops the test execution by calling
// 	//                runtime.Goexit
// 	New(t).
// 		Skip().
// 		SetSubject(New(t).Skip)
// }
//
// func TestTest_String(t *testing.T) {
// 	New(t).
// 		Skip().
// 		SetSubject(New(t).String).
// 		Run(TestCase{})
// }
