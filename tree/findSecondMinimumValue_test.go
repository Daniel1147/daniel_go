package main

import "fmt"
import . "github.com/Daniel1147/tree"
import "testing"

func TestLongestUnivaluePath(t *testing.T) {
    cases := []struct {
        root *TreeNode
        expected int
    } {
        sample1(),
        sample2(),
    }

    for _, testCase := range cases {
        caseDump := testCase.root.Dump()
        // get result
        result := findSecondMinimumValue(testCase.root)

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
    expected int
}) {
    root := NewTreeNode(2)
    root.Left = NewTreeNode(2)
    root.Right = NewTreeNode(5)
    root.Right.Left = NewTreeNode(5)
    root.Right.Right = NewTreeNode(7)

    result.root = root
    result.expected = 5
    return
}

func sample2() (result struct {
    root *TreeNode
    expected int
}) {
    root := NewTreeNode(1)
    root.Left = NewTreeNode(1)
    root.Right = NewTreeNode(5)
    root.Left.Left = NewTreeNode(1)
    root.Left.Right = NewTreeNode(2)

    result.root = root
    result.expected = 2
    return
}
