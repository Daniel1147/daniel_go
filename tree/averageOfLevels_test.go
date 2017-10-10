package tree

import "fmt"
import . "github.com/Daniel1147/tree"
import "testing"

func TestAverageOfLevels(t *testing.T) {
    cases := []struct {
        root *TreeNode
        expected string
    } {
        sample1(),
    }

    for _, testCase := range cases {
        caseDump := testCase.root.Dump()
        // get result
        resultSlice := averageOfLevels(testCase.root)
        result := fmt.Sprintf("%v", resultSlice)

        // build expected
        expected := testCase.expected

        // judge
        if (expected != result) {
            errorMsg := fmt.Sprintf(
                "\ntestCase %v\nexpected %v\nresult %v",
                caseDump,
                expected,
                result,
            )
            t.Errorf(errorMsg)
        }
    }
}

func sample1() (result struct {
    root *TreeNode
    expected string
}) {
    sample := []int{3, 9, 20, NullValue, NullValue, 15, 7}
    root := BuildFromSerialized(sample)

    expectedSlice := []float64{3, 14.5, 11}

    result.root = root
    result.expected = fmt.Sprintf("%v", expectedSlice)
    return
}
