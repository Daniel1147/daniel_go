package tree

import "fmt"
import . "github.com/Daniel1147/tree"
import "testing"

func TestFindTarget(t *testing.T) {
    cases := []struct {
        root *TreeNode
        target int
        expected bool
    } {
        sample1(),
        sample2(),
        sample3(),
    }

    for _, testCase := range cases {
        caseDump := testCase.root.Dump()
        // get result
        result := findTarget(testCase.root, testCase.target)

        // build expected
        expected := testCase.expected

        // judge
        if (expected != result) {
            errorMsg := fmt.Sprintf(
                "\ntestCase %v\ntarget %v\nexpected %v\nresult %v",
                caseDump,
                testCase.target,
                expected,
                result,
            )
            t.Errorf(errorMsg)
        }
    }
}

func sample1() (result struct {
    root *TreeNode
    target int
    expected bool
}) {
    sample := []int{5, 3, 6, 2, 4, NullValue, 7}
    root := BuildFromSerialized(sample)

    result.root = root
    result.target = 9
    result.expected = true
    return
}

func sample2() (result struct {
    root *TreeNode
    target int
    expected bool
}) {
    sample := []int{5, 3, 6, 2, 4, NullValue, 7}
    root := BuildFromSerialized(sample)

    result.root = root
    result.target = 28
    result.expected = false
    return
}

func sample3() (result struct {
    root *TreeNode
    target int
    expected bool
}) {
    sample := []int{7, NullValue, 2866, 2167, 4853, 727}
    root := BuildFromSerialized(sample)

    result.root = root
    result.target = 2894
    result.expected = true
    return
}
