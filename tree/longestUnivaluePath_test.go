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
        sample3(),
        sample4(),
        sample5(),
    }

    for _, testCase := range cases {
        caseDump := testCase.root.Dump()
        // get result
        result := longestUnivaluePath(testCase.root)

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
    root := NewTreeNode(5)
    root.Left = NewTreeNode(4)
    root.Right = NewTreeNode(5)
    root.Left.Left = NewTreeNode(1)
    root.Left.Right = NewTreeNode(1)
    root.Right.Right = NewTreeNode(5)

    result.root = root
    result.expected = 2
    return
}

func sample2() (result struct {
    root *TreeNode
    expected int
}) {
    root := NewTreeNode(1)
    root.Left = NewTreeNode(4)
    root.Right = NewTreeNode(5)
    root.Left.Left = NewTreeNode(4)
    root.Left.Right = NewTreeNode(4)
    root.Right.Right = NewTreeNode(5)

    result.root = root
    result.expected = 2
    return
}

func sample3() (result struct {
    root *TreeNode
    expected int
}) {
    root := NewTreeNode(5)
    root.Left = NewTreeNode(4)
    root.Right = NewTreeNode(5)
    root.Left.Left = NewTreeNode(1)
    root.Left.Right = NewTreeNode(1)
    root.Right.Right = NewTreeNode(5)

    result.root = root
    result.expected = 2
    return
}

func sample4() (result struct {
    root *TreeNode
    expected int
}) {
    root := NewTreeNode(1)
    root.Left = NewTreeNode(1)
    root.Right = NewTreeNode(2)
    root.Left.Left = NewTreeNode(1)
    root.Left.Right = NewTreeNode(1)
    root.Right.Right = NewTreeNode(1)

    result.root = root
    result.expected = 2
    return
}

func sample5() (result struct {
    root *TreeNode
    expected int
}) {
    root := NewTreeNode(1)

    root.Left = NewTreeNode(1)
    root.Left.Left = NewTreeNode(1)
    root.Left.Left.Left = NewTreeNode(1)
    root.Left.Right = NewTreeNode(1)
    root.Left.Right.Right = NewTreeNode(1)

    root.Right = NewTreeNode(1)
    root.Right.Right = NewTreeNode(1)
    root.Right.Right.Right = NewTreeNode(1)

    result.root = root
    result.expected = 6
    return
}
