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

    var leftSecondMin, rightSecondMin int

    if (root.Left.Val == root.Val) {
        leftSecondMin = findSecondMinimumValue(root.Left)
    } else {
        leftSecondMin = root.Left.Val
    }

    if (root.Right.Val == root.Val) {
        rightSecondMin = findSecondMinimumValue(root.Right)
    } else {
        rightSecondMin = root.Right.Val
    }

    if leftSecondMin == -1 && rightSecondMin == -1 {
        return -1
    }

    if leftSecondMin == -1 {
        return rightSecondMin
    }

    if rightSecondMin == -1 {
        return leftSecondMin
    }

    smaller := leftSecondMin
    if (smaller > rightSecondMin) {
        smaller = rightSecondMin
    }

    return smaller
}
