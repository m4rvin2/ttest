package slice

import (
	"testing"

	"github.com/celicoo/ttest"
)

func TestForEach(t *testing.T) {
	// INFO: intermediate steps are yet to be supported by ttest, so the test
	//       below isn't doing much
	ttest.New(t).
		SetSubject(ForEach[any]).
		Run(ttest.TestCase{
			Want: nil,
			Give: []any{
				[5]int{1, 2, 3, 4, 5},
				func(int) {},
			},
		})
}
