package main

import "testing"
import . "github.com/Daniel1147/tree"
import "fmt"

func TestIsBalanced(t *testing.T) {
    cases := []struct {
        root *TreeNode
        expected bool
    } {
        {
            sample1(),
            false,
        },
        {
            sample2(),
            true,
        },
        {
            sample3(),
            true,
        },
    }

    for _, testCase := range cases {
        result := isBalanced(testCase.root)
        if (result != testCase.expected) {
            errorMsg := fmt.Sprintf(
                "\ntestCase %v\nexpected %v\nresult %v",
                testCase.root.Dump(),
                testCase.expected,
                result,
            )
            t.Errorf(errorMsg)
        }
    }
}

func sample1() *TreeNode {
    root := NewTreeNode(0)
    root.Left = NewTreeNode(0)
    root.Left.Left = NewTreeNode(0)

    return root
}

func sample2() *TreeNode {
    root := NewTreeNode(0)
    root.Left = NewTreeNode(0)

    return root
}

func sample3() *TreeNode {
    root := NewTreeNode(0)
    root.Left = NewTreeNode(0)
    root.Right = NewTreeNode(0)

    return root
}
