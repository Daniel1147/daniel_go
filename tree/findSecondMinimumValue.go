package main

// import "fmt"
import . "github.com/Daniel1147/tree"

func findSecondMinimumValue(root *TreeNode) int {
    if root == nil {
        return -1
    }
    if root.Left == nil {
        return -1
    }

    if root.Left.Val != root.Val {
        return root.Left.Val
    }

    if root.Right.Val != root.Val {
        return root.Right.Val
    }

    leftSecondMin := findSecondMinimumValue(root.Left)
    rightSecondMin := findSecondMinimumValue(root.Right)

    if (leftSecondMin == -1 && rightSecondMin == -1) {
        return -1
    }

    if (leftSecondMin == -1 && rightSecondMin != -1) {
        return rightSecondMin
    }

    if (rightSecondMin == -1 && leftSecondMin != -1) {
        return leftSecondMin
    }

    smaller := leftSecondMin
    if smaller > rightSecondMin {
        smaller = rightSecondMin
    }

    return smaller
}
