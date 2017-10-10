package main

// import "fmt"
import . "github.com/Daniel1147/tree"

var initNum = -1147

func longestUnivaluePath(root *TreeNode) int {
    maxLen, _ := realLUP(root, 0, initNum)
    return maxLen
}

func realLUP (
    node *TreeNode,
    parentLen int,
    parentVal int,
) (maxLen int, seriesLen int) {
    if node == nil {
        return
    }

    var currentLen int
    if (parentVal == node.Val) {
        currentLen = parentLen + 1
    }

    leftMax, leftLen := realLUP(node.Left, currentLen, node.Val)
    rightMax, rightLen := realLUP(node.Right, currentLen, node.Val)

    if node.Left != nil && node.Left.Val == node.Val {
        leftLen++
        seriesLen = leftLen
    } else {
        leftLen = 0
    }

    if node.Right != nil && node.Right.Val == node.Val {
        rightLen++
        if (seriesLen < rightLen) {
            seriesLen = rightLen
        }
    } else {
        rightLen = 0
    }

    if (rightLen > currentLen && leftLen > currentLen) {
        seriesLen = 0
        maxLen = rightLen + leftLen
        // fmt.Println(node.Val, maxLen)
    }

    if (leftMax > maxLen) {
        maxLen = leftMax
    }

    if (rightMax > maxLen) {
        maxLen = rightMax
    }
    if (currentLen > maxLen) {
        maxLen = currentLen
    }
    return
}
